#!/bin/sh

go vet ./... || exit

go fmt ./... || exit

# todo: golangci-lint run || exit
