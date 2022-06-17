GO=go

SRC=\
	cmd \
	pkg \


### clean

.PHONY: clean
clean:
	rm -rf \
		antlr-*.jar \
		bin \

	ds=$$(find ${SRC} -name parser -type d) && \
	for d in "$$ds" ; do rm -rf "$$d" ; done

### gen

.PHONY: gen
gen: gen-antlr gen-go

ANTLR_VERSION=4.10.1

.PHONY: gen-antlr
gen-antlr:
	af="antlr-${ANTLR_VERSION}-complete.jar" ; \
	if [ ! -f "$$af" ] ; then \
  		curl "https://www.antlr.org/download/$$af" -o "$$af" ; \
  	fi ; \
  	wd="$$(pwd)" ; \
  	for d in $$(find ${SRC} -name '*.g4' -exec dirname \{\} \; | sort | uniq) ; do \
  	  	rm -rf "$$d/parser" ; \
		for f in $$(find "$$d" -name '*.g4' | sort) ; do \
			echo "$$f" ; \
			( \
				cd $$(dirname "$$f") && \
				java \
					-jar "$$wd/$$af" \
					-Dlanguage=Go \
					$$(basename "$$f") \
					-visitor \
					-o parser \
			) ; \
		done ; \
	done

.PHONY: gen-go
gen-go:
	${GO} generate ./...


### check

.PHONY: check
check: check-nodev check-fmt check-vet

.PHONY: check-nodev
check-nodev:
	r=0 ; \
	for f in $$(( \
		find ${SRC} -name 'dev.go' -or -name '*_dev.go' ; \
		(find ${SRC} -name '*.go' | grep '/dev/') ; \
	) | sort | uniq) ; do \
  		if [[ "$$(head $$f)" != "//go:build !nodev"* ]] ; then \
  		  	r=1; \
  		  	b=$$(cat "$$f") ; \
			(printf "//go:build !nodev\n\n" ; echo "$$b" ) > "$$f" ; \
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


### build

.PHONY: build
build:
	${GO} build \
		-tags nodev \
		-o bin/bane \
		cmd/bane/main.go
