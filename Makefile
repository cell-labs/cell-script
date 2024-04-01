SHELL:=/bin/sh

MKFILE_PATH := ${abspath $(lastword $(MAKEFILE_LIST))}
MKFILE_DIR := $(dir $(MKFILE_PATH))
RELEASE_DIR := ${MKFILE_DIR}/output

.phony: clean antlr grammar dev build test test_cell_examples
clean:
	rm -rf internal/parser
	rm -rf internal/lexer
	rm -rf output
grammar: antlr

antlr:
	go generate ./...
dev:
	go run main.go
fmt:
	cd ${MKFILE_DIR} && go fmt ./...
build:
	@echo "build"
	go build -v -trimpath \
		-o ${RELEASE_DIR}/cell ./cmd/cell
	rm -f cell
	ln -s ${RELEASE_DIR}/cell cell
test:
	@echo "unit test"
	go mod tidy
	git diff --exit-code go.mod go.sum
	go mod verify
	go test -v -gcflag "all=-l" ${MKFILE_DIR}
example:
	@echo "test cell examples"
	make build
	./cell || true
	./cell tests/examples/hi.cell && ./hi
	./cell -t riscv tests/examples/always-true.cell && ckb-debugger --bin always-true