#!/usr/bin/env bash
# Copyright 2021 IBM Corp.
# SPDX-License-Identifier: Apache-2.0

set -x

kubectl delete namespace fybrik-notebook-sample || true
NS="fybrik-system"; kubectl -n $NS get configmap | awk '/sample/{print $1}' | xargs  kubectl delete -n $NS configmap

if [[ "${USE_OPENMETADATA_CATALOG}" -eq 1 ]]; then
  port=8080
  local_port=8081
  prefix=http

  if [[ "${DEPLOY_TLS_TEST_CERTS}" -eq 1 ]]; then
    port=8443
    prefix=https
    local_port=8443
    kubectl get secret test-tls-localhost-certs -n fybrik-system -o json | jq -r '.data."tls.key"' | base64 -d > tls.key
    kubectl get secret test-tls-localhost-certs -n fybrik-system -o json | jq -r '.data."tls.crt"' | base64 -d > tls.crt
    certs=" --cert tls.crt --key tls.key "
  fi
  # Deploy openmetadata asset
  kubectl port-forward svc/openmetadata-connector -n fybrik-system $local_port:$port &
  CATALOGED_ASSET="openmetadata-s3.default.bucket1.\"data.csv\""
  # Wait until curl command succeed
  c=0
  # -k flag is used to skip server verification to avoid errors regarding target host name 'localhost'
  while [[ $(curl $certs -k -X POST $prefix:localhost:${local_port}/getAssetInfo -d '{"assetID": ${CATALOGED_ASSET}, "operationType": "read"}' -H "Content-type: application/json"   -H "X-Request-Datacatalog-Cred: QQQ") != *'assetID'* ]]
  do
    echo "waiting for curl command to createAsset to succeed"
    ((c++)) && ((c==25)) && break
    sleep 1
  done
fi

aws --endpoint-url=http://localhost:9090 s3api --bucket=bucket1 delete-objects --delete='{"Objects": [{"Key": "data.csv"}]}'
