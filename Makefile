GO=go
PYTHON=python3

SRC=cmd pkg

###

.PHONY: check-fmt
check-fmt:
	r=0 ; \
	if [ ! -z "$$(gofmt -l ${SRC})" ] ; then \
		r=1 ; \
	fi ; \
	gofmt ${SRC}

.PHONY: check-vet
check-vet:
	${GO} vet ./...

.PHONY: check-nodev
check-nodev:
	r=0 ; \
	for f in $$((find cmd pkg -name 'dev.go' ; find pkg/dev -name '*.go') | sort | uniq) ; do \
  		if [[ "$$(head $$f)" != "//go:build !nodev"* ]] ; then \
  		  	r=1; \
  		  	b=$$(cat "$$f") ; \
			(printf "//go:build !nodev\n\n" ; echo "$$b" ) > "$$f" ; \
	  	fi ; \
	done ; \
	if [ $$r -ne 0 ] ; then \
		exit 1 ; \
	fi
