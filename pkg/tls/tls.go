// Copyright 2020 IBM Corp.
// SPDX-License-Identifier: Apache-2.0

package tls

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/cert"
	kclient "sigs.k8s.io/controller-runtime/pkg/client"

	"fybrik.io/fybrik/pkg/environment"
)

const (
	TLSEnabledMsg   string = "TLS authentication is enabled"
	TLSDisabledMsg  string = "TLS authentication is disabled"
	MTLSEnabledMsg  string = "Mutual TLS authentication is enabled"
	MTLSDisabledMsg string = "Mutual TLS authentication is disabled"
)

// GetCertificatesFromSecret reads the certificates from kubernetes
// secret. Used when connection between manager and connectors uses tls.
func GetCertificatesFromSecret(client kclient.Client, secretName, secretNamespace string) (map[string][]byte, error) {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: secretNamespace,
			Name:      secretName,
		},
	}
	objectKey := kclient.ObjectKeyFromObject(secret)

	// Read the secret.
	err := client.Get(context.Background(), objectKey, secret)
	if err != nil {
		log.Error().Msg("Error reading cert secret data: " + err.Error())
		return nil, err
	}

	return secret.Data, nil
}

const (
	// TLSCertKeySuffix is the key suffix for tls certificates in a kubernetes secret.
	TLSCertKeySuffix = ".crt"
	TLSCertSuffix    = ".pem"
)

var certsDir = environment.GetDataDir() + "/tls-cert"
var cacertsDir = environment.GetDataDir() + "/tls-cacert"

var certFile = certsDir + "/" + corev1.TLSCertKey
var certPrivateKey = certsDir + "/" + corev1.TLSPrivateKeyKey

// GetServerConfig returns the server config for tls connection between the manager and
// the connectors.
func GetServerConfig(serverLog *zerolog.Logger, client kclient.Client) (*tls.Config, error) {
	// certSecretName is kubernetes secret name which contains the server certificate
	certSecretName := environment.GetCertSecretName()
	// certSecretNamespace is kubernetes secret namespace which contains the server certificate
	certSecretNamespace := environment.GetCertSecretNamespace()
	// caSecretName is kubernetes secret name which contains ca certificate used by the server to
	// validate certificate or the client.  Used when mutual tls is used.
	caSecretName := environment.GetCACERTSecretName()
	//	caSecretNamespace is  kubernetes secret namespace which contains ca certificate used by the server to
	// validate certificate or the client. Used when mutual tls is used.
	caSecretNamespace := environment.GetCACERTSecretNamespace()
	useMTLS := environment.IsUsingMTLS()

	if certSecretName == "" || certSecretNamespace == "" {
		// no server certs provided thus the tls is not used
		return nil, errors.New("no certificates provided")
	}
	serverLog.Info().Msg(TLSEnabledMsg)
	serverCertsData, err := GetCertificatesFromSecret(client, certSecretName, certSecretNamespace)
	if err != nil {
		serverLog.Error().Msg(err.Error())
		return nil, err
	}

	loadedCertServer, err := tls.X509KeyPair(serverCertsData[corev1.TLSCertKey], serverCertsData[corev1.TLSPrivateKeyKey])
	if err != nil {
		serverLog.Error().Msg(err.Error())
		return nil, err
	}
	var config *tls.Config
	if !useMTLS {
		serverLog.Info().Msg(MTLSDisabledMsg)
		config = &tls.Config{
			Certificates: []tls.Certificate{loadedCertServer},
			ClientAuth:   tls.NoClientCert,
			MinVersion:   tls.VersionTLS13,
		}
		return config, nil
	}
	serverLog.Info().Msg(MTLSEnabledMsg)
	CACertsData, err := GetCertificatesFromSecret(client, caSecretName, caSecretNamespace)
	if err != nil {
		return nil, err
	}
	CACertPool := x509.NewCertPool()
	for key, element := range CACertsData {
		// skip non cerificate keys like crt.key if exists in the secret
		if !strings.HasSuffix(key, TLSCertKeySuffix) {
			continue
		}
		if !CACertPool.AppendCertsFromPEM(element) {
			serverLog.Error().Err(err).Msg(err.Error())
			return nil, errors.New("error in GetServerConfig in AppendCertsFromPEM trying to lead key:" + key)
		}
	}

	config = &tls.Config{
		Certificates: []tls.Certificate{loadedCertServer},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    CACertPool,
		MinVersion:   tls.VersionTLS13,
	}

	return config, nil
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func find(root, ext string) ([]string, error) {
	var a []string
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

// GetClientTLSConfig returns the client config for tls connection between the manager and
// the connectors.
func GetClientTLSConfig(clientLog *zerolog.Logger, client kclient.Client) (*tls.Config, error) {
	// Mounted CA certificates
	caFiles, err := find(cacertsDir, ".pem")
	if err != nil {
		clientLog.Error().Msg(err.Error())
		return nil, err
	}
	var caCertPool *x509.CertPool
	if caFiles != nil {
		clientLog.Info().Msg("Found CA Certificates")
		caCertPool = x509.NewCertPool()
		for _, s := range caFiles {
			certs, e := cert.CertsFromFile(s)
			if e != nil {
				clientLog.Error().Msg(err.Error())
				return nil, e
			}
			for _, cert := range certs {
				caCertPool.AddCert(cert)
			}
		}
	} else {
		clientLog.Info().Msg("no CA certificates were found")
	}

	// Mounted manager certificates
	certFileExists := fileExists(certsDir + "/" + corev1.TLSCertKey)
	keyFileExists := fileExists(certsDir + "/" + corev1.TLSPrivateKeyKey)

	// If certificate file exists but not certificate key, or other way around, error out
	if (certFileExists && !keyFileExists) || (!certFileExists && keyFileExists) {
		return nil, errors.New("invalid SSL configuration found, please set both certificate file and certificate key file (one is missing)")
	}
	if !certFileExists && !keyFileExists {
		clientLog.Info().Msg("no certificates were found for the manager")
		if caFiles == nil {
			clientLog.Info().Msg("no TLS config is used")
			return nil, nil
		}
		clientLog.Info().Msg("TLS config contains manager certificates")
		tlsConfig := &tls.Config{
			RootCAs:    caCertPool,
			MinVersion: tls.VersionTLS13,
		}
		return tlsConfig, nil
	}

	clientLog.Info().Msg("found manager certificates")
	managerCerts, err := tls.LoadX509KeyPair(certFile, certPrivateKey)
	if err != nil {
		clientLog.Error().Err(err).Msg("error in X509KeyPair")
		return nil, err
	}
	clientLog.Info().Msg("TLS config contains manager certificates and CA certificates")
	tlsConfig := &tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{managerCerts},
		MinVersion:   tls.VersionTLS13,
	}

	return tlsConfig, nil
}
