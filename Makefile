SHELL:=/bin/sh

MKFILE_PATH := ${abspath $(lastword $(MAKEFILE_LIST))}
MKFILE_DIR := $(dir $(MKFILE_PATH))
RELEASE_DIR := ${MKFILE_DIR}/output

CELL := ${RELEASE_DIR}/cell

.phony: clean antlr grammar dev build test test_cell_examples
clean:
	# rm -rf internal/parser
	# rm -rf internal/lexer
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
	git submodule update --init --recursive
	make ckblibc
	go build -v -trimpath \
		-o ${CELL} ./cmd/cell
	rm -f cell
	ln -s ${CELL} cell
ckblibc:
	@echo " >>> build libdummy.a"
	cd third-party/ckb-c-stdlib && \
	clang --target=riscv64 \
		-march=rv64imc \
		-Wall -Werror -Wextra -Wno-unused-parameter -Wno-nonnull -fno-builtin-printf -fno-builtin-memcmp -O3 -g -fdata-sections -ffunction-sections \
		-I libc \
		-c libc/src/impl.c \
		-o impl.o && \
	riscv64-unknown-elf-ar rcs libdummylibc.a impl.o
	mkdir -p output/pkg
	cp -r third-party/ckb-c-stdlib/libdummylibc.a output/pkg
test:
	@echo "unit test"
	go mod tidy
	git diff --exit-code go.mod go.sum
	go mod verify
	go test -v -gcflag "all=-l" ${MKFILE_DIR}
test/example:
	@echo "test cell examples"
	make build
	${CELL} || true
	${CELL} tests/examples/hi.cell && ./hi
	${CELL} -d -t riscv tests/examples/always-true.cell && ckb-debugger --bin always-true
test/cross:
	@echo "test cross compiling"
	@echo cross hi.ll with linking dummy.c
	which clang
	clang --target=riscv64 \
		-march=rv64imc \
		-Wno-override-module \
		-ffunction-sections -fdata-sections \
		-nostdlib \
		-L output/pkg \
		-ldummylibc \
		-Wl,--gc-sections \
		-o main tests/examples/hi.ll
	ckb-debugger --bin main
