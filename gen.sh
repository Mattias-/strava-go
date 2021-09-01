#!/bin/bash
set -euo pipefail

SWAGGER_URL="https://developers.strava.com/swagger/swagger.json"

main() {
    ensure_go_swagger "v0.27.0"
    rm_gen
    gen
}

ensure_go_swagger() {
    local version=$1
    if ! [ -f ./swagger ]; then
        curl -sL -o ./swagger "https://github.com/go-swagger/go-swagger/releases/download/$version/swagger_linux_amd64"
        chmod +x ./swagger
    fi
}

gen() {
    if ! [ -f ./go.mod ]; then
        go mod init github.com/Mattias-/strava-go
    fi

    curl -sL -o ./swagger.json "$SWAGGER_URL"

    ./swagger generate client \
        -f "./swagger.json" \
        --skip-validation

    go get -u ./...
    go mod tidy
}

rm_gen() {
    rm -rf go.mod go.sum swagger.json models/ client/
}

main "$@"
