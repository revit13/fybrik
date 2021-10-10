export ROOT_DIR ?= ../..
export TOOLS_DIR := $(ROOT_DIR)/hack/tools
export TOOLBIN := $(TOOLS_DIR)/bin
export ABSTOOLBIN := $(shell pwd)/$(TOOLS_DIR)/bin

export HELM_EXPERIMENTAL_OCI=1

.PHONY: test
test: $(ABSTOOLBIN)/helm
	$(eval TMP := $(shell mktemp -d))
	cd $(TMP) && $(ABSTOOLBIN)/helm create mychart
	cd $(TMP) && $(ABSTOOLBIN)/helm package mychart -d $(TMP)  --version 0.7.0
	cd $(TMP) $(ABSTOOLBIN)/helm push $(TMP)/mychart-0.7.0.tgz oci://localhost:5000/fybrik-system/
	rm -rf $(TMP)/mychart-0.7.0.tgz
	TMP_CHART=$(TMP)/mychart go test -v reference.go helm.go helm_test.go
