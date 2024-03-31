SHELL:=/bin/sh

MKFILE_PATH := ${abspath $(lastword $(MAKEFILE_LIST))}
MKFILE_DIR := $(dir $(MKFILE_PATH))
RELEASE_DIR := ${MKFILE_DIR}/output

.phony: clean antlr grammar dev build test test_cell_examples
clean:
	rm -rf internal/parser
	rm -rf internal/lexer
grammar: antlr

antlr:
	go generate ./...
dev:
	go run main.go
fmt:
	cd ${MKFILE_DIR} && go fmt ./...
build:
	@echo "build"
	cd ${MKFILE_DIR} && \
	go build -v -trimpath \
	-o ${RELEASE_DIR}/go-to ${MKFILE_DIR}
test:
	@echo "unit test"
	go test ./...
	@echo "test cell examples"
	./cell || \
	./cell tests/examples/hi.cell