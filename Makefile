.PHONY: %

default: build

build: vet fmt lint
	go build ./...

vet:
	go vet ./...

doc:
	godoc -http=:6060

fmt:
	go fmt ./...

lint:
	golint ./...

test:
	go test -v ./...

update-deps:
	go get -u -v ./...
