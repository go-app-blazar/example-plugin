all: generate format build

clean:
	rm -vf example/*_generated.go
	rm -fv internal/base/example/*_generated.go

generate:
	go generate ./...

format:
	go fmt ./...

test:
	go test ./...

build:
	go build ./...