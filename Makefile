.PHONY: all
all: update-codegen

.PHONY: update-codegen
update-codegen:
	hack/update-codegen.sh

.PHONY: verify-codegen
verify-codegen:
	hack/verify-codegen.sh

.PHONY: clean-codegen
clean-codegen:
	hack/clean-codegen.sh
