include Makefile.env
export DOCKER_TAGNAME ?= latest

.PHONY: license
license: $(TOOLBIN)/license_finder
	$(call license_go,.)
	$(call license_python,secret-provider)

.PHONY: docker-mirror-read
docker-mirror-read:
	$(TOOLS_DIR)/docker_mirror.sh $(TOOLS_DIR)/docker_mirror.conf

.PHONY: build
build:
	$(MAKE) -C pkg/policy-compiler build
	$(MAKE) -C manager manager

.PHONY: test
test:
	$(MAKE) -C pkg/policy-compiler test
	$(MAKE) -C manager test

.PHONY: run-integration-tests
run-integration-tests: export DOCKER_HOSTNAME?=kind-registry:5000
run-integration-tests: export DOCKER_NAMESPACE?=m4d-system
run-integration-tests:
	$(MAKE) kind
	$(MAKE) cluster-prepare
	$(MAKE) docker
	$(MAKE) -C test/services docker-all
	$(MAKE) cluster-prepare-wait
	$(MAKE) -C secret-provider configure-vault
	$(MAKE) -C secret-provider deploy
	$(MAKE) -C manager deploy-crd
	$(MAKE) -C manager deploy_it
	$(MAKE) -C manager wait_for_manager
	$(MAKE) helm
	$(MAKE) -C manager run-integration-tests

.PHONY: run-deploy-tests
run-deploy-tests: export KUBE_NAMESPACE?=m4d-system
run-deploy-tests:
	$(MAKE) kind
	$(MAKE) cluster-prepare
	kubectl config set-context --current --namespace=$(KUBE_NAMESPACE)
	$(MAKE) -C third_party/opa deploy
	kubectl apply -f ./manager/config/prod/deployment_configmap.yaml
	kubectl create secret generic user-vault-unseal-keys --from-literal=vault-root=$(kubectl get secrets vault-unseal-keys -o jsonpath={.data.vault-root} | base64 --decode) 
	$(MAKE) -C connectors deploy
	kubectl get pod --all-namespaces
	kubectl wait --for=condition=ready pod --all-namespaces --all --timeout=120s

.PHONY: cluster-prepare
cluster-prepare:
	$(MAKE) -C third_party/cert-manager deploy
	$(MAKE) -C third_party/registry deploy
	$(MAKE) -C third_party/vault deploy

.PHONY: cluster-prepare-wait
cluster-prepare-wait:
	$(MAKE) -C third_party/cert-manager deploy-wait
	$(MAKE) -C third_party/vault deploy-wait

.PHONY: install
install:
	$(MAKE) -C manager install

.PHONY: deploy
deploy:
	$(MAKE) -C secret-provider deploy
	$(MAKE) -C manager deploy
	$(MAKE) -C connectors deploy

.PHONY: undeploy
undeploy:
	$(MAKE) -C secret-provider undeploy
	$(MAKE) -C manager undeploy
	$(MAKE) -C manager undeploy-crd
	$(MAKE) -C connectors undeploy

.PHONY: docker
docker:
	$(MAKE) -C manager docker-all
	$(MAKE) -C secret-provider docker-all
	$(MAKE) -C connectors docker-all
	$(MAKE) -C test/dummy-mover docker-all

# Build only the docker images needed for integration testing
.PHONY: docker-minimal-it
docker-minimal-it:
	$(MAKE) -C manager docker-all
	$(MAKE) -C secret-provider docker-all
	$(MAKE) -C test/dummy-mover docker-all
	$(MAKE) -C test/services docker-all

.PHONY: docker-build
docker-build:
	$(MAKE) -C manager docker-build
	$(MAKE) -C secret-provider docker-build
	$(MAKE) -C connectors docker-build
	$(MAKE) -C test/dummy-mover docker-build

.PHONY: docker-push
docker-push:
	$(MAKE) -C manager docker-push
	$(MAKE) -C secret-provider docker-push
	$(MAKE) -C connectors docker-push
	$(MAKE) -C test/dummy-mover docker-push

.PHONY: helm
helm:
	$(MAKE) -C modules helm

DOCKER_PUBLIC_HOSTNAME ?= ghcr.io
DOCKER_PUBLIC_NAMESPACE ?= the-mesh-for-data
DOCKER_PUBLIC_NAMES := \
	manager \
	secret-provider \
	egr-connector \
	dummy-mover \
	opa-connector \
	vault-connector
 
define do-docker-retag-and-push-public
	for name in ${DOCKER_PUBLIC_NAMES}; do \
		docker tag ${DOCKER_HOSTNAME}/${DOCKER_NAMESPACE}/$$name:${DOCKER_TAGNAME} ${DOCKER_PUBLIC_HOSTNAME}/${DOCKER_PUBLIC_NAMESPACE}/$$name:$1; \
	done
	DOCKER_HOSTNAME=${DOCKER_PUBLIC_HOSTNAME} DOCKER_NAMESPACE=${DOCKER_PUBLIC_NAMESPACE} DOCKER_TAGNAME=$1 $(MAKE) docker-push
endef

.PHONY: docker-retag-and-push-public
docker-retag-and-push-public:
	$(call do-docker-retag-and-push-public,latest)
ifneq (${TRAVIS_TAG},)
	$(call do-docker-retag-and-push-public,${TRAVIS_TAG})
endif

.PHONY: helm-push-public
helm-push-public:
	DOCKER_HOSTNAME=${DOCKER_PUBLIC_HOSTNAME} DOCKER_NAMESPACE=${DOCKER_PUBLIC_NAMESPACE} make -C modules helm-chart-push
ifneq (${TRAVIS_TAG},)
	DOCKER_HOSTNAME=${DOCKER_PUBLIC_HOSTNAME} DOCKER_NAMESPACE=${DOCKER_PUBLIC_NAMESPACE} DOCKER_TAGNAME=${TRAVIS_TAG} make -C modules helm-chart-push
endif


include hack/make-rules/tools.mk
include hack/make-rules/verify.mk
include hack/make-rules/cluster.mk
