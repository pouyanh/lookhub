#!/bin/sh

helm upgrade --install \
  --namespace="${NAMESPACE}" \
  --kube-apiserver="${KUBE_APISERVER}" \
  --kube-token="${KUBE_TOKEN}" \
  --set environment="${ENVIRONMENT}" \
  --set image.registry="${IMG_REGISTRY}" \
  --set image.repository="${IMG_REPOSITORY}" \
  --set image.tag="${IMG_TAG}" \
  "${CI_PROJECT_NAME}" "${HELM_CHART_DIR}"
