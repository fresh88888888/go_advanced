APP=app

all: fmt test build

.PHONY: run
run: build

.PHONY: doc
doc:
	godoc -http=:6060 -index

.PHONY: build
build:
	GOARCH=amd64 GOOS=darwin go build -o $(APP)-darwin -v ./...
	GOARCH=amd64 GOOS=linux go build -o $(APP)-linux   -v ./...
	GOARCH=amd64 GOOS=windows go build -o $(APP)-windows -v ./...

.PHONY: test
test:
	go test -timeout 20s -race -v ./...

.PHONY: test_coverage
test_coverage:
	go test ./... -coverprofile=coverage.out

.PHONY: lint
lint:
	golangci-lint run --enable-all

.PHONY: debs
debs:
	go get ./...

.PHONY: vet
vet:
	go vet

.PHONY: fmt
fmt:
	gofmt -l -s .

.PHONY: clean
clean:
	go clean
	rm -f $(APP)-darwin
	rm -f $(APP)-linux
	rm -f $(APP)-windows