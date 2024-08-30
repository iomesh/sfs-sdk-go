OEM ?= iomesh.io

.PHONY: all
all: 
	OEM=iomesh.io make update-codegen
	OEM=arcfra.com make update-codegen

.PHONY: update-codegen
update-codegen:
	OEM=$(OEM) go generate ./...
	hack/update-codegen.sh $(OEM)

.PHONY: verify-codegen
verify-codegen:
	hack/verify-codegen.sh $(OEM)

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