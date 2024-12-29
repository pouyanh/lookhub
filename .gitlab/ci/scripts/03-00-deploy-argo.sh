#!/bin/sh

argocd login "${ARGO_ENDPOINT}" --auth-token "${ARGO_TOKEN}" --insecure
argocd app sync "${CI_PROJECT_NAME}"
argocd app wait "${CI_PROJECT_NAME}" --health --timeout 300
