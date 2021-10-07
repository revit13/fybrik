// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package helm

import (
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/release"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	kstatus "sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

func buildTestChart() *chart.Chart {
	testManifestWithHook := `apiVersion: v1
kind: ConfigMap
metadata:
  name: test-cm
data:
  key: value`

	return &chart.Chart{
		Metadata: &chart.Metadata{
			APIVersion: "v1",
			Name:       "test-chart",
			Type:       "application",
			Version:    "0.1.0",
		},
		Templates: []*chart.File{
			{Name: "templates/config.yaml", Data: []byte(testManifestWithHook)},
		},
	}
}

func Log(t *testing.T, label string, err error) {
	if err == nil {
		err = fmt.Errorf("succeeded")
	}
	t.Logf("%s: %s", label, err)
}

var (
	kubeNamespace = "default"
	releaseName   = "test-install-release"
	chartName     = "test-chart"
	tagName       = "0.1.0"
	hostname      = os.Getenv("DOCKER_HOSTNAME")
	namespace     = os.Getenv("DOCKER_NAMESPACE")
	username      = os.Getenv("DOCKER_USERNAME")
	password      = os.Getenv("DOCKER_PASSWORD")
	insecure, _   = strconv.ParseBool(os.Getenv("DOCKER_INSECURE"))
	chartRef      = ChartRef(hostname, namespace, chartName, tagName)
	impl          = new(Impl)
)

func TestHelmRegistry(t *testing.T) {
	var err error
	// Test should only run as integration test if registry is available
	if _, isSet := os.LookupEnv("DOCKER_HOSTNAME"); !isSet {
		t.Skip("No integration environment found. Skipping test...")
	}

	if username != "" && password != "" {
		err = impl.RegistryLogin(hostname, username, password, insecure)
		assert.Nil(t, err)
		Log(t, "registry login", err)
	}

	err = impl.Package("testdata/test-chart", "/tmp/", "")
	assert.Nil(t, err)
	Log(t, "package chart", err)

	ref, err := ParseReference(chartRef)
	assert.Nil(t, err)

	packagePath := "/tmp/test-chart-0.1.0.tgz"
	chartPath := "oci://" + ref.Repo
	err = impl.Push(packagePath, chartPath)
	assert.Nil(t, err)
	Log(t, "push chart", err)

	err = impl.Pull(chartRef, true)
	assert.Nil(t, err)
	Log(t, "pull chart", err)

	pulledChart, err := impl.Load(chartRef)
	assert.Nil(t, err)

	origChart, err := loader.Load(packagePath)
	assert.Nil(t, err)

	assert.Equal(t, origChart.Metadata.Name, pulledChart.Metadata.Name, "expected loaded chart equals saved chart")

	if username != "" && password != "" {
		err = impl.RegistryLogout(hostname)
		assert.Nil(t, err)
		Log(t, "registry logout", err)
	}

}

func TestHelmRelease(t *testing.T) {
	// Test should only run as integration test if registry is available
	if _, isSet := os.LookupEnv("DOCKER_HOSTNAME"); !isSet {
		t.Skip("No integration environment found. Skipping test...")
	}
	var err error
	origChart := buildTestChart()

	_, _ = impl.Uninstall(kubeNamespace, releaseName)
	vals := map[string]interface{}{
		"data": map[string]interface{}{
			"key": "value1",
		},
	}
	_, err = impl.Install(origChart, kubeNamespace, releaseName, vals)
	assert.Nil(t, err)
	Log(t, "install", err)

	_, err = impl.Upgrade(origChart, kubeNamespace, releaseName, vals)
	assert.Nil(t, err)
	Log(t, "upgrade", err)

	assert.Eventually(t, func() bool {
		rel, err := impl.Status(kubeNamespace, releaseName)
		assert.Nil(t, err)
		return rel.Info.Status == release.StatusDeployed
	}, time.Minute, time.Second)
	Log(t, "status", err)

	var resources []*unstructured.Unstructured
	resources, err = impl.GetResources(kubeNamespace, releaseName)
	assert.Nil(t, err)
	assert.Len(t, resources, 1)
	computedResult, _ := kstatus.Compute(resources[0])
	assert.Equal(t, kstatus.CurrentStatus, computedResult.Status)
	Log(t, "getResources", err)

	_, err = impl.Uninstall(kubeNamespace, releaseName)
	assert.Nil(t, err)
	Log(t, "uninstall", err)
}
