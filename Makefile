.PHONY: test lint lint118

test:
	go test ./...
	go test -fuzz=Fuzz -fuzztime 5s ./stringtools
	go test -fuzz=Fuzz -fuzztime 5s ./contexts

lint:
	golangci-lint run -c golangci-lint.yml

lint118:
	golangci-lint run -c golangci-lint.yml --disable gocritic