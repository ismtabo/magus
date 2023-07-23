PACKAGE=github.com/ismtabo/magus
BINARY_NAME=magus
TARGET_DIR=target
VERSION=0.0.0
ARCH=amd64
OS=linux darwin windows
outputs=$(foreach os,${OS},${TARGET_DIR}/${BINARY_NAME}-${os})
go_files=$(shell find . -type f -name '*.go' -not -path "./vendor/*")

os=$(subst ${TARGET_DIR}/${BINARY_NAME}-,,$@)
$(outputs): ${go_files}
	@mkdir -p ${TARGET_DIR}
	@echo "Building ${BINARY_NAME} for ${os}..." && \
	GOARCH=${ARCH} GOOS=${os} go build -ldflags "-X ${PACKAGE}/cmd.versionString=${VERSION}" -o $@ $<

.PHONY: help
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  build         Build the binary for all supported platforms"
	@echo "  install       Install dependencies"
	@echo "  run           Run the application"
	@echo "  clean         Clean the project"
	@echo "  test          Run tests"
	@echo "  test_coverage Run tests with coverage"
	@echo "  dep           Download dependencies"
	@echo "  vet           Run go vet"
	@echo "  lint          Run golangci-lint"

.PHONY: build
build: $(outputs)

.PHONY: install
install:
	go mod tidy

.PHONY: run
run:
	go run main.go

.PHONY: clean
clean:
	go clean
	rm -rf ${TARGET_DIR}

.PHONY: test
test:
	go test ./...

.PHONY: test_coverage
test_coverage:
	go test ./... -coverprofile=coverage.out

.PHONY: dep
dep:
	go mod download

.PHONY: vet
vet:
	go vet

.PHONY: lint
lint:
	golangci-lint run --enable-all
