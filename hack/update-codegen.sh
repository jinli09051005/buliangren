#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

GOPATH=$(go env GOPATH)

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}
bash "${CODEGEN_PKG}"/generate-groups.sh "deepcopy,client,lister,informer" \
	cangbinggu.io/buliangren/pkg/generated \
	cangbinggu.io/buliangren/pkg/apis \
	tiankuixing:v1 \
	--output-base "${GOPATH}/src" \
	--go-header-file "${GOPATH}"/src/cangbinggu.io/buliangren/hack/boilerplate.go.txt

GOPATH="${GOPATH}" bash "${CODEGEN_PKG}"/generate-internal-groups.sh "deepcopy,client,lister,informer,openapi" \
	cangbinggu.io/buliangren/pkg/generated \
	cangbinggu.io/buliangren/pkg/apis \
	cangbinggu.io/buliangren/pkg/apis \
	tiankuixing:v1 \
	--output-base "${GOPATH}/src" \
	--go-header-file "${GOPATH}"/src/cangbinggu.io/buliangren/hack/boilerplate.go.txt
