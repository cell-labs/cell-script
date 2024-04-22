#include <stdbool.h>

#include "bigint/BigInt.h"

typedef struct {
    int64_t size;
    char* ptr;
}string_t;

BigInt* big_int_new() {
    return BigInt_construct(0);
}

BigInt* big_int_clone(BigInt* big_int) {
    return BigInt_clone(big_int, big_int->num_digits);
}

void big_int_free(BigInt* big_int) {
    return BigInt_free(big_int);
}

BigInt* big_init_from_string(string_t str) {
    // Todo: check string size.
    return BigInt_from_string(str.ptr);
}

void big_int_print(const BigInt* big_int) {
    return BigInt_print(big_int);
}

uint32_t big_int_len(const BigInt* big_int) {
    return BigInt_strlen(big_int);
}

string_t big_init_to_string(const BigInt* big_int) {
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
    BigInt* res = big_int_clone(a);
    BOOL ok = BigInt_add(res, b);
    if (!ok) {
        exit(0);
        return *res;
    }
    return *res;
}

BigInt big_int_sub(const BigInt* a, const BigInt* b) {
    BigInt* res = big_int_clone(a);
    BOOL ok = BigInt_sub(res, b);
    if (!ok) {
        exit(0);
        return *res;
    }
    return *res;
}

BigInt big_int_mul(const BigInt* a, const BigInt* b) {
    BigInt* res = big_int_clone(a);
    BOOL ok = BigInt_multiply(res, b);
    if (!ok) {
        exit(0);
        return *res;
    }
    return *res;
}

BigInt big_int_div(const BigInt* a, const BigInt* b) {
    BigInt* quotient = big_int_new();
    BigInt* remainder = big_int_new();
    BOOL ok = BigInt_divide(a, b, quotient, remainder);
    if (!ok) {
        big_int_free(remainder);
        exit(0);
        return *quotient;
    }
    big_int_free(quotient);
    return *quotient;
}

BigInt big_int_mod(const BigInt* a, const BigInt* b) {
    BigInt* quotient = big_int_new();
    BigInt* remainder = big_int_new();
    BOOL ok = BigInt_divide(a, b, quotient, remainder);
    if (!ok) {
        big_int_free(quotient);
        exit(0);
        return *remainder;
    }
    big_int_free(quotient);
    return *remainder;
}

