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
		-o ${RELEASE_DIR}/cell ${MKFILE_DIR}
	ln -s ${RELEASE_DIR}/cell cell
test:
	@echo "unit test"
	go mod tidy
	git diff --exit-code go.mod go.sum
	go mod verify
	go test -v -gcflag "all=-l" ${MKFILE_DIR}
	@echo "test cell examples"
	./cell || \
	./cell tests/examples/hi.cell