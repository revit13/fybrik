// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	"context"

	"emperror.dev/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	app "github.com/ibm/the-mesh-for-data/manager/apis/app/v1alpha1"
	"github.com/ibm/the-mesh-for-data/manager/controllers/utils"
	"k8s.io/apimachinery/pkg/api/equality"
)

// ContextInterface is an interface for communication with a generated resource (e.g. Blueprint)
type ContextInterface interface {
	ResourceExists(ref *app.ResourceReference) bool
	CreateOrUpdateResource(owner *app.ResourceReference, ref *app.ResourceReference, blueprintPerClusterMap map[string]app.BlueprintSpec) error
	DeleteResource(ref *app.ResourceReference) error
	GetResourceStatus(ref *app.ResourceReference) (app.ObservedState, error)
	CreateResourceReference(appName string, appNamespace string) *app.ResourceReference
	GetManagedObject() runtime.Object
}

// Interface for managing Plotter resources

// PlotterInterface context implementation for communication with a single Plotter resource
type PlotterInterface struct {
	Client client.Client
}

// GetManagedObject returns the type of the managed runtime object
func (c *PlotterInterface) GetManagedObject() runtime.Object {
	return &app.Plotter{}
}

// CreateResourceReference returns an identifier (name and namespace) of the generated resource.
func (c *PlotterInterface) CreateResourceReference(appName string, appNamespace string) *app.ResourceReference {
	// Plotter runs in the control plane namespace. Plotter name identifies m4dapplication (name and namespace)
	return &app.ResourceReference{Name: appName + "-" + appNamespace, Namespace: utils.GetSystemNamespace(), Kind: "Plotter"}
}

// ResourceExists checks whether the Plotter resource generated by M4DApplication controller is active
func (c *PlotterInterface) ResourceExists(ref *app.ResourceReference) bool {
	if ref == nil || ref.Namespace == "" {
		return false
	}
	resource := c.GetResourceSignature(ref)
	if err := c.Client.Get(context.Background(), types.NamespacedName{Namespace: ref.Namespace, Name: ref.Name}, resource); err != nil {
		return false
	}
	return true
}

// GetResourceSignature returns the namespaced information of the generated Plotter resource
func (c *PlotterInterface) GetResourceSignature(ref *app.ResourceReference) *app.Plotter {
	return &app.Plotter{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ref.Name,
			Namespace: ref.Namespace,
		},
	}
}

// CreateOrUpdateResource creates a new Plotter resource or updates an existing one
func (c *PlotterInterface) CreateOrUpdateResource(owner *app.ResourceReference, ref *app.ResourceReference, blueprintPerClusterMap map[string]app.BlueprintSpec) error {
	plotter := c.GetResourceSignature(ref)
	if len(blueprintPerClusterMap) == 0 {
		return errors.New("invalid cluster configuration")
	}
	if err := c.Client.Get(context.Background(), types.NamespacedName{Namespace: ref.Namespace, Name: ref.Name}, plotter); err == nil {
		if equality.Semantic.DeepEqual(&plotter.Spec.Blueprints, &blueprintPerClusterMap) {
			// nothing needs to be done
			return nil
		}
	}
	if _, err := ctrl.CreateOrUpdate(context.Background(), c.Client, plotter, func() error {
		plotter.Spec.Blueprints = blueprintPerClusterMap
		plotter.Labels = ownerLabels(types.NamespacedName{Namespace: owner.Namespace, Name: owner.Name})
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// DeleteResource deletes the generated Plotter resource
func (c *PlotterInterface) DeleteResource(ref *app.ResourceReference) error {
	resource := c.GetResourceSignature(ref)
	if err := c.Client.Delete(context.Background(), resource); err != nil {
		return err
	}
	return nil
}

// GetResourceStatus returns the generated Plotter status
func (c *PlotterInterface) GetResourceStatus(ref *app.ResourceReference) (app.ObservedState, error) {
	if ref == nil || ref.Namespace == "" {
		return app.ObservedState{}, nil
	}
	resource := c.GetResourceSignature(ref)
	if err := c.Client.Get(context.Background(), types.NamespacedName{Namespace: ref.Namespace, Name: ref.Name}, resource); err != nil {
		return app.ObservedState{}, err
	}
	return resource.Status.ObservedState, nil
}

// NewPlotterInterface creates a new plotter interface for M4DApplication controller
func NewPlotterInterface(cl client.Client) *PlotterInterface {
	return &PlotterInterface{
		Client: cl,
	}
}