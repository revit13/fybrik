// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package argocd

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/rs/zerolog"
	"k8s.io/apimachinery/pkg/runtime"

	app "fybrik.io/fybrik/manager/apis/app/v1beta1"

	"fybrik.io/fybrik/pkg/environment"
	"fybrik.io/fybrik/pkg/logging"
	"fybrik.io/fybrik/pkg/multicluster"
	argoclient "fybrik.io/fybrik/pkg/multicluster/argocd/auto-generated/client"
)

const (
	clusterMetadataConfigMapSL = "/api/v1/namespaces/fybrik-system/configmaps/cluster-metadata"
	gitRepoBranch              = "master"
)

var (
	scheme = runtime.NewScheme()
)

type gitRepo struct {
	password           string
	username           string
	url                string
	fybrikAppsPath     string
	blueprintsAppsPath string
}

// argocdClusterManager for argocd cluster configuration
type argocdClusterManager struct {
	client            *argoclient.APIClient
	log               zerolog.Logger
	argoCDAppsGitRepo gitRepo
}

func NewArgoCDClusterManager(connectionURL, user, password, gitRepoUrl, gitRepoUser, gitRepoPassword, gitRepoFybrikAppsPath,
	gitRepoBlueprintsAppsPath string) (multicluster.ClusterManager, error) {
	logger := logging.LogInit(logging.SETUP, "ArgoCDManager")
	//log := logging.LogInit(logging.SETUP, "datacatalog client")
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	retryClient := retryablehttp.NewClient()
	retryClient.HTTPClient.Transport = &http.Transport{TLSClientConfig: tlsConfig}

	configuration := argoclient.Configuration{
		Servers: argoclient.ServerConfigurations{
			argoclient.ServerConfiguration{
				URL:         connectionURL,
				Description: "Endpoint URL",
			},
		},
		HTTPClient: retryClient.StandardClient(),
	}

	// https://argo-cd.readthedocs.io/en/stable/developer-guide/api-docs/#authorization
	apiClient := argoclient.NewAPIClient(&configuration)

	sessionReq := argoclient.SessionSessionCreateRequest{
		Username: &user,
		Password: &password,
	}

	sess := apiClient.SessionServiceApi.SessionServiceCreate(context.Background())
	sess = sess.Body(sessionReq)
	sessionResp, httpResp, err := apiClient.SessionServiceApi.SessionServiceCreateExecute(sess)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to get bearer token")
		return nil, err
	}
	if httpResp.StatusCode != http.StatusOK {
		logger.Error().Msg("Failed get bearer token: http status code is " + strconv.Itoa(httpResp.StatusCode))
		return nil, errors.New("http status code is " + strconv.Itoa(httpResp.StatusCode))
	}

	token := fmt.Sprintf("Bearer %s", *sessionResp.Token)
	configuration.DefaultHeader = map[string]string{"Authorization": token}

	logger.Info().Msg("Bearer token successfully fetched")

	logger.Info().Msg("Initializing ArgoCD cluster manager")

	return &argocdClusterManager{
		client: argoclient.NewAPIClient(&configuration),
		log:    logger,
		argoCDAppsGitRepo: gitRepo{
			password:           password,
			username:           user,
			url:                gitRepoUrl,
			fybrikAppsPath:     gitRepoFybrikAppsPath,
			blueprintsAppsPath: gitRepoBlueprintsAppsPath,
		},
	}, nil
}

func (cm *argocdClusterManager) getYamlFile(clusterName string) ([]byte, error) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s/%s/fybrik-%s.yaml", cm.argoCDAppsGitRepo.username,
		cm.argoCDAppsGitRepo.url, gitRepoBranch, cm.argoCDAppsGitRepo.fybrikAppsPath, clusterName)

	resp, err := http.Get(url)
	if err != nil {
		cm.log.Error().Err(err).Msg("Failed to get application yaml file: " + url)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cm.log.Error().Err(err).Msg("Failed to get application yaml file: " + url)
		return nil, err
	}
	return body, nil

}

// GetClusters returns a list of registered clusters
func (cm *argocdClusterManager) GetClusters() ([]multicluster.Cluster, error) {
	cm.log.Info().Msg("list clusters")
	//var clusters []multicluster.Cluster
	req := cm.client.ClusterServiceApi.ClusterServiceList(context.Background())

	clusters, httpResp, err := cm.client.ClusterServiceApi.ClusterServiceListExecute(req)
	if err != nil {
		cm.log.Error().Err(err).Msg("Failed to list clusters")
		return nil, err
	}
	if httpResp.StatusCode != http.StatusOK {
		cm.log.Error().Msg("Failed to list clusters: http status code is " + strconv.Itoa(httpResp.StatusCode))
		return nil, errors.New("http status code is " + strconv.Itoa(httpResp.StatusCode))
	}
	if !clusters.HasItems() {
		cm.log.Error().Msg("Failed to list clusters: no cluster exists")
		return nil, errors.New("no cluster exists")
	}
	clustersList := clusters.GetItems()
	for _, cluster := range clustersList {
		cm.log.Info().Msg("cluster name: " + *cluster.Name)
		_, err := cm.getYamlFile(*cluster.Name)
		if err != nil {
			return nil, err
		}

	}
	return nil, nil
}

func (cm *argocdClusterManager) IsMultiClusterSetup() bool {
	return false
}

// GetBlueprint returns a blueprint matching the given name, namespace and cluster details
func (cm *argocdClusterManager) GetBlueprint(cluster, namespace, name string) (*app.Blueprint, error) {
	if cluster != environment.GetLocalClusterName() {
		return nil, fmt.Errorf("unregistered cluster: %s", cluster)
	}

	return nil, nil
}

// CreateBlueprint creates a blueprint resource or updates an existing one
func (cm *argocdClusterManager) CreateBlueprint(cluster string, blueprint *app.Blueprint) error {
	return cm.UpdateBlueprint(cluster, blueprint)
}

// UpdateBlueprint updates the given blueprint or creates a new one if it does not exist
func (cm *argocdClusterManager) UpdateBlueprint(cluster string, blueprint *app.Blueprint) error {
	return nil
}

// DeleteBlueprint deletes the blueprint resource
func (cm *argocdClusterManager) DeleteBlueprint(cluster, namespace, name string) error {
	return nil
}
