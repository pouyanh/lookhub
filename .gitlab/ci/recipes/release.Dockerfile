FROM alpine:latest

ARG LBL_TITLE=CoreX
ARG LBL_DESC="Pouyanh X Service"
ARG LBL_URL=https://github.com/pouyanh/svcx
ARG BUILD_COMMIT
ARG BUILD_VERSION
ARG BUILD_DATE

LABEL org.opencontainers.image.authors="Pouyan Heyratpour <me@pouyan.dev>" \
    org.opencontainers.image.vendor="Pattack" \
    org.opencontainers.image.title="${LBL_TITLE}" \
    org.opencontainers.image.description="${LBL_DESC}" \
    org.opencontainers.image.url="${LBL_URL}" \
    org.opencontainers.image.source="${LBL_URL}" \
    org.opencontainers.image.revision="${BUILD_COMMIT}" \
    org.opencontainers.image.version="${BUILD_VERSION}" \
    org.opencontainers.image.created="${BUILD_DATE}"

COPY ./.dist /bin
WORKDIR /

ENV TZ="Asia/Tehran"
ENTRYPOINT ["/bin/lookhub"]
HEALTHCHECK --interval=30s --timeout=10s --retries=5 \
    CMD wget --no-verbose --tries=1 --spider http://localhost/healthz:30080
