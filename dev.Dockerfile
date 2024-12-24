FROM golang:1.23

MAINTAINER Pouyan Heyratpour <me@pouyan.dev>

ARG ACCESS_TOKEN

RUN apt-get update && \
    apt-get install -y protobuf-compiler golang-golang-x-tools rsync

RUN go install -v github.com/pouyanh/polywatch/cmd/polywatch@v1.2.0 && \
    go install -v github.com/go-delve/delve/cmd/dlv@latest

RUN git config --global url."https://git@gitlab.snapp.ir/".insteadOf "https://gitlab.snapp.ir/"
RUN printf "machine gitlab.snapp.ir\nlogin git\npassword ${ACCESS_TOKEN}\n" >> ~/.netrc

ENTRYPOINT ["/go/bin/polywatch"]
