.PHONY: test

test:
	go test ./...
	go test -fuzz=Fuzz -fuzztime 5s ./stringtools

lint:
	golangci-lint run -c golangci-lint.yml

lint118:
	golangci-lint run -c golangci-lint.yml --disable gocritic