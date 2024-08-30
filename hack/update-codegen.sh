#!/usr/bin/env bash

# Copyright 2023 The IOMesh Authors.

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

apigroup="$1"

MODULE="github.com/iomesh/sfs-sdk-go"
INPUT_PKG_ROOT="${MODULE}/pkg/apis/${apigroup}"
OUTPUT_PKG_ROOT="${MODULE}/pkg/client/${apigroup}"
SCRIPT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="${SCRIPT_ROOT}/.."
OUTPUT_BASE="${SCRIPT_ROOT}/../../../.."
BOILERPLATE="${SCRIPT_ROOT}/boilerplate.go.txt"

CODEGEN_PKG="${CODEGEN_PKG:-$(
    cd ${REPO_ROOT}
    ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || go list -m -f "{{.Dir}}" k8s.io/code-generator
)}"
source "${CODEGEN_PKG}/kube_codegen.sh"

kube::codegen::gen_helpers \
    --input-pkg-root "${INPUT_PKG_ROOT}" \
    --output-base "${OUTPUT_BASE}" \
    --boilerplate "${BOILERPLATE}"

mkdir -p "${OUTPUT_BASE}/${OUTPUT_PKG_ROOT}/clientset"
mkdir -p "${OUTPUT_BASE}/${OUTPUT_PKG_ROOT}/listers"
mkdir -p "${OUTPUT_BASE}/${OUTPUT_PKG_ROOT}/informers"

kube::codegen::gen_client \
    --with-watch \
    --input-pkg-root "${INPUT_PKG_ROOT}" \
    --output-pkg-root "${OUTPUT_PKG_ROOT}" \
    --output-base "${OUTPUT_BASE}" \
    --boilerplate "${BOILERPLATE}"
