exe = wordle-solver
PREFIX ?= /usr
dest = $(DESTDIR)$(PREFIX)/bin/$(exe)

.PHONY: all clobber run test install uninstall

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
	install -vDs $(exe) $(dest)

install: $(dest)

uninstall:
	rm -vf $(dest)
