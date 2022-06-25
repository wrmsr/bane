GO=go

MOD:=$(shell grep -o '^module .*' go.mod | awk '{print $$2}')

SRC=\
	cmd \
	pkg \

BREW_DEPS=\


### clean

.PHONY: clean
clean:
	rm -rf \
		.cache \
		bin \

	ds=$$(find ${SRC} -name parser -type d) && \
	for d in "$$ds" ; do rm -rf "$$d" ; done


### deps

.PHONY: dep-update
dep-update:
	${GO} get -u ./...
	${GO} mod tidy


### gen

.PHONY: gen
gen: gen-antlr gen-go

ANTLR_VERSION=4.10.1

.PHONY: gen-antlr
gen-antlr:
	af="antlr-${ANTLR_VERSION}-complete.jar" ; \
	mkdir -p .cache ; \
	afp=".cache/$$af" ; \
	if [ ! -f "$$afp" ] ; then \
  		curl "https://www.antlr.org/download/$$af" -o "$$afp" ; \
  	fi ; \
	mtime() { \
		if [[ $$(uname) = "Darwin" ]] ; \
		then stat -f %m "$$1" ; else stat -c %Y "$$1" ; fi \
	} ; \
  	wd="$$(pwd)" ; \
  	for d in $$(find ${SRC} -name '*.g4' -exec dirname \{\} \; | sort | uniq) ; do \
  	  	fs=$$(find "$$d" -name '*.g4' | sort) ; \
  	  	nf=0 ; \
  	  	for f in $$fs ; do \
			mt=$$(mtime "$$f") ; \
			bf="$${f%.*}" ; \
			f2="$$(dirname $$bf)/parser/$$(basename $$bf).interp" ; \
			if [ ! -f "$$f2" ] ; then \
				nf=1 ; \
			else \
				mt2=$$(mtime "$$f2") ; \
				if [ "$$mt" -gt "$$mt2" ] ; then nf=1 ; fi ; \
			fi ; \
  	  	done ; \
		if [ $$nf -ne 0 ] ; then \
			rm -rf "$$d/parser" ; \
			for f in $$fs ; do \
				echo "$$f" ; \
				( \
					cd $$(dirname "$$f") && \
					java \
						-jar "$$wd/$$afp" \
						-Dlanguage=Go \
						$$(basename "$$f") \
						-visitor \
						-o parser \
				) ; \
			done ; \
		fi ; \
	done

.PHONY: gen-go
gen-go:
	${GO} generate ./...


### check

.PHONY: check
check: check-dev check-fmt check-vet

.PHONY: check-dev
check-dev:
	${GO} run "${MOD}/pkg/util/dev/cmd/checkdev" ${SRC}
	r=0 ; \
	for f in $$(( \
		find ${SRC} -name 'dev.go' -or -name '*_dev.go' ; \
		(find ${SRC} -name '*.go' | grep '/dev/') ; \
	) | sort | uniq) ; do \
  		if [[ "$$(head $$f)" != "//go:build !nodev"* ]] ; then \
  		  	r=1; \
			t=$$(mktemp) ; cat "$$f" > "$$t" ; \
			(printf "//go:build !nodev\n\n" ; cat "$$t" ) > "$$f" ; \
	  	fi ; \
	done ; \
	if [ $$r -ne 0 ] ; then exit 1 ; fi

.PHONY: check-fmt
check-fmt:
	r=0 ; \
	if [ ! -z "$$(gofmt -s -l ${SRC})" ] ; then r=1 ; fi ; \
	gofmt -s -w ${SRC} ; \
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
