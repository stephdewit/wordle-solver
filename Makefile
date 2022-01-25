exe=wordle-solver

.PHONY: all clobber run

all: $(exe)

$(exe): *.go
	go build -o $(exe) *.go

clobber:
	rm -vf $(exe)

run:
	go run *.go
