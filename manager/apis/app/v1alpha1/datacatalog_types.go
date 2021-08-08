// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import "fybrik.io/fybrik/pkg/serde"

// Credentials holds the credentials for accessing the data
// Credentials for accesing the data are stored in Vault, in the location represented by Vault property.
type Credentials struct {
	// +required
	Vault Vault `json:"vault"`
}

// Credentials holds the credentials for accessing the data for read, write and copy operation types
type DataStoreCredentials struct {
	// Read holds the credentials for reading the asset
	// +optional
	Read *Credentials `json:"read,omitempty"`

	// Write holds the credentials for writing the asset
	// +optional
	Write *Credentials `json:"write,omitempty"`

	// Copy holds the credentials for copying the asset
	// +optional
	Copy *Credentials `json:"copy,omitempty"`
}

// DataStore contains the details for accesing the data that are sent by catalog connectors
type DataStore struct {
	// +required
	Credentials DataStoreCredentials `json:"credentials"`
	// Connection has the relevant details for accesing the data (url, table, ssl, etc.)
	// +required
	Connection serde.Arbitrary `json:"connection"`
	// Format represents data format (e.g. parquet) as received from catalog connectors
	// +required
	Format string `json:"format"`
}
