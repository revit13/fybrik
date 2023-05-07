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
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"

	app "fybrik.io/fybrik/manager/apis/app/v1beta1"

	"fybrik.io/fybrik/pkg/environment"
	"fybrik.io/fybrik/pkg/logging"
	"fybrik.io/fybrik/pkg/multicluster"
	argoclient "fybrik.io/fybrik/pkg/multicluster/argocd/auto-generated/client"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
	githttp "github.com/go-git/go-git/v5/plumbing/transport/http"
)

const (
	gitRepoBranch      = "main"
	appBlueprintPrefix = "blueprints-"
)

var (
	pushRepoMutex sync.Mutex
	scheme        = runtime.NewScheme()
)

func init() {
	_ = app.AddToScheme(scheme)
}

type gitRepo struct {
	password           string
	username           string
	url                string
	blueprintsAppsPath string
}

// argocdClusterManager for argocd cluster configuration
type argocdClusterManager struct {
	client                     *argoclient.APIClient
	log                        zerolog.Logger
	argoCDAppsGitRepo          gitRepo
	argocdFybrikAppsNamePrefix string
}

func NewArgoCDClusterManager(connectionURL, user, password, gitRepoUrl, gitRepoUser, gitRepoPassword, argocdFybrikAppsNamePrefix,
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
			password:           gitRepoPassword,
			username:           gitRepoUser,
			url:                gitRepoUrl,
			blueprintsAppsPath: gitRepoBlueprintsAppsPath,
		},
		argocdFybrikAppsNamePrefix: argocdFybrikAppsNamePrefix,
	}, nil
}

func (cm *argocdClusterManager) packClusterConfigMap(params map[string]string) *v1.ConfigMap {
	configMap := v1.ConfigMap{}
	configMap.Data = params
	return &configMap
}

func (cm *argocdClusterManager) getClusterInfo(clusterName string) (multicluster.Cluster, error) {
	var cluster multicluster.Cluster
	req := cm.client.ApplicationServiceApi.ApplicationServiceGet(context.Background(),
		cm.argocdFybrikAppsNamePrefix+"-"+clusterName)
	cm.log.Info().Msg("application name: " + cm.argocdFybrikAppsNamePrefix + "-" + clusterName)
	argocdApplication, httpResp, err := cm.client.ApplicationServiceApi.ApplicationServiceGetExecute(req)
	if err != nil {
		cm.log.Error().Err(err).Msg("Failed to get argocd application")
		return cluster, err
	}
	if httpResp.StatusCode != http.StatusOK {
		cm.log.Error().Msg("Failed to get argocd application: http status code is " + strconv.Itoa(httpResp.StatusCode))
		return cluster, errors.New("http status code is " + strconv.Itoa(httpResp.StatusCode))
	}
	fybrikHelmParams := argocdApplication.GetSpec().Source.Helm.GetParameters()
	var params = make(map[string]string)

	for _, helmParam := range fybrikHelmParams {
		switch helmParam.GetName() {
		case "cluster.region":
			params["Region"] = helmParam.GetValue()
			cm.log.Info().Msg("region: " + helmParam.GetValue())
		case "cluster.zone":
			params["Zone"] = helmParam.GetValue()
			cm.log.Info().Msg("zone: " + helmParam.GetValue())
		case "cluster.name":
			params["ClusterName"] = helmParam.GetValue()
			cm.log.Info().Msg("ClusterName: " + helmParam.GetValue())
		case "cluster.vaultAuthPath":
			params["VaultAuthPath"] = helmParam.GetValue()
			cm.log.Info().Msg("VaultAuthPath: " + helmParam.GetValue())
		}
	}

	return multicluster.CreateCluster(*cm.packClusterConfigMap(params)), nil
}

func (cm *argocdClusterManager) cloneGitRepo() (string, *git.Repository, error) {
	cm.log.Info().Msg(cm.argoCDAppsGitRepo.username + cm.argoCDAppsGitRepo.url)
	tmpDir, err := os.MkdirTemp(environment.GetDataDir(), "blueprints-repo")
	r, err := git.PlainClone(tmpDir, false, &git.CloneOptions{
		Auth: &githttp.BasicAuth{
			Username: cm.argoCDAppsGitRepo.username,
			Password: cm.argoCDAppsGitRepo.password,
		},
		URL: cm.argoCDAppsGitRepo.url,
		//Progress: os.Stdout,
	})
	if err != nil {
		cm.log.Error().Err(err).Msg("Failed to clone repo")
		return "", nil, err
	}
	return tmpDir, r, nil
}

// GetClusters returns a list of registered clusters
func (cm *argocdClusterManager) GetClusters() ([]multicluster.Cluster, error) {
	cm.log.Info().Msg("list clusters")
	var clusters []multicluster.Cluster
	req := cm.client.ClusterServiceApi.ClusterServiceList(context.Background())

	clustersList, httpResp, err := cm.client.ClusterServiceApi.ClusterServiceListExecute(req)
	if err != nil {
		cm.log.Error().Err(err).Msg("Failed to list clusters")
		return nil, err
	}
	if httpResp.StatusCode != http.StatusOK {
		cm.log.Error().Msg("Failed to list clusters: http status code is " + strconv.Itoa(httpResp.StatusCode))
		return nil, errors.New("http status code is " + strconv.Itoa(httpResp.StatusCode))
	}
	if !clustersList.HasItems() {
		cm.log.Error().Msg("Failed to list clusters: no cluster exists")
		return nil, errors.New("no cluster exists")
	}
	for _, clusterItem := range clustersList.GetItems() {
		name := clusterItem.GetName()
		cm.log.Info().Msg("cluster name: " + name)
		cluster, err := cm.getClusterInfo(name)
		if err != nil {
			cm.log.Error().Err(err).Msg("Failed to list clusters")
			return nil, err
		}
		clusters = append(clusters, cluster)
	}

	return clusters, nil
}

func (cm *argocdClusterManager) IsMultiClusterSetup() bool {
	return true
}

func (cm *argocdClusterManager) getBlueprintFileName(blueprintName, blueprintNamespace string) string {
	return blueprintName + "-" + blueprintNamespace + ".yaml"
}
func (cm *argocdClusterManager) getBlueprintFilePath(cluster string) string {
	return "blueprints" + "/" + cluster + "/"
}

// GetBlueprint returns a blueprint matching the given name, namespace and cluster details
func (cm *argocdClusterManager) GetBlueprint(cluster, namespace, name string) (*app.Blueprint, error) {
	cm.log.Info().Msg("Get Blueprint " + " cluster " + cluster + "namespace: " + namespace + " name: " + name)
	req := cm.client.ApplicationServiceApi.ApplicationServiceGetManifests(context.Background(), appBlueprintPrefix+cluster)
	req.AppNamespace("argocd")
	resp, httpResp, err := cm.client.ApplicationServiceApi.ApplicationServiceGetManifestsExecute(req)
	if err != nil {
		cm.log.Error().Err(err).Msg("Failed to get application manifest")
		return nil, err
	}
	if httpResp.StatusCode != http.StatusOK {
		cm.log.Error().Msg("Failed to get application manifest: http status code is " + strconv.Itoa(httpResp.StatusCode))
		return nil, errors.New("http status code is " + strconv.Itoa(httpResp.StatusCode))
	}
	cm.log.Info().Msg("print manifest")
	blueprint := app.Blueprint{}
	found := false
	for _, manifest := range resp.GetManifests() {
		err = multicluster.Decode(manifest, scheme, &blueprint)
		if err != nil {
			return nil, err
		}

		if blueprint.Namespace == "" {
			log.Warn().Msg("Retrieved an empty blueprint")
			return nil, nil
		}
		cm.log.Info().Msg("found manifest for " + blueprint.GetName())

		if blueprint.GetName() == name {
			cm.log.Info().Msg(manifest)
			found = true
			break
		}
	}
	if !found {
		err = errors.New("blueprint not found for " + name)
		cm.log.Error().Err(err).Msg("Failed to get blueprint")
		return nil, err
	}
	cm.log.Info().Msg("blueprint successfully read " + blueprint.Namespace)

	return &blueprint, nil
}

// CreateBlueprint creates a blueprint resource or updates an existing one
func (cm *argocdClusterManager) CreateBlueprint(cluster string, blueprint *app.Blueprint) error {
	cm.log.Info().Msg("Create Blueprint " + " cluster " + cluster + "namespace: " + blueprint.Namespace + " name: " + blueprint.Name)
	repoDir, repo, err := cm.cloneGitRepo()
	defer os.RemoveAll(repoDir)
	if err != nil {
		cm.log.Error().Err(err).Msg("Failed to create blueprint")
		return err
	}
	fileName := cm.getBlueprintFileName(blueprint.Name, blueprint.Namespace)
	w, err := repo.Worktree()
	if err != nil {
		cm.log.Error().Err(err).Msg("Failed to create blueprint")
	}

	fullFilename := filepath.Join(repoDir+"/"+cm.getBlueprintFilePath(cluster), fileName)

	content, err := yaml.Marshal(blueprint)
	if err != nil {
		return err
	}
	cm.log.Info().Msg("fullPath: " + fullFilename)
	err = os.MkdirAll(repoDir+"/"+cm.getBlueprintFilePath(cluster), os.ModePerm)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fullFilename, content, 0644)
	if err != nil {
		return err
	}
	cm.log.Info().Msg("do git add of blueprint " + cm.getBlueprintFilePath(cluster) + fileName)
	_, err = w.Add(cm.getBlueprintFilePath(cluster) + fileName)
	if err != nil {
		return err
	}

	// Commits the current staging area to the repository, with the new file
	// just created. We should provide the object.Signature of Author of the
	// commit Since version 5.0.1, we can omit the Author signature, being read
	// from the git config files.
	cm.log.Info().Msg("do git commit of blueprint " + fileName)
	commit, err := w.Commit("create "+fileName+" blueprint", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "fybrik",
			Email: "fybrik@fybrik",
			When:  time.Now(),
		},
	})

	remote, err := repo.Remote("origin")

	cm.log.Info().Msg("commit hash " + commit.String())
	po := &git.PushOptions{
		Auth: &githttp.BasicAuth{
			Username: cm.argoCDAppsGitRepo.username,
			Password: cm.argoCDAppsGitRepo.password,
		},
		RemoteName:      "origin",
		RefSpecs:        []config.RefSpec{config.RefSpec("refs/heads/*:refs/heads/*")},
		Progress:        os.Stdout,
		Force:           false,
		InsecureSkipTLS: true,
	}
	cm.log.Info().Msg("do git push of blueprint " + fileName)
	// mutex for the writing operation
	pushRepoMutex.Lock()
	defer pushRepoMutex.Unlock()
	err = remote.Push(po)
	if err != nil {
		return err
	}
	cm.log.Info().Msg("Successfully created blueprint!")
	return nil
}

// UpdateBlueprint updates the given blueprint or creates a new one if it does not exist
func (cm *argocdClusterManager) UpdateBlueprint(cluster string, blueprint *app.Blueprint) error {
	cm.log.Info().Msg("Update Blueprint " + " cluster " + cluster + "namespace: " + blueprint.Namespace + " name: " + blueprint.Name)
	return cm.CreateBlueprint(cluster, blueprint)
}

// DeleteBlueprint deletes the blueprint resource
func (cm *argocdClusterManager) DeleteBlueprint(cluster, namespace, name string) error {
	repoDir, repo, err := cm.cloneGitRepo()
	defer os.RemoveAll(repoDir)
	if err != nil {
		cm.log.Error().Err(err).Msg("Failed to delete blueprint")
		return err
	}
	fileName := cm.getBlueprintFileName(name, namespace)
	fullFilename := filepath.Join(repoDir+"/"+cm.getBlueprintFilePath(cluster), fileName)

	cm.log.Info().Msg("fullPath: " + fullFilename)

	cm.log.Info().Msg("do git remove of blueprint " + cm.getBlueprintFilePath(cluster) + fileName)
	w, err := repo.Worktree()
	if err != nil {
		cm.log.Error().Err(err).Msg("Failed to delete blueprint")
	}

	_, err = w.Remove(cm.getBlueprintFilePath(cluster) + fileName)
	if err != nil {
		return err
	}

	// Commits the current staging area to the repository, with the new file
	// just created. We should provide the object.Signature of Author of the
	// commit Since version 5.0.1, we can omit the Author signature, being read
	// from the git config files.
	cm.log.Info().Msg("do git remove of blueprint " + fileName)
	commit, err := w.Commit("remove "+fileName+" blueprint", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "fybrik",
			Email: "fybrik@fybrik",
			When:  time.Now(),
		},
	})

	remote, err := repo.Remote("origin")

	cm.log.Info().Msg("commit hash " + commit.String())
	po := &git.PushOptions{
		Auth: &githttp.BasicAuth{
			Username: cm.argoCDAppsGitRepo.username,
			Password: cm.argoCDAppsGitRepo.password,
		},
		RemoteName:      "origin",
		RefSpecs:        []config.RefSpec{config.RefSpec("refs/heads/*:refs/heads/*")},
		Progress:        os.Stdout,
		Force:           false,
		InsecureSkipTLS: true,
	}
	cm.log.Info().Msg("do git push of blueprint deletion" + fileName)
	// mutex for the writing operation
	pushRepoMutex.Lock()
	defer pushRepoMutex.Unlock()
	err = remote.Push(po)
	if err != nil {
		return err
	}
	cm.log.Info().Msg("Successfully deleted blueprint!")
	return nil
}
