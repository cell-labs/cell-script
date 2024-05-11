SHELL:=/bin/sh

MKFILE_PATH := ${abspath $(lastword $(MAKEFILE_LIST))}
MKFILE_DIR := $(dir $(MKFILE_PATH))
RELEASE_DIR := ${MKFILE_DIR}/output

# XUDT
MOLECULEC := moleculec
MOLECULEC2 := ${MKFILE_DIR}/third-party/molecule2-c2/target/release/moleculec-c2

CELL := ${RELEASE_DIR}/cell
.PHONY: clean antlr grammar dev build test test_cell_examples
clean:
	# rm -rf internal/parser
	# rm -rf internal/lexer
	rm -rf output
	rm -f third-party/ckb-c-stdlib/*.o
	rm -f third-party/ckb-c-stdlib/*.a
grammar: antlr

all: build install
debug: build/debug install
antlr:
	go generate ./...
dev:
	go run main.go
fmt:
	cd ${MKFILE_DIR} && go fmt ./...
build:
	@echo " >>> build"
	make clean
	git submodule update --init --recursive
	make ckb-libc
	go build -v -trimpath \
		-o ${CELL} ./cmd/cell
	@echo " >>> sussecfully build cell"
build/debug:
	go build -gcflags=all="-N -l" -o ${CELL} ./cmd/cell
sudt-c:
	@echo " >>> build sudt-c"
	cd third-party/ckb-c-stdlib && \
	clang --target=riscv64 \
		-march=rv64imc \
		-nostdlib \
		-Wall -Werror -Wextra -Wno-unused-parameter -Wno-nonnull -fno-builtin-printf -fno-builtin-memcmp -O3 -fdata-sections -ffunction-sections \
		-I libc \
		-I molecule \
		-I . \
		../sudt.c \
		-o sudt-c && \
	cp sudt-c ../..
	@echo " >>> sussecfully build sudt-c"
molecule-xudt:
	cd third-party/molecule2-c2 && cargo build --release
	@echo "generate mol header files for xudt"
	cd third-party/xudt && ${MOLECULEC} --language c --schema-file xudt_rce.mol > xudt_rce_mol.h
	@echo "generate mol2 header files for xudt"
	cd third-party/xudt && ${MOLECULEC} --language - --schema-file xudt_rce.mol --format json > blockchain_mol2.json
	cd third-party/xudt && ${MOLECULEC2} --input blockchain_mol2.json | clang-format -style=Google > xudt_rce_mol2.h
xudt-c: molecule-xudt
	cd third-party && \
	clang --target=riscv64 \
		-march=rv64imc \
		-nostdlib \
		-Wall -Werror -Wextra -Wno-unused-parameter -Wno-nonnull -fno-builtin-printf -fno-builtin-memcmp -O3 -fdata-sections -ffunction-sections \
		-I ckb-c-stdlib/libc \
		-I ckb-c-stdlib/molecule \
		-I ckb-c-stdlib \
		-I sparse-merkle-tree/c \
		xudt/*.c \
		-o xudt-c && \
	cp xudt-c ..
	@echo " >>> sussecfully build xudt-c"
ckb-libc: ckb-libc-debug ckb-libc-release install
ckb-libc-debug:
	@echo " >>> build libdummylibc-debug.a"
	cd third-party/ckb-c-stdlib && \
	clang --target=riscv64 -v \
		-march=rv64imc \
		-nostdlib \
		-Wall -Werror -Wextra -Wno-unused-parameter -Wno-nonnull -fno-builtin-printf -fno-builtin-memcmp -O3 -g -fdata-sections -ffunction-sections \
		-I libc \
		-I molecule \
		-I . \
		-I ../sparse-merkle-tree/c \
		-c ../xudt/xudt.c \
		-DCKB_C_STDLIB_PRINTF=1 \
		-DCKB_PRINTF_DECLARATION_ONLY=1 && \
	riscv64-unknown-elf-ar rcs libdummylibc-debug.a xudt.o
ckb-libc-release:
	@echo " >>> build libdummylibc.a"
	cd third-party/ckb-c-stdlib && \
	clang --target=riscv64 \
		-march=rv64imc \
		-nostdlib \
		-Wall -Werror -Wextra -Wno-unused-parameter -Wno-nonnull -fno-builtin-printf -fno-builtin-memcmp -O3 -fdata-sections -ffunction-sections \
		-I libc \
		-I molecule \
		-I . \
		-I ../sparse-merkle-tree/c \
		-c ../xudt/xudt.c && \
	riscv64-unknown-elf-ar rcs libdummylibc.a xudt.o
install:
	mkdir -p output/pkg
	cp -r third-party/ckb-c-stdlib/libdummylibc-debug.a output/pkg
	@echo " >>> sussecfully install libdummylibc-debug.a"
	cp -r third-party/ckb-c-stdlib/libdummylibc.a output/pkg
	@echo " >>> sussecfully install libdummylibc.a"
	cp -r pkg/* output/pkg
	@echo " >>> sussecfully install stdlib"

	@echo " >>> manually run following command"
	@echo "source ./install.sh"
test: unittest test/example
unittest:
	@echo "unit test"
	go mod tidy
	git diff --exit-code go.mod go.sum
	go mod verify
	go test -v ${MKFILE_DIR}/compiler/lexer/*.go
	go test -v ${MKFILE_DIR}/compiler/parser/*.go
	go test -v ${MKFILE_DIR}/compiler/passes/bigint/*.go
test/example:
	@echo " >>> test cell examples"
	make build
	${CELL} || true
	${CELL} tests/examples/hi.cell && ./hi
	${CELL} -d -t riscv tests/examples/always-true.cell && ckb-debugger --bin always-true
	${CELL} -d -t riscv tests/examples/helloworld.cell && ckb-debugger --bin helloworld | grep "hello world! 0"
	${CELL} -t riscv tests/examples/table.cell && ckb-debugger --bin table
	${CELL} -t riscv tests/examples/cell-data.cell && ckb-debugger --bin cell-data
	${CELL} -t riscv tests/examples/inputs.cell && ckb-debugger --bin inputs
	${CELL} -t riscv tests/examples/outputs.cell && ckb-debugger --bin outputs
	${CELL} -t riscv tests/examples/sudt.cell && ckb-debugger --bin sudt || true

	${CELL} -t riscv tests/examples/multi-files && ckb-debugger --bin multi-files
	${CELL} -t riscv tests/examples/import-package && ckb-debugger --bin import-package
	
	${CELL} -t riscv tests/examples/brainfuck-vm.cell && ckb-debugger --bin brainfuck-vm
	${CELL} -t riscv tests/examples/byte.cell && ckb-debugger --bin byte
	${CELL} -t riscv tests/examples/xudt-data.cell && ckb-debugger --bin xudt-data

test/cross:
	@echo " >>> test cross compiling"
	@echo cross hi.ll with linking dummy.c
	which clang
	clang --target=riscv64 \
		-march=rv64imc \
		-Wno-override-module \
		-ffunction-sections -fdata-sections \
		-DCKB_C_STDLIB_PRINTF=1 \
		-nostdlib \
		-L output/pkg \
		-ldummylibc-debug \
		-Wl,--gc-sections \
		-o main tests/examples/hi.ll
	ckb-debugger --bin main

