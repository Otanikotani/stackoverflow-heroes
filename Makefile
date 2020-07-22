PROJECT_NAME := "stackoverflow-heroes"
PKG := "github.com/otanikotani/$(PROJECT_NAME)"
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all

all: build

zip: clean.zip build
	zip -j ${PROJECT_NAME}.zip build/*

tidy: dep
	@go mod tidy

dep: ## Get the dependencies
	@go mod download

lint: ## Lint Golang files
	@golint .

test: ## Run unittests
	@go test .

build: tidy ## Build the binary file
	@go build -i -o build/${PROJECT_NAME} $(PKG)

clean.zip:
	@rm -rf build/
	rm ${PROJECT_NAME}.zip || true

clean:
	@rm -rf build/
	rm ${PROJECT_NAME}.zip || true
