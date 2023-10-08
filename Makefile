COVER_FILE=coverage.out

all: deps test

test: deps
	go test -v -count=1 -coverprofile $(COVER_FILE) -cover ./...

deps:
	go mod tidy
	go mod vendor

bump-deps:
	go get -u ./...
	$(MAKE)

lint:
	golangci-lint run
