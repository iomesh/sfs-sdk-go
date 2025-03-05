OEM ?= iomesh.io

.PHONY: all
all: 
	OEM=iomesh.io make update-codegen
	OEM=arcfra.com make update-codegen

.PHONY: update-codegen
update-codegen:
	OEM=$(OEM) go generate ./...
	iidfile=$$(mktemp /tmp/iid-XXXXXX) && \
	docker build --build-arg IMAGE_REGISTRY=$(if $(IMAGE_REGISTRY),$(IMAGE_REGISTRY)/,) -f hack/Dockerfile --iidfile $$iidfile . && \
	docker run --rm -v $$PWD:/go/src/github.com/iomesh/sfs-sdk-go -w /go/src/github.com/iomesh/sfs-sdk-go $$(cat $$iidfile) hack/update-codegen.sh $(OEM) && \
	rm $$iidfile

.PHONY: verify-codegen
verify-codegen:
	iidfile=$$(mktemp /tmp/iid-XXXXXX) && \
	docker build --build-arg IMAGE_REGISTRY=$(if $(IMAGE_REGISTRY),$(IMAGE_REGISTRY)/,) -f hack/Dockerfile --iidfile $$iidfile . && \
	docker run --rm -v $$PWD:/go/src/github.com/iomesh/sfs-sdk-go -w /go/src/github.com/iomesh/sfs-sdk-go $$(cat $$iidfile) hack/verify-codegen.sh $(OEM) && \
	rm $$iidfile

.PHONY: clean
clean:
	OEM=iomesh.io make clean-codegen
	OEM=arcfra.com make clean-codegen
	
.PHONY: clean-codegen
clean-codegen:
	hack/clean-codegen.sh $(OEM)

.PHONY: verify
verify:
	OEM=iomesh.io make verify-codegen
	OEM=arcfra.com make verify-codegen