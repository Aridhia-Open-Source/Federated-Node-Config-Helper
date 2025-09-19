#!/bin/bash

APP_NAME=fn-config-helper

GOOS=windows GOARCH=amd64 go build -o artifacts/$APP_NAME.exe ./main.go

zip -j artifacts/fn-config-helper-win-amd64.zip artifacts/fn-config-helper.exe README.md
rm artifacts/fn-config-helper.exe
