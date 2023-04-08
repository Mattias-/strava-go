SHELL := /usr/bin/env bash -o errexit -o pipefail -o nounset
.SUFFIXES:

OS_NAME := $(shell uname -s | tr A-Z a-z)
ARCH := $(shell uname -m)
GO_SWAGGER_VERSION="v0.30.4"
GO_SWAGGER_URL=https://github.com/go-swagger/go-swagger/releases/download/$(GO_SWAGGER_VERSION)/swagger_$(OS_NAME)_$(ARCH)
SWAGGER_URL=https://developers.strava.com/swagger/swagger.json

generate: clean swagger.json swagger
	go mod init github.com/Mattias-/strava-go
	./swagger generate client \
        -f "./swagger.json" \
        --skip-validation
	go get -u ./...
	go mod tidy

.PHONY: clean
clean:
	$(RM) -rf go.mod go.sum swagger.json swagger models/ client/

swagger.json:
	curl -sL -o ./swagger.json "$(SWAGGER_URL)"

swagger:
	curl -sL -o ./swagger "https://github.com/go-swagger/go-swagger/releases/download/$(GO_SWAGGER_VERSION)/swagger_$(OS_NAME)_$(ARCH)"
	chmod +x ./swagger
