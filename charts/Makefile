ROOT_DIR := ..
include $(ROOT_DIR)/Makefile.env

WITH_OPENSHIFT ?= false
VALUES_FILE ?= "m4d/values.yaml"

.PHONY: kind
kind: export ROOT_DIR := .
kind:
	cd .. && $(MAKE) kind

.PHONY: cert-manager
cert-manager:
	@echo "Installing cert-manager ..."
	kubectl create namespace cert-manager
	helm repo add jetstack https://charts.jetstack.io
	helm repo update
	helm install cert-manager jetstack/cert-manager --version v1.1.0 \
		--set installCRDs=true \
		--namespace cert-manager \
		--wait --timeout 120s

.PHONY: vault
vault:
	@echo "Installing vault ..."
	kubectl create namespace m4d-system || true
	helm repo add hashicorp https://helm.releases.hashicorp.com
	helm install vault hashicorp/vault --version 0.9.1 \
                --set "global.openshift=$(WITH_OPENSHIFT)" \
                --set "server.dev.enabled=true" \
		--values ../third_party/vault/vault-single-cluster/values.yaml \
		--namespace m4d-system \
		--wait --timeout 120s
	kubectl apply -f ../third_party/vault/vault-single-cluster/vault-rbac.yaml -n m4d-system || true

.PHONY:
vault-setup:
	$(MAKE) configure-vault

.PHONY: m4d
m4d:
	@echo "installing m4d ..."
	kubectl create namespace m4d-system || true
	helm install m4d-crd m4d-crd  \
		--namespace m4d-system --wait --timeout 120s
	helm install m4d m4d --values $(VALUES_FILE) \
		--namespace m4d-system --wait --timeout 120s

.PHONY: run-integration-tests
run-integration-tests: export ROOT_DIR := .
run-integration-tests: kind cert-manager vault vault-setup m4d
