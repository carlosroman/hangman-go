# hangman-go 

[![CircleCI](https://circleci.com/gh/carlosroman/hangman-go.svg?style=svg)](https://circleci.com/gh/carlosroman/hangman-go)[![Coverage Status](https://coveralls.io/repos/github/carlosroman/hangman-go/badge.svg?branch=master)](https://coveralls.io/github/carlosroman/hangman-go?branch=master)


## Installation

The project is self contained so the only thing you need to is clone the project

```
$ git clone https://github.com/carlosroman/hangman-go.git
```

## Building

Open a terminal in the project directory and first setup the Go environment by running:

```
$ source .gopath
```

This will set both `GOPATH` to `checkout_dir/code` and `GOBIN` to `checkout_dir/code/bin`. Then to build, run:

```
$ make build
```

The executable can then be found in `checkout_dir/code/bin`

## Runing tests

The easiest way to run the tests is by running:

```
$ make test
```
