#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

CONTROLLER_GEN_PKG="sigs.k8s.io/controller-tools/cmd/controller-gen"
CONTROLLER_GEN_VER="v0.7.0"

# This function installs a Go tools by 'go install' command.
# Parameters:
#  - $1: package name, such as "sigs.k8s.io/controller-tools/cmd/controller-gen"
#  - $2: package version, such as "v0.4.1"
function util::install_tools() {
	local package="$1"
	local version="$2"
	echo "go install ${package}@${version}"
	GO111MODULE=on go install "${package}"@"${version}"
	GOPATH=$(go env GOPATH | awk -F ':' '{print $1}')
	export PATH=$PATH:$GOPATH/bin
}

echo "Generating crd with controller-gen"
util::install_tools ${CONTROLLER_GEN_PKG} ${CONTROLLER_GEN_VER} >/dev/null 2>&1

# Unify the crds used by helm chart and the installation scripts
controller-gen crd:allowDangerousTypes=true paths=./apis/resource...  output:crd:dir=./deploy/helm/aisp/crds

echo "Generating client with controller-gen"
SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo "${GOPATH}"/src/code-generator)}

bash "${CODEGEN_PKG}"/generate-groups.sh "all" \
  ai.plugin/scheduler/pkg/client  ai.plugin/scheduler/apis \
  resource:v1alpha1 \
  --output-base "../../" \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt





