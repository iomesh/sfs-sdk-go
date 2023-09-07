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

GROUP="sfs.iomesh.io"
VERSION="v1"
SCRIPT_ROOT="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
REPO_ROOT="$(dirname SCRIPT_ROOT)"
INPUT_PKG_ROOT="${REPO_ROOT}/pkg/apis"
OUTPUT_PKG_ROOT="${REPO_ROOT}/pkg/client"
DEEPCOPY_FILE_NAME="zz_generated.deepcopy.go"

echo "remove deepcopy file"
rm -f "${INPUT_PKG_ROOT}/${GROUP}/${VERSION}/${DEEPCOPY_FILE_NAME}"

echo "remove client"
rm -rf "${OUTPUT_PKG_ROOT}"
