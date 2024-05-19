.PHONY: xudt-c

# ckb-c-stdlib depends on ckb-c
xudt-c:
	clang --target=riscv64 -v \
		-march=rv64imc \
		-nostdlib \
		-Wall -Werror -Wextra -Wno-unused-parameter -Wno-nonnull -fno-builtin-printf -fno-builtin-memcmp -O3 -fdata-sections -ffunction-sections \
		-I ../ckb-c-stdlib/libc \
		-I ../ckb-c-stdlib/molecule \
		-I ../ckb-c-stdlib \
		-I ../sparse-merkle-tree/c \
		-c xudt.c
	riscv64-unknown-elf-ar rcs libdummylibc.a xudt.o
	@echo " >>> sussecfully build xudt-c"
	cp libdummylibc.a ../../output/pkg
	@echo " >>> sussecfully install xudt-c"
xudt-c/debug:
	clang --target=riscv64 -v \
		-march=rv64imc \
		-nostdlib \
		-Wall -Werror -Wextra -Wno-unused-parameter -Wno-nonnull -fno-builtin-printf -fno-builtin-memcmp -O3 -g -fdata-sections -ffunction-sections \
		-I ../ckb-c-stdlib/libc \
		-I ../ckb-c-stdlib/molecule \
		-I ../ckb-c-stdlib \
		-I ../sparse-merkle-tree/c \
		-c xudt.c \
		-DCKB_C_STDLIB_PRINTF=1 \
		-DCKB_PRINTF_DECLARATION_ONLY=1 && \
	riscv64-unknown-elf-ar rcs libdummylibc-debug.a xudt.o
	@echo " >>> sussecfully build xudt-c"
	cp libdummylibc.a ../../output/pkg
	@echo " >>> sussecfully install xudt-c"
