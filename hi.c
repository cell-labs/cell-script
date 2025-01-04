// # Simple UDT
//
// A simple UDT script using 128 bit unsigned integer range
//
// This UDT has 2 unlocking modes:
//
// 1. If one of the transaction input has a lock script matching the UDT
// script argument, the UDT script will be in owner mode. In owner mode no
// checks is performed, the owner can perform any operations such as issuing
// more UDTs or burning UDTs. By ensuring at least one transaction input has
// a matching lock script, the ownership of UDT can be ensured.
// 2. Otherwise, the UDT script will be in normal mode, where it ensures the
// sum of all input tokens is not smaller than the sum of all output tokens.
//
// Notice one caveat of this UDT script is that only one UDT can be issued
// for each unique lock script. A more sophisticated UDT script might include
// other arguments(such as the hash of the first input) as a unique identifier,
// however for the sake of simplicity, we are happy with this limitation.

// First, let's include header files used to interact with CKB.
#if defined(CKB_SIMULATOR)
#include "ckb_syscall_simulator.h"
#else
#include "ckb_syscalls.h"
#endif
#include "blockchain.h"

// We are limiting the script size loaded to be 32KB at most. This should be
// more than enough. We are also using blake2b with 256-bit hash here, which is
// the same as CKB.
#define BLAKE2B_BLOCK_SIZE 32
#define SCRIPT_SIZE 32768

// Common error codes that might be returned by the script.
#define ERROR_ARGUMENTS_LEN -1
#define ERROR_ENCODING -2
#define ERROR_SYSCALL -3
#define ERROR_SCRIPT_TOO_LONG -21
#define ERROR_OVERFLOWING -51
#define ERROR_AMOUNT -52

// We will leverage gcc's 128-bit integer extension here for number crunching.
typedef unsigned __int128 uint128_t;

const int max_size = 32768;

typedef struct
{
  uint32_t len;
  uint32_t cap;
  uint32_t offset;
  uint8_t *ptr;
} SliceType;

uint8_t *__slice_get_ptr(SliceType *s)
{
  printf("gep = %d", s->ptr);
  return s->ptr;
}

void foo() {
  SliceType s;
  s.ptr = (uint8_t*)malloc(16);
  printf("ptr = %d", s.ptr);
  printf("ptr = %d", malloc(16));
  printf("ptr = %d", s.ptr);
  uint8_t* p = __slice_get_ptr(&s);
  printf("ptr = %d", (uint8_t*)s.ptr);
  printf("ptr = %d", p);
}

// char *strndup(const char *str, size_t n)
// {
//     printf("n = ");
//     if (!str)
//         return 0;

//     size_t length = strlen(str);
//     printf("%d", length);
//     if (length < n)
//         n = length;
//     printf("%d", n);
//     char *result = (char *)malloc(n + 1);
//     memcpy(result, str, n);
//     result[n] = '\0';
//     printf("%d", result);
//     return result;
// }
// void pd(const char*p, int len) {
//     for (int i = 0; i < len; i++)
//     {
//         printf("%d\n", p[i]);
//     }
// }
void px(const char*p, int len) {
    for (int i = 0; i < len; i++)
    {
        printf("%x\n", p[i]);
    }
}

void input_len() {
    int len = ckb_calculate_inputs_len();
    printf("input_len %d", len);
}
uint32_t unpackU32(char *b) {
    return b[0] | b[1]<<8 | b[2]<<16 | b[3]<<24;
}
uint64_t unpackU64(char *b) {
    return b[0] | b[1]<<8 | b[2]<<16 | b[3]<<24 | (uint64_t)b[4]<<32 | (uint64_t)b[5]<<40 | (uint64_t)b[6]<<48 | (uint64_t)b[7]<<56;
}
int main()
{
        // foo();
    input_len();
    char addr[max_size];
    memset(addr, '0', 186);
    printf("%s", addr);

    uint128_t a = 65535;
    uint128_t b = 1;
    uint128_t c = a + b;
    printf("a + b = %d", (int)a);
    printf("a + b = %d", (int)c);

    uint64_t len = max_size;
    printf("len = %d ", len);
    // printf("addr %d", addr);
    int ret = ckb_load_script(addr, &len, 0); //, 0, CKB_SOURCE_GROUP_INPUT);
    // px(addr+49, 32); // hash
    printf("len = %d ", len);
    printf("%s", addr);

    for (int i = 0; i < 7; i++) {
        int ret = ckb_load_cell(addr, &len, 0, i, CKB_SOURCE_INPUT);
        printf("input [%d] len = %d , err = %d, unpack_len = %d", i, len, ret, unpackU32(addr));
        for (int j = 0; j < CKB_CELL_FIELD_OCCUPIED_CAPACITY; j++) {
            len = -1;
            ret = ckb_load_cell_by_field(addr, &len, 0, i, CKB_SOURCE_INPUT, j);
            printf("filed len = %d", len);
            // px(addr, 32);
        }
        // printf("filed [%d] len = %d , err = %d, val = %x", CKB_CELL_FIELD_LOCK_HASH, len, ret, unpackU64(addr));
    }
    // for (int i = 0; i < 6; i++) {
    //     int ret = ckb_load_cell(addr, &len, 0, i, CKB_SOURCE_OUTPUT);
    //     printf("output [%d] len = %d , err = %d", i, len, ret);
    // }
    if (ret != CKB_SUCCESS)
    {
        return CKB_LENGTH_NOT_ENOUGH;
    }
    return CKB_SUCCESS;
}