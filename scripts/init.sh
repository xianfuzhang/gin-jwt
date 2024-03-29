#!/bin/bash

OUT_DIR=${OUT_DIR:-_output}

function setup_env() {
    current_path="$(dirname ${BASH_SOURCE})/.."
    APP_ROOT=$(realpath ${current_path})
    APP_OUT=${APP_ROOT}/${OUT_DIR}
}

function start_build() {
    pushd "${APP_ROOT}" >/dev/null
    mkdir -p "${APP_OUT}"

    CGO_ENABLED=0 go build -v -o "${APP_OUT}" "./$1"

    popd >/dev/null
}

readonly -f setup_env
readonly -f start_build

setup_env
start_build $@