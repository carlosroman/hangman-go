.DEFAULT_GOAL := test

.PHONY: test test-% build setup setup-% coveralls coverprofile-fix

setup: setup-ginkgo setup-dep
	@echo "Setup done"

setup-ginkgo:
	go get -v github.com/onsi/ginkgo/ginkgo

setup-dep:
	curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o ./code/bin/dep
	chmod +x ./code/bin/dep

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

test-clean:
	find code/src/hangman -path ./vendor -o -name '*.coverprofile' -print | xargs rm -f
	rm -rf target
	mkdir -p target

test: test-clean
	ginkgo \
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
	sed -i '/mode: atomic/d' target/hangman.coverprofile
	sed -i '1imode: atomic' target/hangman.coverprofile

coveralls: coverprofile-fix
	goveralls \
	    -coverprofile=target/hangman.coverprofile \
	    -service circle-ci \
	    -repotoken $$COVERALLS_TOKEN
