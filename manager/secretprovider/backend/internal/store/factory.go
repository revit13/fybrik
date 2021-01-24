/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package store

import (
	"context"

	"github.com/go-logr/logr"

	smv1alpha1 "github.com/ibm/the-mesh-for-data/manager/secretprovider/backend/apis/secretmanager/v1alpha1"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Client is a common interface for interacting with SecretStore
// backends
type Client interface {
	GetSecret(ctx context.Context, ref smv1alpha1.RemoteReference) ([]byte, error)
	GetSecretMap(ctx context.Context, ref smv1alpha1.RemoteReference) (map[string][]byte, error)
}

// Factory returns a StoreClient
type Factory interface {
	New(ctx context.Context, log logr.Logger, store smv1alpha1.GenericStore, kubeClient client.Client, kubeReader client.Reader, namespace string) (Client, error)
}
