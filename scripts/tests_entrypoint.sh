#!/bin/bash

go test -v -coverprofile=artifacts/coverage.out ./...
go tool cover -func=artifacts/coverage.out
