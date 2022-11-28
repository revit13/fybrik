#!/usr/bin/env bash
# Copyright 2021 IBM Corp.
# SPDX-License-Identifier: Apache-2.0

FYBRIK_NAMESPACE=fybrik-system

kubectl create ns ${FYBRIK_NAMESPACE}
# Apply certificates for tls authentication
kubectl -n ${FYBRIK_NAMESPACE} apply -f ca-certificate.yaml --wait
kubectl -n ${FYBRIK_NAMESPACE} apply -f openmetadata-connector-certificates.yaml --wait
kubectl -n ${FYBRIK_NAMESPACE} apply -f opa-server-certificates.yaml --wait
kubectl -n ${FYBRIK_NAMESPACE} apply -f opa-connector-certificates.yaml --wait
kubectl -n ${FYBRIK_NAMESPACE} apply -f manager-certificates.yaml --wait
kubectl -n ${FYBRIK_NAMESPACE} apply -f vault-certificates.yaml --wait
kubectl -n ${FYBRIK_NAMESPACE} apply -f arrow-flight-module-certificates.yaml --wait
kubectl -n ${FYBRIK_NAMESPACE} apply -f localhost-certificates.yaml --wait


while ! kubectl get secret test-tls-openmetadata-connector-certs --namespace fybrik-system; do echo "Waiting for my secret. CTRL-C to exit."; sleep 1; done
rm -f openmetadata.keystore.jks || true
kubectl get secret test-tls-openmetadata-connector-certs -n fybrik-system -o json | jq -r '.data."tls.crt"' | base64 -d > tls.crt
docker run -it -v ${PWD}:/local joostdecock/keytool -import -alias localhost -file /local/tls.crt -storepass test12 -keystore /local/openmetadata.keystore.jks -noprompt 
