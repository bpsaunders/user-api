lint_output     := lint.txt
coverage_output := coverage.out
bin             := main

.EXPORT_ALL_VARIABLES:
GO111MODULE = on

.PHONY: all
all: build

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: clean
clean:
	go mod tidy
	rm -f $(lint_output) $(coverage_output) $(bin)

.PHONY: test
test:
	go test ./... -run 'Unit' -coverprofile=coverage.out

.PHONY: build
build: fmt clean
	go build -o ./$(bin)

.PHONY: lint
lint: GO111MODULE = off
lint:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
	gometalinter ./... > $(lint_output); true
