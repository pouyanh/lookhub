#!/bin/sh

helm upgrade --install \
  --namespace="${KUBE_NS}" \
  --kube-apiserver="${KUBE_ENDPOINT}" \
  --kube-token="${KUBE_TOKEN}" \
  --set environment="${ENVIRONMENT}" \
  --set image.registry="${IMG_REGISTRY}" \
  --set image.repository="${IMG_REPOSITORY}" \
  --set image.tag="${IMG_TAG}" \
  "${CI_PROJECT_NAME}" "${HELM_CHART_DIR}"
