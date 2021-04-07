# Copyright 2020 IBM Corp.
# SPDX-License-Identifier: Apache-2.0
# This script is meant for local development with kind

export DOCKER_HOSTNAME=localhost:5000
export DOCKER_NAMESPACE=m4d-system
export HELM_EXPERIMENTAL_OCI=1

make kind
kubectl config use-context kind-control
make -C charts vault
make -C charts cert-manager
make -C third_party/datashim deploy
make docker
make cluster-prepare-wait
make configure-vault
make -C charts m4d
make -C manager wait_for_manager
make -C modules helm-chart-push
