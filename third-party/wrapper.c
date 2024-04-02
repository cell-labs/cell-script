#include "ckb_syscalls.h"
#include "molecule/blockchain.h"

// We are limiting the script size loaded to be 32KB at most. This should be
// more than enough. We are also using blake2b with 256-bit hash here, which is
// the same as CKB.
const uint32_t BLAKE2B_BLOCK_SIZE = 32;
const uint32_t SCRIPT_SIZE = 32768;
const uint32_t MAX_CELLS = 16; //todo
const uint32_t MAX_DATA_SIZE = 4 * 1024 * 1024;

// Common error codes that might be returned by the script.
#define ERROR_ARGUMENTS_LEN -1
#define ERROR_ENCODING -2
#define ERROR_SYSCALL -3
#define ERROR_SCRIPT_TOO_LONG -21
#define ERROR_OVERFLOWING -51
#define ERROR_AMOUNT -52

typedef struct {
  uint64_t size;
  mol_seg_t* data;
}cell_data_t;
cell_data_t EMPTY_CELL_DATA = {0};

bool script_verify() {
  mol_seg_t script_seg;
  mol_seg_t args_seg;
  mol_seg_t args_bytes_seg;
  // First, let's load current running script, so we can extract owner lock
  // script hash from script args.
  unsigned char script[SCRIPT_SIZE];
  uint64_t len = SCRIPT_SIZE;
  int ret = ckb_load_script(script, &len, 0);
  if (ret != CKB_SUCCESS) {
    return false;
  }
  if (len > SCRIPT_SIZE) {
    return false;
  }
  
  script_seg.ptr = (uint8_t *)script;
  script_seg.size = len;
  if (MolReader_Script_verify(&script_seg, false) != MOL_OK) {
    return false;
  }
  args_seg = MolReader_Script_get_args(&script_seg);
  args_bytes_seg = MolReader_Bytes_raw_bytes(&args_seg);
  if (args_bytes_seg.size != BLAKE2B_BLOCK_SIZE) {
    return false;
  }
  return true;
}

uint64_t get_input_cell_data_len(int i) {
  uint64_t len = 0;
  int ret =
      ckb_load_cell_data(NULL, &len, 0, i, CKB_SOURCE_GROUP_INPUT);
  if (ret == CKB_INDEX_OUT_OF_BOUND) {
    return 0;
  }
  if (ret != CKB_SUCCESS) {
    return 0;
  }
  if (len < 16) {
    return 0;
  }
  return len;
}

uint64_t get_output_cell_data_len(int i) {
  uint64_t len = 0;
  int ret =
      ckb_load_cell_data(NULL, &len, 0, i, CKB_SOURCE_GROUP_OUTPUT);
  if (ret == CKB_INDEX_OUT_OF_BOUND) {
    return 0;
  }
  if (ret != CKB_SUCCESS) {
    return 0;
  }
  if (len < 16) {
    return 0;
  }
  return len;
}

cell_data_t get_utxo_inputs() {
  cell_data_t inputs = {0};
  inputs.data = malloc(MAX_CELLS * sizeof(mol_seg_t));
  int i = 0;
  while (1) {
    uint64_t len = get_input_cell_data_len(i);
    uint8_t* data = (uint8_t*)malloc(len * sizeof(uint8_t));
    int ret = ckb_load_cell_data(data, &len, 0, i,
                             CKB_SOURCE_GROUP_INPUT);
    // When `CKB_INDEX_OUT_OF_BOUND` is reached, we know we have iterated
    // through all cells of current type.
    if (ret == CKB_INDEX_OUT_OF_BOUND) {
      break;
    }
    if (ret != CKB_SUCCESS) {
      return EMPTY_CELL_DATA;
    }
    if (len < 16) {
      return EMPTY_CELL_DATA;
    }
    if (i >= (int)MAX_CELLS) {
      for (int j = 0; j <= i; j++) {
        free(inputs.data[j].ptr);
      }
      free(inputs.data);
      return EMPTY_CELL_DATA;
    }
    inputs.data[i].size = len;
    inputs.data[i].ptr = data;
    i += 1;
    inputs.size = i;
  }
  return inputs;
}

cell_data_t get_utxo_outputs() {
  cell_data_t outputs = {0};
  outputs.data = malloc(MAX_CELLS * sizeof(mol_seg_t));
  int i = 0;
  while (1) {
    uint64_t len = get_output_cell_data_len(i);
    uint8_t* data = (uint8_t*)malloc(len * sizeof(uint8_t));
    int ret = ckb_load_cell_data(&data, &len, 0, i,
                             CKB_SOURCE_GROUP_OUTPUT);
    // When `CKB_INDEX_OUT_OF_BOUND` is reached, we know we have iterated
    // through all cells of current type.
    if (ret == CKB_INDEX_OUT_OF_BOUND) {
      break;
    }
    if (ret != CKB_SUCCESS) {
      return EMPTY_CELL_DATA;
    }
    if (len < 16) {
      return EMPTY_CELL_DATA;
    }
    if (i >= (int)MAX_CELLS) {
      for (int j = 0; j <= i; j++) {
        free(outputs.data[j].ptr);
      }
      free(outputs.data);
      return EMPTY_CELL_DATA;
    }
    outputs.data[i].size = len;
    outputs.data[i].ptr = data;
    i += 1;
    outputs.size = i;
  }
  return outputs;
}