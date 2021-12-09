// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"emperror.dev/errors"
	distributionref "github.com/distribution/distribution/reference"
	"github.com/go-logr/logr"
	credentialprovider "github.com/vdemeester/k8s-pkg-credentialprovider"
	credentialprovidersecrets "github.com/vdemeester/k8s-pkg-credentialprovider/secrets"
	"gopkg.in/yaml.v2"
	"helm.sh/helm/v3/pkg/release"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	kstatus "sigs.k8s.io/cli-utils/pkg/kstatus/status"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	ctrlutil "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	app "fybrik.io/fybrik/manager/apis/app/v1alpha1"
	"fybrik.io/fybrik/manager/controllers"
	"fybrik.io/fybrik/manager/controllers/utils"
	"fybrik.io/fybrik/pkg/environment"
	"fybrik.io/fybrik/pkg/helm"
)

// BlueprintReconciler reconciles a Blueprint object
type BlueprintReconciler struct {
	client.Client
	Name   string
	Log    logr.Logger
	Scheme *runtime.Scheme
	Helmer helm.Interface
}

// Reconcile receives a Blueprint CRD
//nolint:dupl
func (r *BlueprintReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("blueprint", req.NamespacedName)
	var err error

	blueprint := app.Blueprint{}
	if err := r.Get(ctx, req.NamespacedName, &blueprint); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	if res, err := r.reconcileFinalizers(&blueprint); err != nil {
		log.V(0).Info("Could not reconcile finalizers: " + err.Error())
		return res, err
	}

	// If the object has a scheduled deletion time, update status and return
	if !blueprint.DeletionTimestamp.IsZero() {
		// The object is being deleted
		log.V(0).Info("Reconcile: Deleting Blueprint " + blueprint.GetName())
		return ctrl.Result{}, nil
	}

	observedStatus := blueprint.Status.DeepCopy()
	log.V(0).Info("Reconcile: Installing/Updating Blueprint " + blueprint.GetName())

	result, err := r.reconcile(ctx, log, &blueprint)
	if err != nil {
		return ctrl.Result{}, errors.Wrap(err, "failed to reconcile blueprint")
	}

	if !equality.Semantic.DeepEqual(&blueprint.Status, observedStatus) {
		if err := r.Client.Status().Update(ctx, &blueprint); err != nil {
			return ctrl.Result{}, errors.WrapWithDetails(err, "failed to update blueprint status", "status", blueprint.Status)
		}
	}

	log.Info("blueprint reconcile cycle completed", "result", result)
	return result, nil
}

// reconcileFinalizers reconciles finalizers for Blueprint
func (r *BlueprintReconciler) reconcileFinalizers(blueprint *app.Blueprint) (ctrl.Result, error) {
	// finalizer
	finalizerName := r.Name + ".finalizer"
	hasFinalizer := ctrlutil.ContainsFinalizer(blueprint, finalizerName)

	// If the object has a scheduled deletion time, delete it and its associated resources
	if !blueprint.DeletionTimestamp.IsZero() {
		// The object is being deleted
		if hasFinalizer { // Finalizer was created when the object was created
			// the finalizer is present - delete the allocated resources
			if err := r.deleteExternalResources(blueprint); err != nil {
				r.Log.V(0).Info("Error while deleting owned resources: " + err.Error())
			}
			// remove the finalizer from the list and update it, because it needs to be deleted together with the object
			ctrlutil.RemoveFinalizer(blueprint, finalizerName)

			if err := r.Update(context.Background(), blueprint); err != nil {
				return ctrl.Result{}, err
			}
		}
		return ctrl.Result{}, nil
	}
	// Make sure this CRD instance has a finalizer
	if !hasFinalizer {
		ctrlutil.AddFinalizer(blueprint, finalizerName)
		if err := r.Update(context.Background(), blueprint); err != nil {
			return ctrl.Result{}, err
		}
	}
	return ctrl.Result{}, nil
}

func (r *BlueprintReconciler) deleteExternalResources(blueprint *app.Blueprint) error {
	errs := make([]string, 0)
	for release := range blueprint.Status.Releases {
		if _, err := r.Helmer.Uninstall(blueprint.Namespace, release); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "; "))
}

func getDomainFromImageName(image string) (string, error) {
	named, err := distributionref.ParseNormalizedNamed(image)
	if err != nil {
		return "", errors.WithMessage(err, "couldn't parse image name: "+image)
	}

	return distributionref.Domain(named), nil
}

func (r *BlueprintReconciler) applyChartResource(ctx context.Context, log logr.Logger, chartSpec app.ChartSpec, args map[string]interface{}, blueprint *app.Blueprint, releaseName string) (ctrl.Result, error) {
	log.Info(fmt.Sprintf("--- Chart Ref ---\n\n%v\n\n", chartSpec.Name))
	kubeNamespace := blueprint.Spec.ModulesNamespace

	args = CopyMap(args)
	for k, v := range chartSpec.Values {
		SetMapField(args, k, v)
	}
	nbytes, _ := yaml.Marshal(args)
	log.Info(fmt.Sprintf("--- Values.yaml ---\n\n%s\n\n", nbytes))

	var registrySuccessfulLogin string

	if chartSpec.ChartPullSecret != "" {
		// obtain ChartPullSecret
		pullSecret := corev1.Secret{}
		pullSecrets := []corev1.Secret{}

		if err := r.Get(ctx, types.NamespacedName{Namespace: utils.GetSystemNamespace(), Name: chartSpec.ChartPullSecret},
			&pullSecret); err == nil {
			// if this is not a dockerconfigjson, ignore
			if pullSecret.Type == "kubernetes.io/dockerconfigjson" {
				pullSecrets = append(pullSecrets, pullSecret)
			}
		} else {
			return ctrl.Result{}, errors.WithMessage(err, "could not find ChartPullSecret: "+chartSpec.ChartPullSecret)
		}

		if len(pullSecrets) != 0 {
			// create a keyring of all dockerconfigjson secrets, to be used for lookup
			keyring := credentialprovider.NewDockerKeyring()
			keyring, _ = credentialprovidersecrets.MakeDockerKeyring(pullSecrets, keyring)
			repoToPull, err := getDomainFromImageName(chartSpec.Name)
			if err != nil {
				return ctrl.Result{}, errors.WithMessage(err, chartSpec.Name+": failed to parse image name")
			}

			creds, withCredentials := keyring.Lookup(chartSpec.Name)
			if withCredentials {
				for _, cred := range creds {
					err := r.Helmer.RegistryLogin(repoToPull, cred.Username, cred.Password, false)
					if err == nil {
						registrySuccessfulLogin = repoToPull
						break
					} else {
						log.Info("Failed to login to helm registry: " + repoToPull)
					}
				}
			} else {
				log.Info("there is a mismatch between helm chart: " + chartSpec.Name +
					" and the registries associated with secret: " + chartSpec.ChartPullSecret)
			}
		}
	}
	tmpDir, err := ioutil.TempDir("", "fybrik-helm-")
	if err != nil {
		return ctrl.Result{}, errors.WithMessage(err, chartSpec.Name+": failed to create temporary directory for chart pull")
	}

	defer func(log logr.Logger) {
		if err = os.RemoveAll(tmpDir); err != nil {
			log.Info("Error while calling RemoveAll on directory created for pulling helm chart")
		}
	}(log)

	err = r.Helmer.Pull(chartSpec.Name, tmpDir)
	if err != nil {
		return ctrl.Result{}, errors.WithMessage(err, chartSpec.Name+": failed chart pull")
	}
	// if we logged into a registry, let us try to logout
	if registrySuccessfulLogin != "" {
		logoutErr := r.Helmer.RegistryLogout(registrySuccessfulLogin)
		if logoutErr != nil {
			return ctrl.Result{}, errors.WithMessage(err, "failed to logout from helm registry: "+registrySuccessfulLogin)
		}
	}
	chart, err := r.Helmer.Load(chartSpec.Name, tmpDir)
	if err != nil {
		return ctrl.Result{}, errors.WithMessage(err, chartSpec.Name+": failed chart load")
	}

	rel, err := r.Helmer.Status(kubeNamespace, releaseName)
	if err == nil && rel != nil {
		rel, err = r.Helmer.Upgrade(chart, kubeNamespace, releaseName, args)
		if err != nil {
			return ctrl.Result{}, errors.WithMessage(err, chartSpec.Name+": failed upgrade")
		}
	} else {
		rel, err = r.Helmer.Install(chart, kubeNamespace, releaseName, args)
		if err != nil {
			return ctrl.Result{}, errors.WithMessage(err, chartSpec.Name+": failed install")
		}
	}
	log.Info(fmt.Sprintf("--- Release Status ---\n\n%s\n\n", rel.Info.Status))
	return ctrl.Result{}, nil
}

// CopyMap copies a map
func CopyMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			cp[k] = CopyMap(vm)
		} else {
			cp[k] = v
		}
	}

	return cp
}

// SetMapField updates a map
func SetMapField(obj map[string]interface{}, k string, v interface{}) bool {
	components := strings.Split(k, ".")
	for n, component := range components {
		if n == len(components)-1 {
			obj[component] = v
		} else {
			m, ok := obj[component]
			if !ok {
				m := make(map[string]interface{})
				obj[component] = m
				obj = m
			} else if obj, ok = m.(map[string]interface{}); !ok {
				return false
			}
		}
	}
	return true
}

// updateModuleState updates the module state
func (r *BlueprintReconciler) updateModuleState(blueprint *app.Blueprint, instanceName string, isReady bool, err string) {
	state := app.ObservedState{
		Ready: isReady,
		Error: err,
	}
	blueprint.Status.ModulesState[instanceName] = state
}

func (r *BlueprintReconciler) reconcile(ctx context.Context, log logr.Logger, blueprint *app.Blueprint) (ctrl.Result, error) {
	modulesNamespace := blueprint.Spec.ModulesNamespace
	// Gather all templates and process them into a list of resources to apply
	// force-update if the blueprint spec is different
	updateRequired := blueprint.Status.ObservedGeneration != blueprint.GetGeneration()
	blueprint.Status.ObservedGeneration = blueprint.GetGeneration()
	// reset blueprint state
	blueprint.Status.ObservedState.Ready = false
	blueprint.Status.ObservedState.Error = ""
	if blueprint.Status.Releases == nil {
		blueprint.Status.Releases = map[string]int64{}
	}
	if blueprint.Status.ModulesState == nil {
		blueprint.Status.ModulesState = make(map[string]app.ObservedState)
	}
	// count the overall number of Helm releases and how many of them are ready
	numReleases, numReady := 0, 0
	for instanceName, module := range blueprint.Spec.Modules {
		// Add debug information to module labels
		if module.Arguments.Labels == nil {
			module.Arguments.Labels = map[string]string{}
		}
		module.Arguments.Labels[app.BlueprintNameLabel] = blueprint.Name
		module.Arguments.Labels[app.BlueprintNamespaceLabel] = blueprint.Namespace
		// Get arguments by type
		var args map[string]interface{}
		args, err := utils.StructToMap(module.Arguments)
		if err != nil {
			return ctrl.Result{}, errors.WithMessage(err, "Blueprint step arguments are invalid")
		}

		releaseName := utils.GetReleaseName(blueprint.Labels[app.ApplicationNameLabel], blueprint.Labels[app.ApplicationNamespaceLabel], instanceName)
		log.V(0).Info("Release name: " + releaseName)
		numReleases++
		// check the release status
		rel, err := r.Helmer.Status(modulesNamespace, releaseName)
		// unexisting release or a failed release - re-apply the chart
		if updateRequired || err != nil || rel == nil || rel.Info.Status == release.StatusFailed {
			// Process templates with arguments
			chart := module.Chart
			if _, err := r.applyChartResource(ctx, log, chart, args, blueprint, releaseName); err != nil {
				blueprint.Status.ObservedState.Error += errors.Wrap(err, "ChartDeploymentFailure: ").Error() + "\n"
				r.updateModuleState(blueprint, instanceName, false, err.Error())
			} else {
				r.updateModuleState(blueprint, instanceName, false, "")
			}
		} else if rel.Info.Status == release.StatusDeployed {
			status, errMsg := r.checkReleaseStatus(releaseName, modulesNamespace)
			if status == corev1.ConditionFalse {
				blueprint.Status.ObservedState.Error += "ResourceAllocationFailure: " + errMsg + "\n"
				r.updateModuleState(blueprint, instanceName, false, errMsg)
			} else if status == corev1.ConditionTrue {
				r.updateModuleState(blueprint, instanceName, true, "")
				numReady++
			}
		}
		blueprint.Status.Releases[releaseName] = blueprint.Status.ObservedGeneration
	}
	// clean-up
	for release, version := range blueprint.Status.Releases {
		if version != blueprint.Status.ObservedGeneration {
			_, err := r.Helmer.Uninstall(modulesNamespace, release)
			if err != nil {
				log.V(0).Info("Error uninstalling release " + release + " : " + err.Error())
			} else {
				delete(blueprint.Status.Releases, release)
			}
		}
	}
	// check if all releases reached the ready state
	if numReady == numReleases {
		// all modules have been orhestrated successfully - the data is ready for use
		blueprint.Status.ObservedState.Ready = true
		return ctrl.Result{}, nil
	}

	// the status is unknown yet - continue polling
	if blueprint.Status.ObservedState.Error == "" {
		return ctrl.Result{RequeueAfter: 2 * time.Second}, nil
	}
	return ctrl.Result{}, nil
}

// NewBlueprintReconciler creates a new reconciler for Blueprint resources
func NewBlueprintReconciler(mgr ctrl.Manager, name string, helmer helm.Interface) *BlueprintReconciler {
	return &BlueprintReconciler{
		Client: mgr.GetClient(),
		Name:   name,
		Log:    ctrl.Log.WithName("controllers").WithName(name),
		Scheme: mgr.GetScheme(),
		Helmer: helmer,
	}
}

// SetupWithManager registers Blueprint controller
func (r *BlueprintReconciler) SetupWithManager(mgr ctrl.Manager) error {
	// 'UpdateFunc' and 'CreateFunc' used to judge if the event came from within the blueprint's namespace.
	// If that is true, the event will be processed by the reconciler.
	// If it's not then it is a rogue event created by someone outside of the control plane.

	blueprintNamespace := utils.GetSystemNamespace()
	p := predicate.Funcs{
		CreateFunc: func(e event.CreateEvent) bool {
			return e.Object.GetNamespace() == blueprintNamespace
		},
		UpdateFunc: func(e event.UpdateEvent) bool {
			return e.ObjectOld.GetNamespace() == blueprintNamespace
		},
	}
	numReconciles := environment.GetEnvAsInt(controllers.BlueprintConcurrentReconcilesConfiguration, controllers.DefaultBlueprintConcurrentReconciles)
	mgr.GetLogger().Info(fmt.Sprintf("Concurrent blueprint reconciles: %d", numReconciles))

	return ctrl.NewControllerManagedBy(mgr).
		WithOptions(controller.Options{MaxConcurrentReconciles: numReconciles}).
		For(&app.Blueprint{}).
		WithEventFilter(p).
		Complete(r)
}

func (r *BlueprintReconciler) getExpectedResults(kind string) (*app.ResourceStatusIndicator, error) {
	// Assumption: specification for each resource kind is done in one place.
	ctx := context.Background()

	var moduleList app.FybrikModuleList
	if err := r.List(ctx, &moduleList, client.InNamespace(utils.GetSystemNamespace())); err != nil {
		return nil, err
	}
	for _, module := range moduleList.Items {
		for _, res := range module.Spec.StatusIndicators {
			if res.Kind == kind {
				return res.DeepCopy(), nil
			}
		}
	}
	return nil, nil
}

// checkResourceStatus returns the computed state and an error message if exists
func (r *BlueprintReconciler) checkResourceStatus(res *unstructured.Unstructured) (corev1.ConditionStatus, string) {
	// get indications how to compute the resource status based on a module spec
	expected, err := r.getExpectedResults(res.GetKind())
	if err != nil {
		// Could not retrieve the list of modules, will retry later
		return corev1.ConditionUnknown, ""
	}
	if expected == nil {
		// use kstatus to compute the status of the resources for them the expected results have not been specified
		// Current status of a deployed release indicates that the resource has been successfully reconciled
		// Failed status indicates a failure
		computedResult, err := kstatus.Compute(res)
		if err != nil {
			r.Log.V(0).Info("Error computing the status of " + res.GetKind() + " : " + err.Error())
			return corev1.ConditionUnknown, ""
		}
		switch computedResult.Status {
		case kstatus.FailedStatus:
			return corev1.ConditionFalse, computedResult.Message
		case kstatus.CurrentStatus:
			return corev1.ConditionTrue, ""
		default:
			return corev1.ConditionUnknown, ""
		}
	}
	// use expected values to compute the status
	if r.matchesCondition(res, expected.SuccessCondition) {
		return corev1.ConditionTrue, ""
	}
	if r.matchesCondition(res, expected.FailureCondition) {
		return corev1.ConditionFalse, getErrorMessage(res, expected.ErrorMessage)
	}
	return corev1.ConditionUnknown, ""
}

func getErrorMessage(res *unstructured.Unstructured, fieldPath string) string {
	// convert the unstructured data to Labels interface to use Has and Get methods for retrieving data
	labelsImpl := utils.UnstructuredAsLabels{Data: res}
	if !labelsImpl.Has(fieldPath) {
		return ""
	}
	return labelsImpl.Get(fieldPath)
}

func (r *BlueprintReconciler) matchesCondition(res *unstructured.Unstructured, condition string) bool {
	selector, err := labels.Parse(condition)
	if err != nil {
		r.Log.V(0).Info("condition " + condition + "failed to parse: " + err.Error())
		return false
	}
	// get selector requirements, 'selectable' property is ignored
	requirements, _ := selector.Requirements()
	// convert the unstructured data to Labels interface to leverage the package capability of parsing and evaluating conditions
	labelsImpl := utils.UnstructuredAsLabels{Data: res}
	for _, req := range requirements {
		if !req.Matches(labelsImpl) {
			return false
		}
	}
	return true
}

func (r *BlueprintReconciler) checkReleaseStatus(releaseName string, namespace string) (corev1.ConditionStatus, string) {
	// get all resources for the given helm release in their current state
	resources, err := r.Helmer.GetResources(namespace, releaseName)
	if err != nil {
		r.Log.V(0).Info("Error getting resources: " + err.Error())
		return corev1.ConditionUnknown, ""
	}
	// return True if all resources are ready, False - if any resource failed, Unknown - otherwise
	numReady := 0
	for _, res := range resources {
		state, errMsg := r.checkResourceStatus(res)
		r.Log.V(0).Info("Status of " + res.GetKind() + " " + res.GetName() + " is " + string(state))
		if state == corev1.ConditionFalse {
			return state, errMsg
		}
		if state == corev1.ConditionTrue {
			numReady++
		}
	}
	if numReady == len(resources) {
		return corev1.ConditionTrue, ""
	}
	return corev1.ConditionUnknown, ""
}
