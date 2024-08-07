enum ErrorCode
{
  // 0 is the only success code. We can use 0 directly.

  // inherit from simple_udt
  ERROR_ARGUMENTS_LEN = -1,
  ERROR_ENCODING = -2,
  ERROR_SYSCALL = -3,
  ERROR_SCRIPT_TOO_LONG = -21,
  ERROR_OVERFLOWING = -51,
  ERROR_AMOUNT = -52,

  // error code is starting from 40, to avoid conflict with
  // common error code in other scripts.
  ERROR_CANT_LOAD_LIB = 40,
  ERROR_CANT_FIND_SYMBOL,
  ERROR_INVALID_RCE_ARGS,
  ERROR_NOT_ENOUGH_BUFF,
  ERROR_INVALID_FLAG,
  ERROR_INVALID_ARGS_FORMAT,
  ERROR_INVALID_WITNESS_FORMAT,
  ERROR_INVALID_MOL_FORMAT,
  ERROR_BLAKE2B_ERROR,
  ERROR_HASH_MISMATCHED,
  ERROR_RCRULES_TOO_DEEP, // 50
  ERROR_TOO_MANY_RCRULES,
  ERROR_RCRULES_PROOFS_MISMATCHED,
  ERROR_SMT_VERIFY_FAILED,
  ERROR_RCE_EMERGENCY_HALT,
  ERROR_NOT_VALIDATED,
  ERROR_TOO_MANY_LOCK,
  ERROR_ON_BLACK_LIST,
  ERROR_ON_BLACK_LIST2,
  ERROR_NOT_ON_WHITE_LIST,
  ERROR_TYPE_FREEZED, // 60
  ERROR_APPEND_ONLY,
  ERROR_EOF,
  ERROR_TOO_LONG_PROOF,
};
