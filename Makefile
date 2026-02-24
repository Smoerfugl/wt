# Simple Makefile for building the wt CLI

BINARY := wt
CMD := .
GO := go
LDFLAGS := -s -w

.PHONY: all build install run clean fmt vet test help

all: build

build:
	$(GO) build -ldflags="$(LDFLAGS)" -o $(BINARY) $(CMD)

install:
	$(GO) install $(CMD)

# Build then run; pass arguments with ARGS variable, e.g.:
#   make run ARGS="list"
run: build
	./$(BINARY) $(ARGS)

clean:
	-rm -f $(BINARY)

fmt:
	$(GO) fmt ./...

vet:
	$(GO) vet ./...

test:
	$(GO) test ./...

help:
	@echo "Targets:"
	@echo "  make        (same as 'make build')"
	@echo "  make build  Build binary ($(BINARY))"
	@echo "  make install Install binary to $$(GOBIN) or $$(GOPATH)/bin"
	@echo "  make run ARGS=\"list\" Build and run with ARGS"
	@echo "  make clean  Remove built binary"
	@echo "  make fmt    Run go fmt"
	@echo "  make vet    Run go vet"
	@echo "  make test   Run go test ./..."
