.PHONY: test lint lint118

test:
	go test -coverprofile="coverage.txt" -covermode=atomic -p 1 ./...
	go test -fuzz=Fuzz -fuzztime 5s ./stringtools
	go test -fuzz=Fuzz -fuzztime 5s ./contexts

lint:
	golangci-lint run -c .golangci.yml

lint118:
	golangci-lint run -c .golangci.yml --disable gocritic