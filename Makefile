GO=go
PYTHON=python3

###

.PHONY: check-vet
check-vet:
	${GO} vet ./...

.PHONY: check-nodev
check-nodev:
	for f in $$((find cmd pkg -name 'dev.go' ; find pkg/dev -name '*.go') | sort | uniq) ; do \
  		if [[ "$$(head $$f)" != "//go:build !nodev"* ]] ; then \
			echo "$$f" ; \
	  	fi ; \
	done
