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
	make ckb-libc
	go build -v -trimpath \
		-o ${CELL} ./cmd/cell
	cp -r pkg/* output/pkg
	@echo "sussecfully build cell"
build/debug:
	go build -gcflags=all="-N -l" ./cmd/cell
ckb-libc: ckb-libc-release
ckb-libc-debug:
	@echo " >>> build libdummy-debug.a"
	cd third-party/ckb-c-stdlib && \
	clang --target=riscv64 -v \
		-march=rv64imc \
		-Wall -Werror -Wextra -Wno-unused-parameter -Wno-nonnull -fno-builtin-printf -fno-builtin-memcmp -O3 -g -fdata-sections -ffunction-sections \
		-I libc \
		-I . \
		-c libc/src/impl.c \
		-DCKB_C_STDLIB_PRINTF=1 \
		-o impl.o && \
	riscv64-unknown-elf-ar rcs libdummylibc-debug.a impl.o
	mkdir -p output/pkg
	cp -r third-party/ckb-c-stdlib/libdummylibc-debug.a output/pkg
	@echo "sussecfully build libdummy-debug.a"
ckb-libc-release:
	@echo " >>> build libdummy.a"
	cd third-party/ckb-c-stdlib && \
	clang --target=riscv64 \
		-march=rv64imc \
		-Wall -Werror -Wextra -Wno-unused-parameter -Wno-nonnull -fno-builtin-printf -fno-builtin-memcmp -O3 -g -fdata-sections -ffunction-sections \
		-I libc \
		-I . \
		-c ../wrapper.c && \
	riscv64-unknown-elf-ar rcs libdummylibc.a wrapper.o
	mkdir -p output/pkg
	cp -r third-party/ckb-c-stdlib/libdummylibc.a output/pkg
	@echo "sussecfully build libdummy.a"
install:
	@echo "manually run following command"
	@echo "source ./install.sh"
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
	${CELL} -t riscv tests/examples/always-true.cell && ckb-debugger --bin always-true
	${CELL} -t riscv tests/examples/helloworld.cell && ckb-debugger --bin helloworld
	${CELL} -t riscv tests/examples/table.cell && ckb-debugger --bin table
	${CELL} -t riscv tests/examples/cell-data.cell && ckb-debugger --bin cell-data
	${CELL} -t riscv tests/examples/inputs.cell && ckb-debugger --bin inputs
	${CELL} -t riscv tests/examples/outputs.cell && ckb-debugger --bin outputs

	${CELL} -t riscv tests/examples/multi-files && ckb-debugger --bin multi-files
	${CELL} -t riscv tests/examples/import-package && ckb-debugger --bin import-package
test/cross:
	@echo "test cross compiling"
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

