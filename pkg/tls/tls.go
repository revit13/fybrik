// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package tls

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
	corev1 "k8s.io/api/core/v1"

	"fybrik.io/fybrik/pkg/environment"
	"fybrik.io/fybrik/pkg/utils"
)

const (
	TLSEnabledMsg   string = "TLS authentication is enabled"
	TLSDisabledMsg  string = "TLS authentication is disabled"
	MTLSEnabledMsg  string = "Mutual TLS authentication is enabled"
	MTLSDisabledMsg string = "Mutual TLS authentication is disabled"
)

var certsDir = environment.GetDataDir() + "/tls-cert"
var cacertsDir = environment.GetDataDir() + "/tls-cacert"

var certFile = certsDir + "/" + corev1.TLSCertKey
var certPrivateKeyFile = certsDir + "/" + corev1.TLSPrivateKeyKey
var CACertFileSuffix = ".crt"

// isCertificateProvided returns true if the certificate file and private key were provided as expected.
// Otherwise it returns false.
func isCertificateProvided(certFileExists, keyFileExists bool) (bool, error) {
	if certFileExists && keyFileExists {
		return true, nil
	}
	if certFileExists || keyFileExists {
		return false, errors.New("invalid SSL configuration found, " +
			"please set both certificate name and namespace (one is missing)")
	}
	return false, nil
}

// getCertificate returns a certificate for the server/client if such provided.
func getCertificate() (*tls.Certificate, error) {
	// Mounted cert files.
	certFileExists := utils.IsPathExists(certFile)
	keyFileExists := utils.IsPathExists(certPrivateKeyFile)

	certProvided, err := isCertificateProvided(certFileExists, keyFileExists)
	if err != nil {
		return nil, err
	}
	if !certProvided {
		return nil, nil
	}

	cert, err := tls.LoadX509KeyPair(certFile, certPrivateKeyFile)
	if err != nil {
		return nil, err
	}
	return &cert, nil
}

func find(root, ext string) ([]string, error) {
	var a []string
	dirExists := utils.IsPathExists(cacertsDir)
	if !dirExists {
		return nil, nil
	}
	err := filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return a, nil
}

// getCACertPool returns the CA certificate if such provided.
func getCACertPool() (*x509.CertPool, error) {
	var CACertPool *x509.CertPool
	var err error
	certFiles, err := find(cacertsDir, CACertFileSuffix)
	if err != nil {
		return nil, err
	}
	if certFiles != nil {
		CACertPool = x509.NewCertPool()
		for _, cacertFile := range certFiles {
			caCert, err := os.ReadFile(cacertFile)
			if err != nil {
				return nil, err
			}
			if !CACertPool.AppendCertsFromPEM(caCert) {
				return nil, errors.New("error in AppendCertsFromPEM trying to load CA certificate")
			}
		}
		return CACertPool, nil
	}
	return nil, nil
}

// GetServerConfig returns the server config for tls connection between the manager and
// the connectors.
func GetServerConfig(serverLog *zerolog.Logger) (*tls.Config, error) {
	useMTLS := environment.IsUsingMTLS()
	var err error

	loadedCertServer, err := getCertificate()
	if err != nil {
		serverLog.Error().Msg(err.Error())
		return nil, err
	} else if loadedCertServer == nil {
		return nil, errors.New("TLS certificate for server is missing")
	} else {
		serverLog.Log().Msg("TLS certificate for server was provided")
	}
	serverLog.Info().Msg(TLSEnabledMsg)
	var config *tls.Config
	if !useMTLS {
		serverLog.Info().Msg(MTLSDisabledMsg)
		//nolint:gosec // ignore G402: TLS MinVersion too low
		config = &tls.Config{
			Certificates: []tls.Certificate{*loadedCertServer},
			// Do not use mutual TLS
			ClientAuth: tls.NoClientCert,
			MinVersion: environment.GetMinTLSVersion(serverLog),
		}
		return config, nil
	}
	serverLog.Info().Msg(MTLSEnabledMsg)
	caCertPool, err := getCACertPool()
	if err != nil {
		return nil, err
	}
	if caCertPool != nil {
		serverLog.Log().Msg("private CA certificates were provided in GetServerConfig")
	} else {
		serverLog.Log().Msg("private CA certificates were not provided in GetServerConfig")
	}
	//nolint:gosec // ignore G402: TLS MinVersion too low
	config = &tls.Config{
		Certificates: []tls.Certificate{*loadedCertServer},
		// configure mutual TLS
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
		MinVersion: environment.GetMinTLSVersion(serverLog),
	}

	return config, nil
}

// GetClientTLSConfig returns the client config for tls connection between the manager and
// the connectors.
func GetClientTLSConfig(clientLog *zerolog.Logger) (*tls.Config, error) {
	caCertPool, err := getCACertPool()
	if err != nil {
		return nil, err
	} else if caCertPool != nil {
		clientLog.Log().Msg("private CA certificates were provided in GetClientTLSConfig")
	} else {
		clientLog.Log().Msg("private CA certificates were not provided in GetClientTLSConfig")
	}
	cert, err := getCertificate()
	if err != nil {
		clientLog.Error().Msg(err.Error())
		return nil, err
	}
	if cert == nil && caCertPool == nil {
		clientLog.Log().Msg("no TLS certificates were provided")
		return nil, nil
	}
	var providedCert tls.Certificate
	if cert != nil {
		clientLog.Log().Msg("client TLS certificates were provided")
		providedCert = *cert
	} else {
		clientLog.Log().Msg("client TLS certificates were not provided")
	}

	//nolint:gosec // ignore G402: TLS MinVersion too low
	tlsConfig := &tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{providedCert},
		MinVersion:   environment.GetMinTLSVersion(clientLog),
	}

	return tlsConfig, nil
}
