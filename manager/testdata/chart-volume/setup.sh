#!/usr/bin/env bash
# Copyright 2021 IBM Corp.
# SPDX-License-Identifier: Apache-2.0

set -x
set -e

export TOOLBIN=../../../hack/tools/bin
export MODULE_VERSION="0.0.0-master"

TMP=$(mktemp -d)
${TOOLBIN}/helm pull oci://ghcr.io/fybrik/arrow-flight-module-chart --version ${MODULE_VERSION} -d=${TMP}
cd ${TMP}
tar -xvf ${TMP}/arrow-flight-module-chart-0.0.0-master.tgz  

POD_NAME=$(kubectl get pods -n fybrik-system -o=name | sed "s/^.\{4\}//" |grep manager)
${TOOLBIN}/kubectl cp ${TMP}/arrow-flight-module-chart ${POD_NAME}:/opt/fybrik -n fybrik-system


