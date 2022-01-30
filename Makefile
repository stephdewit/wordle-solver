exe = wordle-solver

.PHONY: all clobber run test

all: $(exe)

$(exe): *.go
	go build -o $(exe) *.go

clobber:
	rm -vf $(exe)

run: $(exe)
	./$(exe)

test:
	go test -v *.go
