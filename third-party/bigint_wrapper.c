#include <stdbool.h>

#define FILE void

#include "bigint/BigInt.h"
#include "ckb-c-stdlib/ckb_syscall_apis.h"

typedef struct {
    int64_t size;
    char* ptr;
}string_t;

BigInt* big_int_new() {
    BigInt* new_big_int = malloc(sizeof(BigInt));
    new_big_int->digits = (unsigned char*)"";
    new_big_int->is_negative = false;
    new_big_int->num_digits = 0;
    new_big_int->num_allocated_digits = 0;
    return new_big_int;
}

BigInt* big_int_clone(const BigInt* big_int) {
    return BigInt_clone(big_int, big_int->num_digits);
}

void big_int_free(BigInt* big_int) {
    return BigInt_free(big_int);
}

BigInt* big_int_from_string(string_t str) {
    // Todo: check string size.
    return BigInt_from_string(str.ptr);
}

void big_int_print(const BigInt* big_int) {
    const unsigned char* base = big_int->digits;
    const unsigned char* digits = &base[big_int->num_digits-1];
    if (big_int->is_negative) ckb_printf("-");
    while(digits >= base) {
        ckb_printf("%c", '0' + *(digits--));
    }
}

uint32_t big_int_len(const BigInt* big_int) {
    return BigInt_strlen(big_int);
}

string_t big_int_to_string(const BigInt* big_int) {
    string_t res;
    res.ptr = BigInt_to_new_string(big_int);
    res.size = big_int->num_digits;
    return res;
}

bool big_int_assign(BigInt* target, const BigInt* source) {
    return BigInt_assign(target, source);
}

bool big_int_gt(const BigInt* a, const BigInt* b) {
    return BigInt_compare(a, b) == 1;
}

bool big_int_gte(const BigInt* a, const BigInt* b) {
    int res = BigInt_compare(a, b);
    return res == 1 || res == 0;
}

bool big_int_lt(const BigInt* a, const BigInt* b) {
    return BigInt_compare(a, b) == -1;
}

bool big_int_lte(const BigInt* a, const BigInt* b) {
    int res = BigInt_compare(a, b);
    return res == -1 || res == 0;
}

bool big_int_equal(const BigInt* a, const BigInt* b) {
    return BigInt_compare(a, b) == 0;
}

BigInt big_int_add(const BigInt* a, const BigInt* b) {
    BigInt* res = big_int_clone((BigInt*)a);
    BOOL ok = BigInt_add(res, b);
    if (!ok) {
        ckb_exit(0);
        return *res;
    }
    return *res;
}

BigInt big_int_sub(const BigInt* a, const BigInt* b) {
    BigInt* res = big_int_clone(a);
    BOOL ok = BigInt_subtract(res, b);
    if (!ok) {
        ckb_exit(0);
        return *res;
    }
    return *res;
}

BigInt big_int_mul(const BigInt* a, const BigInt* b) {
    BigInt* res = big_int_clone(a);
    BOOL ok = BigInt_multiply(res, b);
    if (!ok) {
        ckb_exit(0);
        return *res;
    }
    return *res;
}

BigInt big_int_div(const BigInt* a, const BigInt* b) {
    BigInt* quotient = big_int_new();
    BigInt* remainder = big_int_new();
    BOOL ok = BigInt_divide((BigInt*)a, (BigInt*)b, quotient, remainder);
    if (!ok) {
        big_int_free(remainder);
        ckb_exit(0);
        return *quotient;
    }
    big_int_free(quotient);
    return *quotient;
}

BigInt big_int_mod(const BigInt* a, const BigInt* b) {
    BigInt* quotient = big_int_new();
    BigInt* remainder = big_int_new();
    BOOL ok = BigInt_divide((BigInt*)a, (BigInt*)b, quotient, remainder);
    if (!ok) {
        big_int_free(quotient);
        ckb_exit(0);
        return *remainder;
    }
    big_int_free(quotient);
    return *remainder;
}

