#!/usr/bin/env bash


set -o errexit
set -o nounset
set -o pipefail

package_path="/Users/yuanmh/Desktop/code/k8s-crds-clientsets"

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
# ../vendor/k8s.io/code-generator/generate-groups.sh \
#   "deepcopy,client,informer,lister" \
#   k8s-crds-clientsets/generated \
#   k8s-crds-clientsets/apis \
#   compute.company.com:v1 \
#   --go-header-file $(pwd)/boilerplate.go.txt \
#   --output-base $(pwd)/../../

../vendor/k8s.io/code-generator/generate-groups.sh "deepcopy,client,informer,lister" \
../pkg/client \
../pkg/apis \
  nodehealth:v1alpha1 \
  --go-header-file ../hack/boilerplate.go.txt

