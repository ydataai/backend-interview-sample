-include .env
export $(shell sed 's/=.*//' .env)

GOPATH=$(shell go env GOPATH)

K = kubectl

.PHONY: help controller-gen build generate manifests fmt vet

CRD_OPTIONS ?= "crd:trivialVersions=false"

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

help:	# The following lines will print the available commands when entering just 'make'. ⚠️ This needs to be the first target, ever
ifeq ($(UNAME), Linux)
	@grep -P '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
else
	@awk -F ':.*###' '$$0 ~ FS {printf "%15s%s\n", $$1 ":", $$2}' \
		$(MAKEFILE_LIST) | grep -v '@awk' | sort
endif

controller-gen:
ifeq (, $(shell which controller-gen))
	@{ \
	set -e ;\
	CONTROLLER_GEN_TMP_DIR=$$(mktemp -d) ;\
	cd $$CONTROLLER_GEN_TMP_DIR ;\
	go mod init tmp ;\
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.6.1 ;\
	rm -rf $$CONTROLLER_GEN_TMP_DIR ;\
	}
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

build:	### Build
	go build -a pkg/main.go pkg/configuration.go pkg/scheme.go

generate:	controller-gen	### Generate code
	$(CONTROLLER_GEN) object paths=./pkg/models/kube/v1alpha1/...

manifests:	controller-gen	### Generate manifests
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./pkg/models/kube/..." output:crd:artifacts:config=config/crd/bases

fmt:    ### Run go fmt against code
	go fmt ./...

vet:	### Run go vet against code
	go vet ./...

test:   ### Runs application's tests in verbose mode
	go test -v ./pkg/...
