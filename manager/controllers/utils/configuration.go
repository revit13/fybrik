// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/onsi/ginkgo"
)

// Attributes that are defined in a config map or the runtime environment
const (
	CatalogConnectorServiceAddressKey string = "CATALOG_CONNECTOR_URL"
	VaultAddressKey                   string = "VAULT_ADDRESS"
	// Current communication between the manager and Vault assumes a Vault token
	// which is passed as an env var.
	// This should be changed to the following:
	// 1. get a Vault token form Vault using the manager role
	// 2. continue communication with Vault using this Vault token
	// The manager role should be passed as env var after it was configured
	// with the appropriate permissions for the manager.
	VaultSecretKey   string = "VAULT_TOKEN"
	VaultTTLKey      string = "VAULT_TTL"
	VaultModulesRole string = "VAULT_MODULES_ROLE"
)

// GetSystemNamespace returns the namespace of control plane
func GetSystemNamespace() string {
	if data, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace"); err == nil {
		if ns := strings.TrimSpace(string(data)); len(ns) > 0 {
			return ns
		}
	}
	return "m4d-system"
}

// GetModulesRole returns the modules assigned authentification role for accessing dataset credentials
func GetModulesRole() string {
	return os.Getenv(VaultModulesRole)
}

// GetVaultAddress returns the address and port of the vault system,
// which is used for managing data set credentials
func GetVaultAddress() string {
	return os.Getenv(VaultAddressKey)
}

// GetVaultToken returns the token this module uses to authenticate with vault
func GetVaultToken() string {
	return os.Getenv(VaultSecretKey)
}

// GetVaultAuthTTL returns the amount of time the authorization issued by vault is valid
func GetVaultAuthTTL() string {
	return os.Getenv(VaultTTLKey)
}

// GetDataCatalogServiceAddress returns the address where data catalog is running
func GetDataCatalogServiceAddress() string {
	return os.Getenv(CatalogConnectorServiceAddressKey)
}

func SetIfNotSet(key string, value string, t ginkgo.GinkgoTInterface) {
	if _, b := os.LookupEnv(key); !b {
		if err := os.Setenv(key, value); err != nil {
			t.Fatalf("Could not set environment variable %s", key)
		}
	}
}

func DefaultTestConfiguration(t ginkgo.GinkgoTInterface) {
	SetIfNotSet(CatalogConnectorServiceAddressKey, "localhost:50085", t)
	SetIfNotSet(VaultAddressKey, "http://127.0.0.1:8200/", t)
	SetIfNotSet("RUN_WITHOUT_VAULT", "1", t)
	SetIfNotSet("ENABLE_WEBHOOKS", "false", t)
	SetIfNotSet("CONNECTION_TIMEOUT", "120", t)
	SetIfNotSet("MAIN_POLICY_MANAGER_CONNECTOR_URL", "localhost:50090", t)
	SetIfNotSet("MAIN_POLICY_MANAGER_NAME", "MOCK", t)
	SetIfNotSet("USE_EXTENSIONPOLICY_MANAGER", "false", t)
}
