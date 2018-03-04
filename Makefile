.PHONY: test build

build:
	go install \
        code/src/hangman/server/server.go

test-service:
	go test \
		-race -v  \
		code/src/hangman/services/*.go

test-server:
	go test \
		-race -v  \
		code/src/hangman/server/handlers/*.go

test: test-service test-server
