exe = wordle-solver
PREFIX ?= /usr
dest = $(DESTDIR)$(PREFIX)/bin/$(exe)
INSTALL = install

.PHONY: all clobber run test install install-strip uninstall

all: $(exe)

$(exe): *.go
	go build -o $(exe) .

clobber:
	rm -vf $(exe)

run:
	go run .

test:
	go test -v .

$(dest): $(exe)
	$(INSTALL) -vD $(exe) $(dest)

install: $(dest)

install-strip:
	$(MAKE) INSTALL='$(INSTALL) -s' install

uninstall:
	rm -vf $(dest)
