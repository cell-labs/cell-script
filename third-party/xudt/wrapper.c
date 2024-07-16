#include "ckb_syscalls.h"
#include "molecule/blockchain.h"

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
typedef struct
{
  int64_t size;
  uint8_t *ptr;
} String;
uint8_t* haha = (uint8_t*)"dispickableMe@4";
String mol_seg_to_string(mol_seg_t *seg)
{
  String s;
  s.size = 15;
  s.ptr = haha;
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

String parse_args() {
  mol_seg_t args = parse_xudt_args_as_mol_seg();
  return mol_seg_to_string(&args);
}
