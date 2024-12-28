#!/bin/sh

helm upgrade --install \
  --namespace="${NAMESPACE}" \
  --kube-apiserver="${KUBE_APISERVER}" \
  --kube-token="${KUBE_TOKEN}" \
  --set okd.host="${OKD_HOST}" \
  --set environment="${ENVIRONMENT}" \
  --set image.repository="${IMAGE_REPO}" \
  --set image.tag="${IMAGE_TAG}" \
  "${CI_PROJECT_NAME}" "${HELM_CHART_DIR}"
