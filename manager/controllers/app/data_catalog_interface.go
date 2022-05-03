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
// - assetID: DataSetID as it appears in fybrik-application
// - catalogID: the destination catalog identifier
// - resourceMetadataFromRequirements: resource metadata from the flow requirements
// - info: connection and credential details
// Returns:
// - an error if happened
// - the new asset identifier
func (r *FybrikApplicationReconciler) RegisterAsset(assetID string, catalogID string,
	resourceMetadataFromRequirements *datacatalog.ResourceMetadata, info *app.DatasetDetails,
	input *app.FybrikApplication) (string, error) {
	r.Log.Trace().Msg("RegisterAsset")
	details := datacatalog.ResourceDetails{}
	if info.Details != nil {
		details.Connection = info.Details.Connection
		details.DataFormat = info.Details.Format
	}

	var resourceMetadata datacatalog.ResourceMetadata
	if resourceMetadataFromRequirements != nil {
		resourceMetadata = *resourceMetadataFromRequirements.DeepCopy()
	} else {
		resourceMetadata.Name = assetID
	}
	// Update the Geography with the allocated storage region
	if info.ResourceMetadata != nil {
		resourceMetadata.Geography = info.ResourceMetadata.Geography
	}

	creds := ""
	if utils.IsVaultEnabled() {
		creds = utils.GetVaultAddress() + vault.PathForReadingKubeSecret(info.SecretRef.Namespace, info.SecretRef.Name)
	}

	request := datacatalog.CreateAssetRequest{
		ResourceMetadata:     *resourceMetadata.DeepCopy(),
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
