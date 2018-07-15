.DEFAULT_GOAL := test

.PHONY: lint test test-% build setup setup-% coveralls coverprofile-fix docker docker-%

NS ?= hangman
VERSION ?= latest

DOCKER ?= docker
DOCKER_COMPOSE_FILE := ./Docker/docker-compose.yml
DOCKER_COMPOSE := docker-compose -f $(DOCKER_COMPOSE_FILE)

lint:
	@golangci-lint \
	    run \
	    code/src/hangman/...

setup: setup-ginkgo setup-dep
	@echo "Setup done"

setup-ginkgo:
	@go get -v github.com/onsi/ginkgo/ginkgo

setup-golangci-lint:
	@curl -L -s https://github.com/golangci/golangci-lint/releases/download/v${GOLANGCI_LINT_VERSION}/golangci-lint-${GOLANGCI_LINT_VERSION}-linux-amd64.tar.gz -o /tmp/golangci-lint-linux-amd64.tar.gz
	@tar -xzf /tmp/golangci-lint-linux-amd64.tar.gz  --strip 1 --directory ./code/bin

setup-dep:
	@curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o ./code/bin/dep
	@chmod +x ./code/bin/dep

build:
	@go install \
        code/src/hangman/server/server.go

test-service:
	@go test \
		-race -v  \
		./code/src/hangman/services/...

test-server:
	@go test \
		-race -v  \
		./code/src/hangman/server/handlers/...

test-clean:
	@find code/src/hangman -path ./vendor -o -name '*.coverprofile' -print | xargs rm -f
	@rm -rf target
	@mkdir -p target

test: test-clean
	@ginkgo \
        -r \
        --randomizeAllSpecs \
        --randomizeSuites \
        --failOnPending \
        --cover \
        --coverprofile=hangman.coverprofile \
        --outputdir=target \
        --trace \
        --race \
        --compilers=2 \
        ./code/src/hangman/

coverprofile-fix:
	@sed -i '/mode: atomic/d' target/hangman.coverprofile
	@sed -i '1imode: atomic' target/hangman.coverprofile

coveralls: coverprofile-fix
	@goveralls \
	    -coverprofile=target/hangman.coverprofile \
	    -service circle-ci \
	    -repotoken $$COVERALLS_TOKEN

docker-build-server:
	@$(DOCKER) build \
	     -f ./Docker/Dockerfile.server \
	     -t $(NS)/server:$(VERSION) .
