current_time = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
git_description = $(shell git describe --always --dirty --tags --long)
linker_flags = '-s -X main.buildTime=${current_time} -X main.version=${git_description}'

dev-bundler:
	go run ./main.go bundler

.PHONY: build
build:
	GIN_MODE=release go build -ldflags=${linker_flags} -o=./build/main ./main.go

format:
	go fmt ./