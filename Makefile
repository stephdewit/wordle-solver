sources-solver := $(shell find solver/ -name '*.go') go.mod
sources-cli    := $(shell find cli/ -name '*.go') $(sources-solver)
sources-api    := $(shell find api/ -name '*.go') $(sources-solver)
sources        := $(sources-cli) $(sources-api)
exe = wordle-solver
exe-api = $(exe)-api
PREFIX ?= /usr
dest = $(DESTDIR)$(PREFIX)/bin/$(exe)
INSTALL = install

namespace = stephdewit
image-api = $(exe)
image-ui  = $(exe)-ui

.PHONY: all clean clobber run test install install-strip uninstall build-api-image build-ui-image push-api push-ui check-version

all: $(exe) $(exe-api)

$(exe): $(sources-cli)
	go build -o $(exe) ./cli

$(exe-api): $(sources-api)
	go build -o $(exe-api) ./api

clean:
	rm -vf cover.*

clobber: clean
	rm -vf $(exe) $(exe-api)

run:
	go run ./cli

cover.out: $(sources)
	go test -v -coverprofile cover.out -coverpkg ./... ./...

cover.html: cover.out
	go tool cover -html=cover.out -o cover.html

test: cover.html

$(dest): $(exe)
	$(INSTALL) -vD $(exe) $(dest)

install: $(dest)

install-strip:
	$(MAKE) INSTALL='$(INSTALL) -s' install

uninstall:
	rm -vf $(dest)

build-api-image: check-version
	docker build . \
		-t $(namespace)/$(image-api):$(VERSION) \
		-t $(namespace)/$(image-api):latest

build-ui-image: check-version
	docker build ui/ \
		-t $(namespace)/$(image-ui):$(VERSION) \
		-t $(namespace)/$(image-ui):latest

push-api: check-version
	docker push $(namespace)/$(image-api):$(VERSION)
	docker push $(namespace)/$(image-api):latest

push-ui: check-version
	docker push $(namespace)/$(image-ui):$(VERSION)
	docker push $(namespace)/$(image-ui):latest

check-version:
ifndef VERSION
	$(error VERSION is undefined)
endif
