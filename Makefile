.PHONY: test build

build:
	go install \
        code/src/hangman/server/server.go

test-service:
	go test \
		-race -v  \
		./code/src/hangman/services/...

test-server:
	go test \
		-race -v  \
		./code/src/hangman/server/handlers/...

test:
	ginkgo \
        -r \
        --randomizeAllSpecs \
        --randomizeSuites \
        --failOnPending \
        --cover \
        --trace \
        --race \
        --compilers=2 \
        ./code/src/hangman/
