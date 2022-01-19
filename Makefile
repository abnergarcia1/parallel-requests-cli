
default: build

.PHONY: build
build:
	go build

.PHONY: install
install:
	go install

.PHONY: lint
lint:
	@echo "==> Checking source code against linters..."
	golangci-lint run

.PHONY: test
test:
	go test $(go list ./...)