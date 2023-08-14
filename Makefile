GOCMD=go

all: deps lint test

test: deps
	$(GOCMD) test -v -count=1 -cover ./...

deps:
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor

lint:
	golangci-lint run