### project

GO=go

MOD:=$$(go list -m)

SUBMODS:=$$(find . -name go.mod | sed -n 's/^\.\/\(.*\)\/go\.mod$$/\1/p')

MSRCS:=$$( \
	go list ./... | \
	awk ' \
		BEGIN{L=length("github.com/wrmsr/bane")+2} \
		{S=substr($$1,L); P=index(S,"/"); if(P>0){ print substr(S,1,P-1) } else { print S }} \
	' | sort | uniq \
)

SRCS=\
	core \
	sql \
	x \

MODS:=$$(echo ${SRCS} | xargs -n1 -I% echo ./%/...)

PKGS:=$$(echo ${SRCS} | xargs -n1 -I% ${GO} run ${MOD}/core/dev/cmd/list -find -ignorefiles ./%/...)


### echos

.PHONY: mod
mod:
	@echo ${MOD}

.PHONY: msrcs
msrcs:
	@echo ${MSRCS}

.PHONY: submods
submods:
	@echo ${SUBMODS}

.PHONY: srcs
srcs:
	@echo ${SRCS}

.PHONY: mods
mods:
	@echo ${MODS}

.PHONY: pkgs
pkgs:
	@echo ${PKGS}


### versions

define get-version
	$$(grep '^$(1)=' .versions | cut -d= -f2)
endef

GO_VERSION:=$(call get-version,'GO')
PYTHON_VERSION:=$(call get-version,'PYTHON')


### deps

BREW_DEPS=\
	graphviz \
	socat \
	sqlite3 \

PYTHON_DEPS=\
	ipython \
	llvmlite \
	numpy \
	pandas \
	pyyaml \
	torch \
	torchvision \
	\
	'git+https://github.com/wrmsr/tinygrad' \


### clean

.PHONY: clean
clean:
	rm -rf \
		.cache \
		.venv* \
		bin \

	${MAKE} clean-antlr

.PHONY: clean-antlr
clean-antlr:
	ds=$$(find ${SRCS} -name parser -type d) && \
	for d in $$ds ; do rm -rf "$$d" ; done

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
	# ${GO} get -u ./...
	# ${MAKE} dep-tidy
	# github.com/wrmsr/bane v0.0.0-00000000000000-000000000000
	# GOV0=$(curl -s 'https://go.dev/VERSION\?m\=text')
	# GOV1=$(go version | egrep -o 'go1\.[0-9]+(\.[0-9]+)?')
	# curl -s 'https://www.python.org/ftp/python/' | egrep -o 'href="3\.[0-9]+\.[0-9]+/"' | tr -dC '0-9.\n'
	for D in . sql x ; do (cd "$$D" && go get -u ./... && go mod tidy) ; done && \
	find . -name go.mod | xargs sed -i '' 's/\(github.com\/wrmsr\/bane\) v.*/\1 v0.0.0-00000000000000-000000000000/g'


.PHONY: dep-tidy
dep-tidy:
	${GO} mod tidy


### gen

.PHONY: gen
gen: gen-antlr gen-go

ANTLR_VERSION=4.12.0

.PHONY: gen-antlr
gen-antlr:
	${GO} run "${MOD}/core/antlr/dev/gen" ${ANTLR_VERSION} ${SRCS}

.PHONY: gen-go
gen-go:
	${GO} generate ./...
	${GO} run "${MOD}/core/dev/cmd/checkdev" -q ${SRCS}


### check

.PHONY: check
check: check-dev check-fmt check-vet

.PHONY: check-dev
check-dev:
	${GO} run "${MOD}/core/dev/cmd/checkdev" ${SRCS}

.PHONY: check-fmt
check-fmt:
	r=0 ; \
	if [ ! -z "$$(${GO} run cmd/gofmt -s -l ${SRCS})" ] ; then r=1 ; fi ; \
	${GO} run cmd/gofmt -s -w ${SRCS} ; \
	if [ $$r -ne 0 ] ; then exit 1 ; fi

.PHONY: check-vet
check-vet:
	${GO} vet -composites=false ${PKGS}


### test

.PHONY: test
test:
	${GO} test ./...

.PHONY: test-verbose
test-verbose:
	${GO} test -v ./...


### docker

DOCKER_COMPOSE=docker-compose -f docker/docker-compose-ext.yml -f docker/docker-compose.yml

.PHONY: docker-stop
docker-stop:
	${DOCKER_COMPOSE} stop

.PHONY: docker-reup
docker-reup: docker-stop
	${DOCKER_COMPOSE} rm -f
	${DOCKER_COMPOSE} build bane-dev
	${DOCKER_COMPOSE} up

.PHONY: docker-invalidate
docker-invalidate:
	date +%s > docker/.dockertimestamp

.PHONY: docker-enable-ptrace
docker-enable-ptrace:
	docker run --platform linux/x86_64 --privileged -it ubuntu sh -c 'echo 0 > /proc/sys/kernel/yama/ptrace_scope'


### ci

.PHONY: ci
ci:
	${DOCKER_COMPOSE} build bane-ci
	${DOCKER_COMPOSE} run \
		--rm \
		$$BANE_CI_DOCKER_OPTS \
		-e BANE_CI_OUTPUT_DIR="$$BANE_CI_OUTPUT_DIR" \
		bane-ci \
		make _ci

.PHONY: _ci
_ci:
	# FIXME: make test
	echo hi


### utils

.PHONY: imports
imports:
	@${GO} run "${MOD}/core/dev/cmd/imports" ${SRCS}


### python

PYENV_ROOT:=$$(sh -c "if [ -z '$${PYENV_ROOT}' ] ; then echo '$${HOME}/.pyenv' ; else echo '$${PYENV_ROOT%/}' ; fi")
PYENV_BIN:=$$(sh -c "if [ -f '$${HOME}/.pyenv/bin/pyenv' ] ; then echo '$${HOME}/.pyenv/bin/pyenv' ; else echo pyenv ; fi")

PYTHON=.venv/bin/python

.PHONY: py-venv
py-venv:
	if [ ! -d .venv ] ; then \
		$(PYENV_BIN) install -s $(PYTHON_VERSION) && \
		"$(PYENV_ROOT)/versions/$(PYTHON_VERSION)/bin/python" -mvenv .venv && \
		${PYTHON} -mpip install --upgrade pip setuptools wheel ; \
	fi

.PHONY: py-deps
py-deps: py-venv
	${PYTHON} -mpip install ${PYTHON_DEPS}
