#!/bin/sh

set -x

BUILD_COMMIT="$(git rev-list -1 HEAD)"
BUILD_BRANCH="$(git rev-parse --abbrev-ref HEAD)"
BUILD_VERSION="$(git describe --tags --always)"
BUILD_DATE="$(TZ=Asia/Tehran date '+%F %T%:z')"

cat <<EOF >"${BUILD_ENV_FILE}"
BUILD_COMMIT=${BUILD_COMMIT}
BUILD_BRANCH=${BUILD_BRANCH}
BUILD_VERSION=${BUILD_VERSION}
BUILD_DATE=${BUILD_DATE}
EOF

LDFLAGS="-w -s \
  -X '${CI_SERVER_HOST}/${CI_PROJECT_PATH}/config.Commit=${BUILD_COMMIT}'\
  -X '${CI_SERVER_HOST}/${CI_PROJECT_PATH}/config.Branch=${BUILD_BRANCH}'\
  -X '${CI_SERVER_HOST}/${CI_PROJECT_PATH}/config.Version=${BUILD_VERSION}'\
  -X '${CI_SERVER_HOST}/${CI_PROJECT_PATH}/config.BuildDate=${BUILD_DATE}'"

GOOS=linux GOARH=amd64 CGO_ENABLED=0 \
	go build -a -ldflags "${LDFLAGS}" -o .dist/lookhub ./cmd/server/...
