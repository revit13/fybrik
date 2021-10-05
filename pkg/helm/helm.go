// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package helm

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"

	helmMain "helm.sh/helm/v3/cmd/helm"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/kube"
	"helm.sh/helm/v3/pkg/release"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var (
	debugOption = (os.Getenv("HELM_DEBUG") == "true")
)

const chartPath = "/opt/fybrik/charts/"

func getConfig(kubeNamespace string) (*action.Configuration, error) {
	actionConfig := new(action.Configuration)

	if kubeNamespace == "" {
		kubeNamespace = "default"
	}

	config := &genericclioptions.ConfigFlags{
		Namespace: &kubeNamespace,
	}
	err := actionConfig.Init(config, kubeNamespace, os.Getenv("HELM_DRIVER"), debug)
	if err != nil {
		return nil, err
	}

	return actionConfig, err
}

func debug(format string, v ...interface{}) {
	if debugOption {
		format = fmt.Sprintf("[debug] %s\n", format)
		_ = log.Output(2, fmt.Sprintf(format, v...))
	}
}

const protocolPrefix = "oci"

func ChartRef(hostname string, namespace string, name string, tagname string) string {
	return fmt.Sprintf("%s://%s/%s/%s:%s", protocolPrefix, hostname, namespace, name)
}

// Interface of a helm operations
type Interface interface {
	Uninstall(kubeNamespace string, releaseName string) (*release.UninstallReleaseResponse, error)
	Install(chart *chart.Chart, kubeNamespace string, releaseName string, vals map[string]interface{}) (*release.Release, error)
	Upgrade(chart *chart.Chart, kubeNamespace string, releaseName string, vals map[string]interface{}) (*release.Release, error)
	Status(kubeNamespace string, releaseName string) (*release.Release, error)
	RegistryLogin(hostname string, username string, password string, insecure bool) error
	RegistryLogout(hostname string) error
	Push(chart *chart.Chart, ref string) error
	Pull(ref string) error
	Load(ref string) error
	Package(chartPath string, destinationPath string, version string)
	GetResources(kubeNamespace string, releaseName string) ([]*unstructured.Unstructured, error)
}

// Fake implementation
type Fake struct {
	release   *release.Release
	resources []*unstructured.Unstructured
}

// Uninstall helm release
func (r *Fake) Uninstall(kubeNamespace string, releaseName string) (*release.UninstallReleaseResponse, error) {
	res := &release.UninstallReleaseResponse{}
	r.release = nil
	return res, nil
}

// Install helm release
func (r *Fake) Install(chart *chart.Chart, kubeNamespace string, releaseName string, vals map[string]interface{}) (*release.Release, error) {
	r.release = &release.Release{
		Name: releaseName,
		Info: &release.Info{Status: release.StatusDeployed},
	}
	return r.release, nil
}

// Upgrade helm release
func (r *Fake) Upgrade(kubeNamespace string, releaseName string, vals map[string]interface{}) (*release.Release, error) {
	r.release = &release.Release{
		Name: releaseName,
		Info: &release.Info{Status: release.StatusDeployed},
	}
	return r.release, nil
}

// Status of helm release
func (r *Fake) Status(kubeNamespace string, releaseName string) (*release.Release, error) {
	return r.release, nil
}

// RegistryLogin to docker registry v2
func (r *Fake) RegistryLogin(hostname string, username string, password string, insecure bool) error {
	return nil
}

// RegistryLogout to docker registry v2
func (r *Fake) RegistryLogout(hostname string) error {
	return nil
}

// ChartPush helm chart to repo
func (r *Fake) Push(chart *chart.Chart, ref string) error {
	return nil
}

// ChartPull helm chart from repo
func (r *Fake) Pull(ref string, untar bool, untarDir string) error {
	return nil
}

// Load helm chart
func (r *Fake) Load(ref string) (*chart.Chart, error) {
	return nil, nil
}

// GetResources returns allocated resources for the specified release (their current state)
func (r *Fake) GetResources(kubeNamespace string, releaseName string) ([]*unstructured.Unstructured, error) {
	return r.resources, nil
}

func NewEmptyFake() *Fake {
	return &Fake{
		release:   &release.Release{Info: &release.Info{}},
		resources: make([]*unstructured.Unstructured, 0),
	}
}

func NewFake(release *release.Release, resources []*unstructured.Unstructured) *Fake {
	return &Fake{
		release:   release,
		resources: resources,
	}
}

// Impl implementation
type Impl struct {
}

// Uninstall helm release
func (r *Impl) Uninstall(kubeNamespace string, releaseName string) (*release.UninstallReleaseResponse, error) {
	cfg, err := getConfig(kubeNamespace)
	if err != nil {
		return nil, err
	}
	uninstall := action.NewUninstall(cfg)
	return uninstall.Run(releaseName)
}

// Load helm chart
func (r *Impl) Load(path string) (*chart.Chart, error) {
	chart, err := loader.Load(path)
	if err == nil {
		return chart, err
	}
	return chart, err
}

// Install helm release
func (r *Impl) Install(chart *chart.Chart, kubeNamespace string, releaseName string, vals map[string]interface{}) (*release.Release, error) {
	cfg, err := getConfig(kubeNamespace)
	if err != nil {
		return nil, err
	}
	install := action.NewInstall(cfg)
	install.ReleaseName = releaseName
	install.Namespace = kubeNamespace
	return install.Run(chart, vals)
}

// Upgrade helm release
func (r *Impl) Upgrade(chart *chart.Chart, kubeNamespace string, releaseName string, vals map[string]interface{}) (*release.Release, error) {
	cfg, err := getConfig(kubeNamespace)
	if err != nil {
		return nil, err
	}
	upgrade := action.NewUpgrade(cfg)
	upgrade.Namespace = kubeNamespace
	return upgrade.Run(releaseName, chart, vals)
}

// Status of helm release
func (r *Impl) Status(kubeNamespace string, releaseName string) (*release.Release, error) {
	cfg, err := getConfig(kubeNamespace)
	if err != nil {
		return nil, err
	}
	status := action.NewStatus(cfg)
	return status.Run(releaseName)
}

// RegistryLogin to docker registry v2
func (r *Impl) RegistryLogin(hostname string, username string, password string, insecure bool) error {
	cfg, err := getConfig("")
	if err != nil {
		return err
	}
	return helmMain.ExperimentalRegistryLogin(cfg, nil, hostname, username, password, insecure)
}

// RegistryLogout to docker registry v2
func (r *Impl) RegistryLogout(hostname string) error {
	cfg, err := getConfig("")
	if err != nil {
		return err
	}
	return helmMain.ExperimentalRegistryLogout(cfg, nil, hostname)
}

// Package helm chart from repo
func (r *Impl) Package(chartPath string, destinationPath string, version string) error {
	client := action.NewPackage()
	if version != "" {
		client.Version = version
	}
	client.Destination = destinationPath

	_, err := client.Run(chartPath, nil)
	return err
}

// Push helm chart to repo
func (r *Impl) Push(chart string, remote string) error {
	cfg, err := getConfig("")
	if err != nil {
		return err
	}

	client := experimental.NewPushWithOpts(experimental.WithPushConfig(cfg))
	_, err = client.Run(chart, remote)

	return err

// ChartLoad helm chart from cache
func (r *Impl) ChartLoad(ref string) (*chart.Chart, error) {
	// check for chart mounted in container
	chart, err := loader.Load(chartPath + ref)
	if err == nil {
		return chart, err
	}

	cfg, err := getConfig("")
	if err != nil {
		return nil, err
	}
	load := action.NewChartLoad(cfg)
	return load.Run(ref)
}

// Pull helm chart from repo
func (r *Impl) Pull(ref string, untar bool, untarDir string) error {
	chartRef, err := ParseReference(ref)
	if err != nil {
		return err
	}
	cfg, err := getConfig("")
	if err != nil {
		return err
	}
	client := action.NewPullWithOpts(action.WithConfig(cfg))
	client.Version = chartRef.Tag
	client.Untar = untar
	client.UntarDir = untarDir
	_, err = client.Run("oci://" + chartRef.Repo)
	return err
}

// GetResources returns allocated resources for the specified release (their current state)
func (r *Impl) GetResources(kubeNamespace string, releaseName string) ([]*unstructured.Unstructured, error) {
	resources := make([]*unstructured.Unstructured, 0)
	var rel *release.Release
	var config *action.Configuration
	var err error
	var resourceList kube.ResourceList
	config, err = getConfig(kubeNamespace)
	if err != nil || config == nil {
		return resources, err
	}
	status := action.NewStatus(config)
	rel, err = status.Run(releaseName)
	if err != nil {
		return resources, err
	}
	resourceList, err = config.KubeClient.Build(bytes.NewBufferString(rel.Manifest), false)
	if err != nil {
		return resources, err
	}

	for _, res := range resourceList {
		if err := res.Get(); err != nil {
			return resources, err
		}
		obj := res.Object
		if unstr, ok := obj.(*unstructured.Unstructured); ok {
			resources = append(resources, unstr)
		} else {
			return resources, errors.New("invalid runtime object")
		}
	}
	return resources, nil
}
