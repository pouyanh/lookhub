#!/bin/sh

go env -w GOPROXY="${GO_PROXY}"
go env GOPROXY
go mod tidy
go mod verify
go mod vendor
