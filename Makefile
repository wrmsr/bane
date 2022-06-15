.PHONY: check-nodev
check-nodev:
	for f in $$((find cmd pkg -name 'dev.go' ; find pkg/dev -name '*.go') | sort | uniq) ; do \
		echo "$$f" ; \
	done
