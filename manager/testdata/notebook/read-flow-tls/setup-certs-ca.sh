#!/usr/bin/env bash
# Copyright 2021 IBM Corp.
# SPDX-License-Identifier: Apache-2.0

ROOT_DIR := ../../../../

# Apply certificates for tls authentication
kubectl -n fybrik-system apply -f ca-certificate.yaml
kubectl -n fybrik-system apply -f katalog-connector-certificates.yaml
kubectl -n fybrik-system apply -f opa-connector-certificates.yaml
kubectl -n fybrik-system apply -f manager-certificates.yaml

kubectl get secret test-tls-ca-certs -o json | jq -r '.data."ca.crt"' | base64 -d > /tmp/ca.crt

cp katalog-connector-Dockerfile ${ROOT_DIR}/connectors/katalog/Dockerfile
cp opa-connector-Dockerfile ${ROOT_DIR}/connectors/opa/Dockerfile
cp manager-Dockerfile ${ROOT_DIR}/manager/Dockerfile





