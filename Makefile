.PHONY: *
SHELL = /bin/bash

tests_ci:
	docker rmi --force test_go_cfg
	docker build -f tests.Dockerfile . -t test_go_cfg
	docker run --rm test_go_cfg
	docker rmi test_go_cfg

tests:
	go test ./...

coverage:
	go test -coverprofile=artifacts/coverage.out ./...
	go tool cover -html=artifacts/coverage.out -o artifacts/coverage.html

debug_server:
	./scripts/debug-server.sh

build_deb:
	./scripts/build_deb.sh

build_win:
	./scripts/build_exe.sh
