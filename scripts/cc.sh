cmd="clang -o a.out --target=riscv64 \
    -march=rv64imc \
    -nostdlib \
    -Wall -Werror -Wextra -Wno-unused-parameter -Wno-nonnull -fno-builtin-printf -fno-builtin-memcmp -O3 -fdata-sections -ffunction-sections \
    -I third-party/ckb-c-stdlib/libc \
    -I third-party/ckb-c-stdlib/molecule \
    -I third-party/ckb-c-stdlib \
    -I third-party/sparse-merkle-tree/c \
    -DCKB_C_STDLIB_PRINTF=1 \
    -DCKB_PRINTF_DECLARATION_ONLY=1 \
    $1"
echo $cmd
$cmd
