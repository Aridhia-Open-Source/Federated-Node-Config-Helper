#!/bin/bash

APP_NAME=fn-config-helper

GOOS=windows GOARCH=amd64 go build -o artifacts/$APP_NAME.exe ./main.go
