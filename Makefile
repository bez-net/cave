VERSION=1.0

cave: *.go **/*.go
	GO111MODULE=on go build -ldflags '-X main.version=${VERSION}' -o bin/$@

.PHONY: all
all: cave

darwin-build: *.go **/*.go
	GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -ldflags '-X main.version=${VERSION}' -o bin/cave-darwin

linux-build: *.go **/*.go
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -ldflags '-s -w -X main.version=${VERSION}' -o bin/cave-linux

check:
	GO111MODULE=on go test ./...

fmt:
	go fmt ./...

clean:
	rm -rf cave