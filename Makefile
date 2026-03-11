sources := $(shell find . -name '*.go') go.mod
exe = wordle-solver
exe-api = wordle-solver-api
PREFIX ?= /usr
dest = $(DESTDIR)$(PREFIX)/bin/$(exe)
INSTALL = install

.PHONY: all clean clobber run test install install-strip uninstall

all: $(exe) $(exe-api)

$(exe): $(sources)
	go build -o $(exe) ./cli

$(exe-api): $(sources)
	go build -o $(exe-api) ./api

clean:
	rm -vf cover.*

clobber: clean
	rm -vf $(exe) $(exe-api)

run:
	go run ./cli

cover.out: $(sources)
	go test -v -coverprofile cover.out -coverpkg ./...

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
