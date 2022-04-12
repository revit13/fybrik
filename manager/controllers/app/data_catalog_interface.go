// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package app

import (
	app "fybrik.io/fybrik/manager/apis/app/v1alpha1"
	"fybrik.io/fybrik/manager/controllers/utils"
	"fybrik.io/fybrik/pkg/model/datacatalog"
	"fybrik.io/fybrik/pkg/vault"

	"github.com/rs/zerolog/log"
)

// RegisterAsset registers a new asset in the specified catalog
// Input arguments:
// - catalogID: the destination catalog identifier
// - info: connection and credential details
// Returns:
// - an error if happened
// - the new asset identifier
func (r *FybrikApplicationReconciler) RegisterAsset(assetID string, catalogID string, info *app.DatasetDetails,
	input *app.FybrikApplication) (string, error) {
	r.Log.Trace().Msg("RegisterAsset")
	details := datacatalog.ResourceDetails{}
	if info.Details != nil {
		infoDetails := info.Details
		details.Connection = infoDetails.Connection
		details.DataFormat = infoDetails.Format
	}

	resourceMetadata := datacatalog.ResourceMetadata{
		Name: assetID,
	}

	creds := ""
	if utils.IsVaultEnabled() {
		creds = vault.PathForReadingKubeSecret(info.SecretNamespace, info.SecretName)
	}

	request := datacatalog.CreateAssetRequest{
		ResourceMetadata:     resourceMetadata,
		Details:              details,
		Credentials:          creds,
		DestinationCatalogID: catalogID,
		DestinationAssetID:   assetID,
	}

	credentialPath := ""
	if utils.IsVaultEnabled() {
		credentialPath = utils.GetVaultAddress() + vault.PathForReadingKubeSecret(input.Namespace, input.Spec.SecretRef)
	}

	var err error
	var response *datacatalog.CreateAssetResponse
	if response, err = r.DataCatalog.CreateAsset(&request, credentialPath); err != nil {
		log.Error().Err(err).Msg("failed to receive the catalog connector response")
		return "", err
	}

	return response.AssetID, nil
}
