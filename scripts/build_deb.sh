#!/bin/bash

APP_NAME=fn-config-helper

mkdir -p build/deb/$APP_NAME/usr/local/bin/

GOOS=linux GOARCH=amd64 go build -o build/deb/$APP_NAME/usr/local/bin/$APP_NAME ./main.go
pushd build/deb
dpkg-deb --build $APP_NAME ../../artifacts
popd

rm -r build/deb/$APP_NAME/usr
