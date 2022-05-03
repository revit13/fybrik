// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-retryablehttp"

	"fybrik.io/fybrik/pkg/connectors/datacatalog/clients"
	"fybrik.io/fybrik/pkg/model/datacatalog"
	"fybrik.io/fybrik/pkg/model/policymanager"
)

const (
	headerCredentials = "X-Request-Cred" // #nosec G101 -- This is a false positive
	policyEndpoint    = "/v1/data/dataapi/authz/verdict"
)

type ConnectorController struct {
	OpaServerURL  string
	OpaClient     *retryablehttp.Client
	CatalogClient clients.DataCatalog
}

func NewConnectorController(opaServerURL string, catalogClient clients.DataCatalog) *ConnectorController {
	return &ConnectorController{
		OpaServerURL:  opaServerURL,
		OpaClient:     retryablehttp.NewClient(),
		CatalogClient: catalogClient,
	}
}

func (r *ConnectorController) GetPoliciesDecisions(c *gin.Context) {
	// Parse request
	var request policymanager.GetPolicyDecisionsRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Avoid calling data catalog if the resource metadata is not empty
	if !datacatalog.ResourceMetadataNotEmpty(request.Resource.Metadata) {
		// Get asset metadata from catalog connector
		requestToCatalog := &datacatalog.GetAssetRequest{
			AssetID:       request.Resource.ID,
			OperationType: datacatalog.READ,
		}
		assetInfo, err := r.CatalogClient.GetAssetInfo(requestToCatalog, c.GetHeader(headerCredentials))
		if err != nil {
			// TODO: better error propagation
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Enrich request with catalog information
		request.Resource.Metadata = &assetInfo.ResourceMetadata
	}

	// Add "input" hierarchy
	inputStruct := map[string]interface{}{"input": &request}
	// Marshal request as JSON
	requestBody, err := json.Marshal(&inputStruct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Send request to OPA
	endpoint := fmt.Sprintf("%s/%s", strings.TrimRight(r.OpaServerURL, "/"), strings.TrimLeft(policyEndpoint, "/"))
	responseFromOPA, err := r.OpaClient.Post(endpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Read response from OPA
	defer responseFromOPA.Body.Close()
	responseFromOPABody, err := io.ReadAll(responseFromOPA.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Handle errors from OPA
	if responseFromOPA.StatusCode != http.StatusOK {
		// TODO: better error handling for OPA errors
		c.JSON(responseFromOPA.StatusCode, gin.H{"error": string(responseFromOPABody)})
		return
	}

	// Unmarshal as GetPolicyDecisionsResponse for the sake of validation
	var response policymanager.GetPolicyDecisionsResponse
	if err := json.Unmarshal(responseFromOPABody, &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
