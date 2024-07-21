#include "ckb_syscalls.h"
#include "blake2b.h"
#ifndef MOL2_EXIT
#define MOL2_EXIT ckb_exit
#endif
int ckb_exit(signed char);
#include "molecule/blockchain.h"
#include "molecule/blockchain-api2.h"
#include "error.h"
#include "slice.h"
// We are limiting the script size loaded to be 32KB at most. This should be
// more than enough. We are also using blake2b with 256-bit hash here, which is
// the same as CKB.
const uint32_t BLAKE2B_BLOCK_SIZE = 32;
const uint32_t SCRIPT_SIZE = 32768;
const uint32_t MAX_CELLS = 16; // todo
const uint32_t MAX_DATA_SIZE = 4 * 1024 * 1024;

// Common error codes that might be returned by the script.
#define ERROR_ARGUMENTS_LEN -1
#define ERROR_ENCODING -2
#define ERROR_SYSCALL -3
#define ERROR_SCRIPT_TOO_LONG -21

typedef struct
{
  uint64_t data;
} cell_meta_data;

typedef struct
{
  uint32_t size;
  uint32_t cap;
  uint32_t offset;
  cell_meta_data *ptr;
} cell_data_t;

cell_data_t EMPTY_CELL_DATA = {0};

bool script_verify()
{
  mol_seg_t script_seg;
  // mol_seg_t args_seg;
  // mol_seg_t args_bytes_seg;
  // First, let's load current running script, so we can extract owner lock
  // script hash from script args.
  unsigned char script[SCRIPT_SIZE];
  uint64_t len = SCRIPT_SIZE;
  int ret = ckb_load_script(script, &len, 0);
  if (ret != CKB_SUCCESS)
  {
    return false;
  }
  if (len > SCRIPT_SIZE)
  {
    return false;
  }

  script_seg.ptr = (uint8_t *)script;
  script_seg.size = len;
  if (MolReader_Script_verify(&script_seg, false) != MOL_OK)
  {
    return false;
  }
  // should be checked in sudt
  // args_seg = MolReader_Script_get_args(&script_seg);
  // args_bytes_seg = MolReader_Bytes_raw_bytes(&args_seg);
  // if (args_bytes_seg.size != BLAKE2B_BLOCK_SIZE)
  // {
  //   return false;
  // }
  return true;
}

bool is_owner_mode()
{
  unsigned char script[SCRIPT_SIZE];
  uint64_t len = SCRIPT_SIZE;
  int ret = ckb_load_script(script, &len, 0);
  if (ret != CKB_SUCCESS)
  {
    return false;
  }
  if (len > SCRIPT_SIZE)
  {
    return false;
  }

  mol_seg_t script_seg;
  script_seg.ptr = (uint8_t *)script;
  script_seg.size = len;
  if (MolReader_Script_verify(&script_seg, false) != MOL_OK)
  {
    return false;
  }

  mol_seg_t args_seg = MolReader_Script_get_args(&script_seg);
  mol_seg_t args_bytes_seg = MolReader_Bytes_raw_bytes(&args_seg);
  if (args_bytes_seg.size != BLAKE2B_BLOCK_SIZE)
  {
    return false;
  }

  // With owner lock script extracted, we will look through each input in the
  // current transaction to see if any unlocked cell uses owner lock.
  int owner_mode = 0;
  size_t i = 0;
  while (1)
  {
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
    if (ret == CKB_INDEX_OUT_OF_BOUND)
    {
      break;
    }
    if (ret != CKB_SUCCESS)
    {
      return false;
    }
    if (len != BLAKE2B_BLOCK_SIZE)
    {
      return false;
    }
    if (memcmp(buffer, args_bytes_seg.ptr, BLAKE2B_BLOCK_SIZE) == 0)
    {
      owner_mode = 1;
      break;
    }
    i += 1;
  }
  return (owner_mode == 1);
}

uint64_t get_input_cell_data_len(int i)
{
  uint64_t len = 0;
  int ret =
      ckb_load_cell_data(NULL, &len, 0, i, CKB_SOURCE_GROUP_INPUT);
  if (ret == CKB_INDEX_OUT_OF_BOUND)
  {
    return 0;
  }
  if (ret != CKB_SUCCESS)
  {
    return 0;
  }
  if (len < 16)
  {
    return 0;
  }
  return len;
}

uint64_t get_output_cell_data_len(int i)
{
  uint64_t len = 0;
  int ret =
      ckb_load_cell_data(NULL, &len, 0, i, CKB_SOURCE_GROUP_OUTPUT);
  if (ret == CKB_INDEX_OUT_OF_BOUND)
  {
    return 0;
  }
  if (ret != CKB_SUCCESS)
  {
    return 0;
  }
  if (len < 16)
  {
    return 0;
  }
  return len;
}

cell_data_t get_utxo_inputs()
{
  cell_data_t inputs = {0};
  inputs.ptr = malloc(MAX_CELLS * sizeof(mol_seg_t));
  int i = 0;
  while (1)
  {
    uint64_t len = get_input_cell_data_len(i);
    if (len < 16)
    {
      return EMPTY_CELL_DATA;
    }
    // uint8_t* data = (uint8_t*)malloc(len * sizeof(uint8_t));
    uint64_t cur_data = 0;
    int ret = ckb_load_cell_data(&cur_data, &len, 0, i,
                                 CKB_SOURCE_GROUP_INPUT);
    // When `CKB_INDEX_OUT_OF_BOUND` is reached, we know we have iterated
    // through all cells of current type.
    if (ret == CKB_INDEX_OUT_OF_BOUND)
    {
      break;
    }
    if (ret != CKB_SUCCESS)
    {
      return EMPTY_CELL_DATA;
    }
    if (len < 16)
    {
      return EMPTY_CELL_DATA;
    }
    if (i >= (int)MAX_CELLS)
    {
      free(inputs.ptr);
      return EMPTY_CELL_DATA;
    }
    inputs.ptr[i].data = cur_data;
    i += 1;
    inputs.size = i;
  };
  return inputs;
}

cell_data_t get_utxo_outputs()
{
  cell_data_t outputs = {0};
  outputs.ptr = malloc(MAX_CELLS * sizeof(mol_seg_t));
  int i = 0;
  while (1)
  {
    uint64_t len = get_output_cell_data_len(i);
    uint64_t cur_data = 0;
    int ret = ckb_load_cell_data(&cur_data, &len, 0, i,
                                 CKB_SOURCE_GROUP_OUTPUT);
    // When `CKB_INDEX_OUT_OF_BOUND` is reached, we know we have iterated
    // through all cells of current type.
    if (ret == CKB_INDEX_OUT_OF_BOUND)
    {
      break;
    }
    if (ret != CKB_SUCCESS)
    {
      return EMPTY_CELL_DATA;
    }
    if (len < 16)
    {
      return EMPTY_CELL_DATA;
    }
    if (i >= (int)MAX_CELLS)
    {
      free(outputs.ptr);
      return EMPTY_CELL_DATA;
    }
    outputs.ptr[i].data = cur_data;
    i += 1;
    outputs.size = i;
  }
  return outputs;
}

void syscall_exit(int8_t code)
{
  ckb_exit(code);
  return;
}

// xUDT
const uint32_t XUDT_FLAG_SIZE = 4;
const uint32_t MAX_LOCK_SCRIPT_HASH_COUNT = 2048; // maximum inputs
typedef struct
{
  int64_t size;
  uint8_t *ptr;
} String;
String mol_seg_to_string(mol_seg_t *seg)
{
  String s;
  s.size = seg->size;
  s.ptr = seg->ptr;
  return s;
}

mol_seg_t parse_xudt_args_as_mol_seg()
{
  mol_seg_t args;
  args.ptr = NULL;
  args.size = 0;
  unsigned char script[SCRIPT_SIZE];
  uint64_t len = SCRIPT_SIZE;
  int ret = ckb_load_script(script, &len, 0);
  if (ret != CKB_SUCCESS)
  {
    return args;
  }
  if (len > SCRIPT_SIZE)
  {
    return args;
  }

  mol_seg_t script_seg;
  script_seg.ptr = (uint8_t *)script;
  script_seg.size = len;
  mol_seg_t args_seg = MolReader_Script_get_args(&script_seg);
  mol_seg_t args_bytes_seg = MolReader_Bytes_raw_bytes(&args_seg);
  if (args_bytes_seg.size < BLAKE2B_BLOCK_SIZE)
  {
    return args;
  }

  if (args_bytes_seg.size >= (BLAKE2B_BLOCK_SIZE + XUDT_FLAG_SIZE))
  {
    // uint32_t val = *(uint32_t *)(args_bytes_seg.ptr + BLAKE2B_BLOCK_SIZE);
    return args_bytes_seg;
  }
  return args;
}

// get xudt args raw bytes
String parse_args()
{
  mol_seg_t args = parse_xudt_args_as_mol_seg();
  return mol_seg_to_string(&args);
}

typedef struct
{
  uint32_t len;
  uint32_t cap;
  uint32_t offset;
  uint8_t *bytes;
} ByteSlice;
ByteSlice make_byte_slice(uint32_t len, uint32_t cap, uint32_t offset, uint8_t *ptr)
{
  ByteSlice bs;
  bs.len = len;
  bs.cap = cap;
  bs.bytes = ptr;
  return bs;
}
// load hashes from all inputs cells
ByteSlice lock_scripts()
{
  uint32_t hashes_count = 0;
  uint8_t *hashes = malloc(sizeof(uint8_t) * BLAKE2B_BLOCK_SIZE * MAX_LOCK_SCRIPT_HASH_COUNT);
  memset(hashes, 0, BLAKE2B_BLOCK_SIZE * MAX_LOCK_SCRIPT_HASH_COUNT);
  ByteSlice bs = make_byte_slice(0, 0, 0, NULL);
  // collect hashes
  size_t i = 0;
  while (1)
  {
    uint8_t buffer[BLAKE2B_BLOCK_SIZE];
    uint64_t len = BLAKE2B_BLOCK_SIZE;
    int err = ckb_checked_load_cell_by_field(buffer, &len, 0, i, CKB_SOURCE_INPUT,
                                             CKB_CELL_FIELD_LOCK_HASH);
    if (err == CKB_INDEX_OUT_OF_BOUND)
    {
      break;
    }
    if (err != 0)
    {
      return bs;
    }
    if (hashes_count >= MAX_LOCK_SCRIPT_HASH_COUNT)
    {
      ckb_exit(ERROR_TOO_MANY_LOCK);
      return bs;
    }

    memcpy(hashes + hashes_count * BLAKE2B_BLOCK_SIZE, buffer,
           BLAKE2B_BLOCK_SIZE);
    hashes_count += 1;
    i += 1;
  }
  // make a string slice
  return make_byte_slice(hashes_count * BLAKE2B_BLOCK_SIZE, hashes_count * BLAKE2B_BLOCK_SIZE, 0, hashes);
}

typedef struct
{
  int64_t err;
  bool val;
} OptionBool;
OptionBool check_owner_mode(int64_t source, int64_t field)
{
  int err = CKB_SUCCESS;
  bool owner_mode = false;
  size_t i = 0;
  uint8_t buffer[BLAKE2B_BLOCK_SIZE];

  while (1)
  {
    uint64_t len = BLAKE2B_BLOCK_SIZE;
    err = ckb_checked_load_cell_by_field(buffer, &len, 0, i, (size_t)source, (size_t)field);
    if (err == CKB_INDEX_OUT_OF_BOUND)
    {
      err = CKB_SUCCESS;
      break;
    }
    if (err == CKB_ITEM_MISSING)
    {
      i += 1;
      err = CKB_SUCCESS;
      continue;
    }
    if (err != CKB_SUCCESS)
    {
      OptionBool ret = {err, owner_mode};
      return ret;
    }
    // skip memcpy here, which is not actually used later
    owner_mode = true;
    OptionBool ret = {err, owner_mode};
    return ret;
  }
  OptionBool ret = {err, owner_mode};
  return ret;
}

typedef uint8_t byte;
typedef struct
{
  String code_hash;
  byte hash_type;
  String args;
} Script;
typedef struct
{
  uint32_t len;
  uint32_t cap;
  uint32_t offset;
  Script *data;
} ScriptSlice;
int64_t checked_load_script_vec_size(uint8_t *ptr, uint32_t size, uint64_t *real_size)
{
  if (size < MOL_NUM_T_SIZE)
  {
    return ERROR_INVALID_MOL_FORMAT;
  }
  mol_num_t full_size = mol_unpack_number(ptr);
  *real_size = full_size;
  if (*real_size > size)
  {
    *real_size = 0;
    return ERROR_INVALID_MOL_FORMAT;
  }
  return CKB_SUCCESS;
}

uint64_t get_vec_size(String args)
{
  uint64_t real_size = 0;
  checked_load_script_vec_size(args.ptr, args.size, &real_size);
  return real_size;
}

const uint32_t BLAKE160_SIZE = 160 / 8;
const uint32_t RAW_EXTENSION_SIZE = 65536;
String get_witness()
{
  uint64_t witness_len = 0;
  // at the beginning of the transactions including RCE,
  // there is no "witness" in CKB_SOURCE_GROUP_INPUT
  // here we use the first witness of CKB_SOURCE_GROUP_OUTPUT
  // same logic is applied to rce_validator
  size_t source = CKB_SOURCE_GROUP_INPUT;
  int err = ckb_load_witness(NULL, &witness_len, 0, 0, source);
  if (err == CKB_INDEX_OUT_OF_BOUND)
  {
    source = CKB_SOURCE_GROUP_OUTPUT;
    err = ckb_load_witness(NULL, &witness_len, 0, 0, source);
  }
  String str;
  str.ptr = NULL;
  str.size = 0;
  if (err != 0)
  {
    return str;
  }
  if (witness_len <= 0)
  {
    err = ERROR_INVALID_MOL_FORMAT;
    return str;
  }

  uint8_t *witness_bytes = malloc(RAW_EXTENSION_SIZE);
  ckb_load_witness(witness_bytes, &witness_len, 0, 0, source);
  str.ptr = witness_bytes;
  str.size = witness_len;
  return str;
}
String get_blake2b(String data)
{
  uint8_t *hash = malloc(BLAKE2B_BLOCK_SIZE);
  String str;
  str.ptr = hash;
  str.size = BLAKE2B_BLOCK_SIZE;
  int err = blake2b(hash, BLAKE2B_BLOCK_SIZE, data.ptr, data.size, NULL, 0);
  if (err != 0)
  {
    str.size = 0;
    return str;
  }
  return str;
}

// uint32_t read_from_witness(uintptr_t arg[], uint8_t *ptr, uint32_t len,
//                                   uint32_t offset) {
//   int err;
//   uint64_t output_len = len;
//   err = ckb_load_witness(ptr, &output_len, offset, arg[0], arg[1]);
//   if (err != 0) {
//     return 0;
//   }
//   if (output_len > len) {
//     return len;
//   } else {
//     return output_len;
//   }
// }
// int make_cursor_from_witness(WitnessArgsType *witness, bool *use_input_type) {
//   uint64_t witness_len = 0;
//   // at the beginning of the transactions including RCE,
//   // there is no "witness" in CKB_SOURCE_GROUP_INPUT
//   // here we use the first witness of CKB_SOURCE_GROUP_OUTPUT
//   // same logic is applied to rce_validator
//   size_t source = CKB_SOURCE_GROUP_INPUT;
//   int err = ckb_load_witness(NULL, &witness_len, 0, 0, source);
//   if (err == CKB_INDEX_OUT_OF_BOUND) {
//     source = CKB_SOURCE_GROUP_OUTPUT;
//     err = ckb_load_witness(NULL, &witness_len, 0, 0, source);
//     *use_input_type = false;
//   } else {
//     *use_input_type = true;
//   }
//   if (err != 0) {
//     return 0;
//   }
//   if (witness_len <= 0) {
//     return ERROR_INVALID_MOL_FORMAT;
//   }

//   mol2_cursor_t cur;
//   cur.offset = 0;
//   cur.size = witness_len;

//   mol2_data_source_t *ptr = (mol2_data_source_t *)malloc(DEFAULT_DATA_SOURCE_LENGTH);

//   ptr->read = read_from_witness;
//   ptr->total_size = witness_len;
//   // pass index and source as args
//   ptr->args[0] = 0;
//   ptr->args[1] = source;

//   ptr->cache_size = 0;
//   ptr->start_point = 0;
//   ptr->max_cache_size = MAX_CACHE_SIZE;
//   cur.data_source = ptr;

//   *witness = make_WitnessArgs(&cur);

//   err = 0;

//   return err;
// }
// typedef struct ScriptVecOptVTable {
//   bool (*is_none)(struct ScriptVecOptType *);
//   bool (*is_some)(struct ScriptVecOptType *);
//   struct ScriptVecType (*unwrap)(struct ScriptVecOptType *);
// } ScriptVecOptVTable;
// typedef struct ScriptVecOptType {
//   mol2_cursor_t cur;
//   ScriptVecOptVTable *t;
// } ScriptVecOptType;
// struct XudtWitnessInputType;
// typedef struct XudtWitnessInputVTable {
//   struct ScriptOptType (*owner_script)(struct XudtWitnessInputType *);
//   struct BytesOptType (*owner_signature)(struct XudtWitnessInputType *);
//   struct ScriptVecOptType (*raw_extension_data)(struct XudtWitnessInputType *);
//   struct BytesVecType (*extension_data)(struct XudtWitnessInputType *);
// } XudtWitnessInputVTable;
// typedef struct XudtWitnessInputType {
//   mol2_cursor_t cur;
//   XudtWitnessInputVTable *t;
// } XudtWitnessInputType;
// XudtWitnessInputVTable *GetXudtWitnessInputVTable(void) {
//   static XudtWitnessInputVTable s_vtable;
//   static int inited = 0;
//   if (inited) return &s_vtable;
//   s_vtable.owner_script = NULL;
//   s_vtable.owner_signature = NULL;
//   s_vtable.raw_extension_data = NULL;
//   s_vtable.extension_data = NULL;
//   return &s_vtable;
// }
// struct XudtWitnessInputType make_XudtWitnessInput(mol2_cursor_t *cur) {
//   XudtWitnessInputType ret;
//   ret.cur = *cur;
//   ret.t = GetXudtWitnessInputVTable();
//   return ret;
// }
// int load_raw_extension_data(uint8_t **var_data, uint32_t *var_len) {
//   int err = 0;
//   bool use_input_type = true;
//   WitnessArgsType witness_args;
//   err = make_cursor_from_witness(&witness_args, &use_input_type);
//   if (err != 0) {
//     return err;
//   }

//   BytesOptType input;
//   if (use_input_type) {
//     input = witness_args.t->input_type(&witness_args);
//   } else {
//     input = witness_args.t->output_type(&witness_args);
//   }

//   if (input.t->is_none(&input)) {
//     return ERROR_INVALID_MOL_FORMAT;
//   }

//   struct mol2_cursor_t bytes = input.t->unwrap(&input);
//   // convert Bytes to XudtWitnessInputType
//   XudtWitnessInputType witness_input = make_XudtWitnessInput(&bytes);
//   ScriptVecOptType script_vec =
//       witness_input.t->raw_extension_data(&witness_input);

//   uint8_t raw_extension_data[RAW_EXTENSION_SIZE];
//   memset(raw_extension_data, 0, RAW_EXTENSION_SIZE);
//   uint32_t read_len =
//       mol2_read_at(&script_vec.cur, raw_extension_data, RAW_EXTENSION_SIZE);
//   if (read_len != script_vec.cur.size) {
//     return ERROR_INVALID_MOL_FORMAT;
//   }

//   *var_data = raw_extension_data;
//   *var_len = read_len;
//   return err;
// }
