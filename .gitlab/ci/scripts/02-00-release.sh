#!/bin/sh

set -x

docker login -u "${CI_REGISTRY_USER}" -p "${CI_REGISTRY_PASSWORD}" "${CI_REGISTRY}"
docker build -f .gitlab/ci/recipes/release.Dockerfile \
	--build-arg LBL_TITLE="${CI_PROJECT_TITLE}" \
	--build-arg LBL_DESC="${CI_PROJECT_DESCRIPTION}" \
	--build-arg LBL_URL="${CI_PROJECT_URL}" \
	--build-arg BUILD_COMMIT="${BUILD_COMMIT}" \
	--build-arg BUILD_VERSION="${BUILD_VERSION}" \
	--build-arg BUILD_DATE="${BUILD_DATE}" \
	--no-cache \
	-t "${IMG}" .

docker login -u "${IMG_PUSH_USERNAME}" -p "${IMG_PUSH_PASSWORD}" "${IMG_REGISTRY}"
docker push "${IMG}"
