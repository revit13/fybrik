// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package datacatalog

import (
	"errors"

	"fybrik.io/fybrik/pkg/model/taxonomy"
)

// OperationType Type of operation requested for the asset
// +kubebuilder:validation:Enum=read;
type OperationType string

// List of operationType
const (
	READ OperationType = "read"
)

// ResourceMetadata defines model for resource metadata
type ResourceMetadata struct {
	// Name of the resource
	Name string `json:"name,omitempty"`
	// Owner of the resource
	Owner string `json:"owner,omitempty"`
	// Geography of the resource
	Geography string `json:"geography,omitempty"`
	// Tags associated with the asset
	Tags *taxonomy.Tags `json:"tags,omitempty"`
	// Columns associated with the asset
	Columns []ResourceColumn `json:"columns,omitempty"`
}

// ResourceColumn represents a column in a tabular resource
type ResourceColumn struct {
	// Name of the column
	Name string `json:"name"`
	// Tags associated with the column
	Tags *taxonomy.Tags `json:"tags,omitempty"`
}

// ResourceDetails includes asset connection details
type ResourceDetails struct {
	// Connection information
	Connection taxonomy.Connection `json:"connection"`
	// Data format
	DataFormat taxonomy.DataFormat `json:"dataFormat,omitempty"`
}

// checkMetadataConflicts returns an error if there is a conflict in the values
// of meta1 and meta2
func checkMetadataConflicts(meta1, meta2 *ResourceMetadata) error {
	if meta1 == nil || meta2 == nil {
		return nil
	}
	if meta1.Tags != nil && meta2.Tags != nil {
		for k, v := range meta1.Tags.Items {
			if val, ok := meta2.Tags.Items[k]; ok {
				if val != v {
					return errors.New("conflicts in ResourceMetadata tags")
				}
			}
		}
	}
	if meta1.Geography != "" && meta2.Geography != "" && meta1.Geography != meta2.Geography {
		return errors.New("geography conflict in ResourceMetadata geography field")
	}
	if meta1.Owner != "" && meta2.Owner != "" && meta1.Owner != meta2.Owner {
		return errors.New("conflict in ResourceMetadata owner field")
	}
	if meta1.Name != "" && meta2.Name != "" && meta1.Name != meta2.Name {
		return errors.New("conflict in ResourceMetadata name field")
	}
	return nil
}

// MergeMetadata merges the metadata in newMetaData into intoMetadata
// Returns an error if there is a conflict
func MergeMetadata(intoMetadata, newMetaData *ResourceMetadata) error {
	if newMetaData == nil {
		return nil
	}
	if newMetaData.Columns != nil {
		intoMetadata.Columns = append(intoMetadata.Columns, newMetaData.Columns...)
	}
	if err := checkMetadataConflicts(intoMetadata, newMetaData); err != nil {
		return err
	}
	if newMetaData.Tags != nil {
		newTags := newMetaData.Tags.Items
		if intoMetadata.Tags == nil {
			intoMetadata.Tags = &taxonomy.Tags{}
			intoMetadata.Tags.Items = make(map[string]interface{})
		}
		oldTags := intoMetadata.Tags.Items
		for k, v := range newTags {
			oldTags[k] = v
		}
	}

	if newMetaData.Geography != "" {
		intoMetadata.Geography = newMetaData.Geography
	}
	if newMetaData.Owner != "" {
		intoMetadata.Owner = newMetaData.Owner
	}
	if newMetaData.Name != "" {
		intoMetadata.Name = newMetaData.Name
	}
	return nil
}

// Return true if resourceMetadata is not empty. Otherwise return false.
func ResourceMetadataEmpty(resourceMetadata *ResourceMetadata) bool {
	if resourceMetadata == nil {
		return true
	}
	if resourceMetadata.Columns != nil ||
		(resourceMetadata.Tags != nil &&
			len(resourceMetadata.Tags.Items) != 0) ||
		resourceMetadata.Geography != "" ||
		resourceMetadata.Name != "" ||
		resourceMetadata.Owner != "" {
		return false
	}
	return true
}
