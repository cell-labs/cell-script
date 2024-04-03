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

#ifdef CKB_SIMULATOR
int simulator_main() {
#else
int main() {
#endif
  // First, let's load current running script, so we can extract owner lock
  // script hash from script args.
  unsigned char script[SCRIPT_SIZE];
  uint64_t len = SCRIPT_SIZE;
  int ret = ckb_load_script(script, &len, 0);
  if (ret != CKB_SUCCESS) {
    return ERROR_SYSCALL;
  }
  if (len > SCRIPT_SIZE) {
    return ERROR_SCRIPT_TOO_LONG;
  }
  mol_seg_t script_seg;
  script_seg.ptr = (uint8_t *)script;
  script_seg.size = len;

  if (MolReader_Script_verify(&script_seg, false) != MOL_OK) {
    return ERROR_ENCODING;
  }

  mol_seg_t args_seg = MolReader_Script_get_args(&script_seg);
  mol_seg_t args_bytes_seg = MolReader_Bytes_raw_bytes(&args_seg);
  if (args_bytes_seg.size != BLAKE2B_BLOCK_SIZE) {
    return ERROR_ARGUMENTS_LEN;
  }

  // With owner lock script extracted, we will look through each input in the
  // current transaction to see if any unlocked cell uses owner lock.
  int owner_mode = 0;
  size_t i = 0;
  while (1) {
    uint8_t buffer[BLAKE2B_BLOCK_SIZE];
    uint64_t len = BLAKE2B_BLOCK_SIZE;
    // There are 2 points worth mentioning here:
    //
    // * First, we are using the checked version of CKB syscalls, the checked
    // versions will return an error if our provided buffer is not enough to
    // hold all returned data. This can help us ensure that we are processing
    // enough data here.
    // * Second, `CKB_CELL_FIELD_LOCK_HASH` is used here to directly load the
    // lock script hash, so we don't have to manually calculate the hash again
    // here.
    ret = ckb_checked_load_cell_by_field(buffer, &len, 0, i, CKB_SOURCE_INPUT,
                                         CKB_CELL_FIELD_LOCK_HASH);
    if (ret == CKB_INDEX_OUT_OF_BOUND) {
      break;
    }
    if (ret != CKB_SUCCESS) {
      return ret;
    }
    if (len != BLAKE2B_BLOCK_SIZE) {
      return ERROR_ENCODING;
    }
    if (memcmp(buffer, args_bytes_seg.ptr, BLAKE2B_BLOCK_SIZE) == 0) {
      owner_mode = 1;
      break;
    }
    i += 1;
  }

  // When owner mode is triggered, we won't perform any checks here, the owner
  // is free to make any changes here, including token issurance, minting, etc.
  if (owner_mode) {
    return CKB_SUCCESS;
  }

  // When the owner mode is not enabled, however, we will then need to ensure
  // the sum of all input tokens is not smaller than the sum of all output
  // tokens. First, let's loop through all input cells containing current UDTs,
  // and gather the sum of all input tokens.
  uint128_t input_amount = 0;
  i = 0;
  while (1) {
    uint128_t current_amount = 0;
    len = 16;
    // The implementation here does not require that the transaction only
    // contains UDT cells for the current UDT type. It's perfectly fine to mix
    // the cells for multiple different types of UDT together in one
    // transaction. But that also means we need a way to tell one UDT type from
    // another UDT type. The trick is in the `CKB_SOURCE_GROUP_INPUT` value used
    // here. When using it as the source part of the syscall, the syscall would
    // only iterate through cells with the same script as the current running
    // script. Since different UDT types will naturally have different
    // script(the args part will be different), we can be sure here that this
    // loop would only iterate through UDTs that are of the same type as the one
    // identified by the current running script.
    //
    // In the case that multiple UDT types are included in the same transaction,
    // this simple UDT script will be run multiple times to validate the
    // transaction, each time with a different script containing different
    // script args, representing different UDT types.
    //
    // A different trick used here, is that our current implementation assumes
    // that the amount of UDT is stored as unsigned 128-bit little endian
    // integer in the first 16 bytes of cell data. Since RISC-V also uses little
    // endian format, we can just read the first 16 bytes of cell data into
    // `current_amount`, which is just an unsigned 128-bit integer in C. The
    // memory layout of a C program will ensure that the value is set correctly.
    ret = ckb_load_cell_data((uint8_t *)&current_amount, &len, 0, i,
                             CKB_SOURCE_GROUP_INPUT);
    // When `CKB_INDEX_OUT_OF_BOUND` is reached, we know we have iterated
    // through all cells of current type.
    if (ret == CKB_INDEX_OUT_OF_BOUND) {
      break;
    }
    if (ret != CKB_SUCCESS) {
      return ret;
    }
    if (len < 16) {
      return ERROR_ENCODING;
    }
    input_amount += current_amount;
    // Like any serious smart contract out there, we will need to check for
    // overflows.
    if (input_amount < current_amount) {
      return ERROR_OVERFLOWING;
    }
    i += 1;
  }

  // With the sum of all input UDT tokens gathered, let's now iterate through
  // output cells to grab the sum of all output UDT tokens.
  uint128_t output_amount = 0;
  i = 0;
  while (1) {
    uint128_t current_amount = 0;
    len = 16;
    // Similar to the above code piece, we are also looping through output cells
    // with the same script as current running script here by using
    // `CKB_SOURCE_GROUP_OUTPUT`.
    ret = ckb_load_cell_data((uint8_t *)&current_amount, &len, 0, i,
                             CKB_SOURCE_GROUP_OUTPUT);
    if (ret == CKB_INDEX_OUT_OF_BOUND) {
      break;
    }
    if (ret != CKB_SUCCESS) {
      return ret;
    }
    if (len < 16) {
      return ERROR_ENCODING;
    }
    output_amount += current_amount;
    // Like any serious smart contract out there, we will need to check for
    // overflows.
    if (output_amount < current_amount) {
      return ERROR_OVERFLOWING;
    }
    i += 1;
  }

  // When both value are gathered, we can perform the final check here to
  // prevent non-authorized token issurance.
  if (input_amount < output_amount) {
    return ERROR_AMOUNT;
  }
  return CKB_SUCCESS;
}