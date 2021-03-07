// Copyright 2021 IBM Corp.
// SPDX-License-Identifier: Apache-2.0
package connector

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"

	"github.com/ibm/the-mesh-for-data/connectors/katalog/pkg/connector/utils"
	"github.com/ibm/the-mesh-for-data/connectors/katalog/pkg/taxonomy"
	connectors "github.com/ibm/the-mesh-for-data/pkg/connectors/protobuf"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kclient "sigs.k8s.io/controller-runtime/pkg/client"
)

type DataCredentialsService struct {
	connectors.UnimplementedDataCredentialServiceServer

	client kclient.Client
}

func (s *DataCredentialsService) GetCredentialsInfo(ctx context.Context, req *connectors.DatasetCredentialsRequest) (*connectors.DatasetCredentials, error) {
	namespace, name, err := utils.SplitNamespacedName(req.DatasetId)
	if err != nil {
		return nil, err
	}

	// Get the secret name from the asset
	asset, err := getAsset(ctx, s.client, namespace, name)
	if err != nil {
		return nil, err
	}

	secretName := asset.Spec.SecretRef.Name

	// Read the secret
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      secretName,
		},
	}
	objectKey, err := kclient.ObjectKeyFromObject(secret)
	if err != nil {
		return nil, err
	}
	err = s.client.Get(ctx, objectKey, secret)
	if err != nil {
		return nil, err
	}

	// Get the data fields as strings
	data := map[string]string{}
	for key, value := range secret.Data {
		data[key] = string(value)
	}

	// Decode into Authentication structure
	authn := &taxonomy.Authentication{}
	switch secret.Type {
	case corev1.SecretTypeOpaque:
		err = utils.DecodeToStruct(data, authn)
		if err != nil {
			return nil, errors.Wrap(err, "Invalid fields in Secret data")
		}
	case corev1.SecretTypeBasicAuth:
		username := data[corev1.BasicAuthUsernameKey]
		password := data[corev1.BasicAuthPasswordKey]
		authn.Username = &username
		authn.Password = &password
	default:
		// TODO(roee88): add SSHAuth and TLSAuth as in corev1.SecretType
		return nil, fmt.Errorf("unknown secret type %s", secret.Type)
	}

	// Map to current connectors API
	return &connectors.DatasetCredentials{
		DatasetId: req.DatasetId,
		Creds: &connectors.Credentials{
			AccessKey: utils.EmptyIfNil(authn.AccessKey),
			SecretKey: utils.EmptyIfNil(authn.SecretKey),
			ApiKey:    utils.EmptyIfNil(authn.ApiKey),
			Username:  utils.EmptyIfNil(authn.Username),
			Password:  utils.EmptyIfNil(authn.Password),
		}}, nil
}
