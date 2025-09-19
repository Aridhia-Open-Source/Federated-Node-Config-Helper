.PHONY: *
SHELL = /bin/bash

tests:
	go test ./...

coverage:
	go test -coverprofile=artifacts/coverage.out ./...
	go tool cover -html=artifacts/coverage.out -o artifacts/coverage.html

debug_server:
	./debug-server.sh
