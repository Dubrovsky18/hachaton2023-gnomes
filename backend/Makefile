PROJECT_PKG = github.com/Dubrovsky18/hachaton2023-gnomes
BUILD_DIR = build
VERSION ?=$(shell git describe --tags --exact-match 2>/dev/null || git symbolic-ref -q --short HEAD)


run:
	go run -v ./cmd/app/main.go serve