HELM_VALUES ?= \
	--set hello=world1

CHART := ${DOCKER_NAME}
HELM_TAGNAME ?= '0.0.0'
HELM_RELEASE ?= rel1-${DOCKER_NAME}
CHART_SCHEME ?= oci
TEMP := /tmp
CHART_PATH ?= ${CHART_SCHEME}://${DOCKER_HOSTNAME}/${DOCKER_NAMESPACE}/

export HELM_EXPERIMENTAL_OCI=1
export GODEBUG=x509ignoreCN=0

.PHONY: helm-login
helm-login: $(TOOLBIN)/helm
ifneq (${DOCKER_PASSWORD},)
	@$(ABSTOOLBIN)/helm registry login -u "${DOCKER_USERNAME}" -p "${DOCKER_PASSWORD}" ${DOCKER_HOSTNAME}
endif

.PHONY: helm-verify
helm-verify: $(TOOLBIN)/helm
	$(ABSTOOLBIN)/helm lint ../${CHART}
	$(ABSTOOLBIN)/helm install --dry-run ${HELM_RELEASE} ../${CHART} ${HELM_VALUES}

.PHONY: helm-uninstall
helm-uninstall: $(TOOLBIN)/helm
	$(ABSTOOLBIN)/helm uninstall ${HELM_RELEASE} || true

.PHONY: helm-install
helm-install: $(TOOLBIN)/helm
	$(ABSTOOLBIN)/helm install ${HELM_RELEASE} ../${CHART} ${HELM_VALUES}

.PHONY: helm-chart-push
helm-chart-push: helm-login $(TOOLBIN)/helm
	$(ABSTOOLBIN)/helm package ../${CHART} --version=${HELM_TAGNAME} --destination=${TEMP}
	$(ABSTOOLBIN)/helm push ${TEMP}/${CHART}-${HELM_TAGNAME}.tgz ${CHART_PATH}
	rm -rf ${TEMP}/${CHART}-${HELM_TAGNAME}.tgz

.PHONY: helm-chart-pull
helm-chart-pull: helm-login $(TOOLBIN)/helm
	$(ABSTOOLBIN)/helm pull ${CHART_PATH}/${CHART} --version ${HELM_TAGNAME}

.PHONY: helm-chart-list
helm-chart-list: $(TOOLBIN)/helm
	$(ABSTOOLBIN)/helm chart list

.PHONY: helm-chart-install
helm-chart-install: $(TOOLBIN)/helm
	$(ABSTOOLBIN)/helm install ${HELM_RELEASE} ${CHART_PATH}/${CHART} --version ${HELM_TAGNAME} ${HELM_VALUES}
	$(ABSTOOLBIN)/helm list

.PHONY: helm-template
helm-template: $(TOOLBIN)/helm
	$(ABSTOOLBIN)/helm template ${HELM_RELEASE} ../${CHART} ${HELM_VALUES}

.PHONY: helm-debug
helm-debug: $(ABSTOOLBIN)/helm
	$(ABSTOOLBIN)/helm template ${HELM_RELEASE} ../${CHART} ${HELM_VALUES} --debug

.PHONY: helm-actions
helm-actions:
	$(ABSTOOLBIN)/helm show values ../${CHART} | yq -y -r .actions

.PHONY: helm-all
helm-all: helm-verify helm-chart-push helm-chart-pull helm-uninstall helm-chart-install

