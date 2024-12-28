#!/bin/sh

CGO_ENABLED=1 \
	go test -count=1 -covermode=atomic -race -coverprofile=coverage.out -test.short ./...
go tool cover -func coverage.out
