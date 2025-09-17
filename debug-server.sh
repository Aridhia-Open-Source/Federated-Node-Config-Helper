#!/bin/bash

if [[ ! dlv ]]; then
    echo "delve missing, please installing it with"
    echo "go install github.com/go-delve/delve/cmd/dlv@latest"
    exit 1
fi

dlv debug --headless --listen=:2345 --api-version=2 --log
