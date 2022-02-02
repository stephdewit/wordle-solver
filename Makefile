exe = wordle-solver
PREFIX ?= /usr
dest = $(DESTDIR)$(PREFIX)/bin/$(exe)
INSTALL = install

.PHONY: all clobber run test install install-strip uninstall

all: $(exe)

$(exe): *.go
	go build -o $(exe) *.go

clobber:
	rm -vf $(exe)

run: $(exe)
	./$(exe)

test:
	go test -v *.go

$(dest): $(exe)
	$(INSTALL) -vD $(exe) $(dest)

install: $(dest)

install-strip:
	$(MAKE) INSTALL='$(INSTALL) -s' install

uninstall:
	rm -vf $(dest)
