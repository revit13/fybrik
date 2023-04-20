ROOT_DIR := ../../..
include $(ROOT_DIR)/Makefile.env
include $(ROOT_DIR)/hack/make-rules/tools.mk

USER_ID := $(shell id -u)
GROUP_ID := $(shell id -g)

GIT_HOST := fybrik.io
GIT_USER_ID := argocd
GIT_REPO_ID := argocd-go
GIT_REPO_ID_MODELS := argocd-go-models
GIT_REPO_ID_CLIENT := argocd-go-client

AUTO_GENERATED = auto-generated
FYBRIK_VERSION ?= v1.2.1

IGNORE := $(shell bash -c "sed -n /=/p  ${ROOT_DIR}/hack/tools/requirements.env | sed 's/=/:=/' | sed 's/^/export /' > makeenv")
include makeenv

ABS_ROOT_DIR := $(abspath ${ROOT_DIR})
OPENAPI_GENERATOR_IMG := openapitools/openapi-generator-cli:v${OPENAPI_GENERATOR_VERSION}

.PHONY: generate-client-argocd
generate-client-argocd:
	rm -Rf ${AUTO_GENERATED}/
	docker run --rm \
		-v ${PWD}:/local \
		-u "${USER_ID}:${GROUP_ID}" \
		${OPENAPI_GENERATOR_IMG} generate -g go \
		--additional-properties=packageName=openapiclient,isGoSubmodule=false \
		--global-property=apis,models,supportingFiles,apiDocs=false \
		--git-host=${GIT_HOST} \
		--git-user-id=${GIT_USER_ID} \
		--git-repo-id=${GIT_REPO_ID_MODELS} \
		-o /local/${AUTO_GENERATED}/client \
		-i /local/swagger.json
	rm -f $(ROOT_DIR)/pkg/multicluster/argocd/auto-generated/client/go.mod $(ROOT_DIR)/pkg/multicluster/argocd/auto-generated/client/go.sum $(ROOT_DIR)/pkg/multicluster/argocd/auto-generated/client/.travis.yml
	rm -f $(ROOT_DIR)/pkg/multicluster/argocd/auto-generated/client/README.md $(ROOT_DIR)/pkg/multicluster/argocd/auto-generated/client/git_push.sh
	rm -f $(ROOT_DIR)/pkg/multicluster/argocd/auto-generated/client/.gitignore $(ROOT_DIR)/pkg/multicluster/argocd/auto-generated/client/.openapi-generator-ignore
	rm -r $(ROOT_DIR)/pkg/multicluster/argocd/auto-generated/client/api


.PHONY: check
check:
	go fmt ./...
	go vet ./...
	go mod tidy -compat=1.19
	docker run --rm -v ${PWD}:/app -w /app golangci/golangci-lint:v1.50.1 golangci-lint run

clean:
	rm makeenv
