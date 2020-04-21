# Makefile to build the command lines and tests in Seele project.
# This Makefile doesn't consider Windows Environment. If you use it in Windows, please be careful.

SHELL := /bin/bash

#BASEDIR = $(shell pwd)
BASEDIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

DEFAULT_PLUGINS_DIR = ${BASEDIR}/plugins
DEFAULT_PLUGINS_NAME = plugins_demo.so
HOSTNAME = 10.14.41.51
HOSTPORT = 3000
NAMESPACE = bar

ldflagsDebug=""
# -s -w
ldflagsRelase="-s -w"


#go run <name> -h <host> -p <port> -n <namespace> -s <set>

.PHONY: test

default: test

test:
	go run examples/touch.go -h ${HOSTNAME} -p ${HOSTPORT} -n ${NAMESPACE} -s demoset

benchmark:
	cd tools/benchmark && \
	go run benchmark.go -h ${HOSTNAME} -n ${NAMESPACE} -k 10000000 -w RU,80 -R -o S:50 -T 10 -c 20 -L 5,1

#-h 10.14.41.51 -n bar -k 10000000 -w RU,80 -R -o S:50 -T 10 -c 20 -L 5,1

build-darwin:
	mkdir -p "build/bin" && \
	export CGO_ENABLED=0 && export GOOS=darwin && export GOARCH=amd64 && \
	go build -v -ldflags ${ldflagsRelase} -o ./build/bin/benchmark-darwin ./tools/benchmark
	@echo "Done benchmark built for darwin"

build-linux:
	mkdir -p "build/bin" && \
	export CGO_ENABLED=0 && export GOOS=linux && export GOARCH=amd64 && \
	go build -v -ldflags ${ldflagsRelase} -o ./build/bin/benchmark-linux ./tools/benchmark
	@echo "Done benchmark built for linux"
