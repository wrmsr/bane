GO=go

MOD:=$(shell grep -o '^module .*' go.mod | awk '{print $$2}')

SRC=\
	cmd \
	pkg \

BREW_DEPS=\
	graphviz \
	socat \
	sqlite3 \


### clean

.PHONY: clean
clean:
	rm -rf \
		.cache \
		bin \

	ds=$$(find ${SRC} -name parser -type d) && \
	for d in "$$ds" ; do rm -rf "$$d" ; done


### sys

.PHONY: brew
brew:
	brew install ${BREW_DEPS}


### deps

.PHONY: deps
deps:
	${GO} get ./...

.PHONY: dep-update
dep-update:
	${GO} get -u ./...
	${MAKE} dep-tidy

.PHONY: dep-tidy
dep-tidy:
	${GO} mod tidy


### gen

.PHONY: gen
gen: gen-antlr gen-go

ANTLR_VERSION=4.10.1

.PHONY: gen-antlr
gen-antlr:
	${GO} run "${MOD}/pkg/util/antlr/dev/gen" ${ANTLR_VERSION} ${SRC}

.PHONY: gen-go
gen-go:
	${GO} generate ./...
	${GO} run "${MOD}/pkg/util/dev/cmd/checkdev" -q ${SRC}


### check

.PHONY: check
check: check-dev check-fmt check-vet

.PHONY: check-dev
check-dev:
	${GO} run "${MOD}/pkg/util/dev/cmd/checkdev" ${SRC}

.PHONY: check-fmt
check-fmt:
	r=0 ; \
	if [ ! -z "$$(${GO} run cmd/gofmt -s -l ${SRC})" ] ; then r=1 ; fi ; \
	${GO} run cmd/gofmt -s -w ${SRC} ; \
	if [ $$r -ne 0 ] ; then exit 1 ; fi

.PHONY: check-vet
check-vet:
	${GO} vet -composites=false ./...


### test

.PHONY: test
test:
	go test ./...

.PHONY: test-verbose
test-verbose:
	go test -v ./...


### build

.PHONY: build
build:
	${GO} build \
		-tags nodev \
		-o bin/bane \
		"${MOD}/cmd/bane"


### docker

DOCKER_COMPOSE=docker-compose -f docker/docker-compose.yml

.PHONY: docker-stop
docker-stop:
	${DOCKER_COMPOSE} stop

.PHONY: docker-reup
docker-reup: docker-stop
	${DOCKER_COMPOSE} rm -f
	${DOCKER_COMPOSE} up

.PHONY: docker-invalidate
docker-invalidate:
	date +%s > docker/.dockertimestamp
	docker rmi -f docker_bane-dev


### utils

.PHONY: imports
imports:
	${GO} run "${MOD}/pkg/util/dev/cmd/imports" ${SRC}
