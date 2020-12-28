#############################
# Global vars
#############################
PROJECT_NAME := $(shell basename $(shell pwd))
PROJECT_VER  ?= $(shell git describe --tags --always --dirty | sed -e '/^v/s/^v\(.*\)$$/\1/g')
# Last released version (not dirty)
PROJECT_VER_TAGGED  := $(shell git describe --tags --always --abbrev=0 | sed -e '/^v/s/^v\(.*\)$$/\1/g')

SRCDIR       ?= .
GO            = go

# The root module (from go.mod)
PROJECT_MODULE  ?= $(shell $(GO) list -m)

GEN_ARCH        ?= $(shell uname -i)
GEN_KERNEL      ?= $(shell uname -r)
GEN_KERNELNAME  ?= $(shell uname -s)



#############################
# Targets
#############################
all: build

build: clean compile

include build/compile.mk
include build/deps.mk
include build/util.mk
include build/release.mk

build-install:
	@echo "Preparing files in temporary directory"
	@rm -rf ./tmp ./installbundles
	@mkdir ./tmp ./installbundles
	@cp ./bin/${GOOS}/tsak ./tmp
	@cp ./install/setup.sh ./tmp
	@cp /usr/local/lib/libclips.so ./tmp
	@echo "Preparing installation bundle"
	@makeself ./tmp ./installbundles/tsak-install-${GOOS}-${GEN_KERNEL}.sh "TSAK installation for ${GOOS}" ./setup.sh

.PHONY: all build clean
