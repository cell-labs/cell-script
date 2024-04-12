
#ifndef _BLOCKCHAIN_MOL2_API2_H_
#define _BLOCKCHAIN_MOL2_API2_H_

#define MOLECULE2_API_VERSION_MIN 5000

#include "molecule2_reader.h"

#ifdef __cplusplus
extern "C" {
#endif /* __cplusplus */

// ----forward declaration--------
struct ScriptVecType;
struct ScriptVecVTable;
struct ScriptVecVTable *GetScriptVecVTable(void);
struct ScriptVecType make_ScriptVec(mol2_cursor_t *cur);
uint32_t ScriptVec_len_impl(struct ScriptVecType *);
struct ScriptType ScriptVec_get_impl(struct ScriptVecType *, uint32_t, bool *);
struct ScriptVecOptType;
struct ScriptVecOptVTable;
struct ScriptVecOptVTable *GetScriptVecOptVTable(void);
struct ScriptVecOptType make_ScriptVecOpt(mol2_cursor_t *cur);
bool ScriptVecOpt_is_none_impl(struct ScriptVecOptType *);
bool ScriptVecOpt_is_some_impl(struct ScriptVecOptType *);
struct ScriptVecType ScriptVecOpt_unwrap_impl(struct ScriptVecOptType *);
struct XudtWitnessInputType;
struct XudtWitnessInputVTable;
struct XudtWitnessInputVTable *GetXudtWitnessInputVTable(void);
struct XudtWitnessInputType make_XudtWitnessInput(mol2_cursor_t *cur);
struct ScriptOptType XudtWitnessInput_get_owner_script_impl(
    struct XudtWitnessInputType *);
struct BytesOptType XudtWitnessInput_get_owner_signature_impl(
    struct XudtWitnessInputType *);
struct ScriptVecOptType XudtWitnessInput_get_raw_extension_data_impl(
    struct XudtWitnessInputType *);
struct BytesVecType XudtWitnessInput_get_extension_data_impl(
    struct XudtWitnessInputType *);
struct RCRuleType;
struct RCRuleVTable;
struct RCRuleVTable *GetRCRuleVTable(void);
struct RCRuleType make_RCRule(mol2_cursor_t *cur);
mol2_cursor_t RCRule_get_smt_root_impl(struct RCRuleType *);
uint8_t RCRule_get_flags_impl(struct RCRuleType *);
struct RCCellVecType;
struct RCCellVecVTable;
struct RCCellVecVTable *GetRCCellVecVTable(void);
struct RCCellVecType make_RCCellVec(mol2_cursor_t *cur);
uint32_t RCCellVec_len_impl(struct RCCellVecType *);
mol2_cursor_t RCCellVec_get_impl(struct RCCellVecType *, uint32_t, bool *);
struct RCDataType;
struct RCDataVTable;
struct RCDataVTable *GetRCDataVTable(void);
struct RCDataType make_RCData(mol2_cursor_t *cur);
uint32_t RCData_item_id_impl(struct RCDataType *);
struct RCRuleType RCData_as_RCRule_impl(struct RCDataType *);
struct RCCellVecType RCData_as_RCCellVec_impl(struct RCDataType *);
struct SmtProofType;
struct SmtProofVTable;
struct SmtProofVTable *GetSmtProofVTable(void);
struct SmtProofType make_SmtProof(mol2_cursor_t *cur);
uint32_t SmtProof_len_impl(struct SmtProofType *);
uint8_t SmtProof_get_impl(struct SmtProofType *, uint32_t, bool *);
struct SmtProofEntryType;
struct SmtProofEntryVTable;
struct SmtProofEntryVTable *GetSmtProofEntryVTable(void);
struct SmtProofEntryType make_SmtProofEntry(mol2_cursor_t *cur);
uint8_t SmtProofEntry_get_mask_impl(struct SmtProofEntryType *);
mol2_cursor_t SmtProofEntry_get_proof_impl(struct SmtProofEntryType *);
struct SmtProofEntryVecType;
struct SmtProofEntryVecVTable;
struct SmtProofEntryVecVTable *GetSmtProofEntryVecVTable(void);
struct SmtProofEntryVecType make_SmtProofEntryVec(mol2_cursor_t *cur);
uint32_t SmtProofEntryVec_len_impl(struct SmtProofEntryVecType *);
struct SmtProofEntryType SmtProofEntryVec_get_impl(
    struct SmtProofEntryVecType *, uint32_t, bool *);
struct SmtUpdateItemType;
struct SmtUpdateItemVTable;
struct SmtUpdateItemVTable *GetSmtUpdateItemVTable(void);
struct SmtUpdateItemType make_SmtUpdateItem(mol2_cursor_t *cur);
mol2_cursor_t SmtUpdateItem_get_key_impl(struct SmtUpdateItemType *);
uint8_t SmtUpdateItem_get_packed_values_impl(struct SmtUpdateItemType *);
struct SmtUpdateItemVecType;
struct SmtUpdateItemVecVTable;
struct SmtUpdateItemVecVTable *GetSmtUpdateItemVecVTable(void);
struct SmtUpdateItemVecType make_SmtUpdateItemVec(mol2_cursor_t *cur);
uint32_t SmtUpdateItemVec_len_impl(struct SmtUpdateItemVecType *);
struct SmtUpdateItemType SmtUpdateItemVec_get_impl(
    struct SmtUpdateItemVecType *, uint32_t, bool *);
struct SmtUpdateActionType;
struct SmtUpdateActionVTable;
struct SmtUpdateActionVTable *GetSmtUpdateActionVTable(void);
struct SmtUpdateActionType make_SmtUpdateAction(mol2_cursor_t *cur);
struct SmtUpdateItemVecType SmtUpdateAction_get_updates_impl(
    struct SmtUpdateActionType *);
mol2_cursor_t SmtUpdateAction_get_proof_impl(struct SmtUpdateActionType *);
struct XudtDataType;
struct XudtDataVTable;
struct XudtDataVTable *GetXudtDataVTable(void);
struct XudtDataType make_XudtData(mol2_cursor_t *cur);
mol2_cursor_t XudtData_get_lock_impl(struct XudtDataType *);
struct BytesVecType XudtData_get_data_impl(struct XudtDataType *);

// ----definition-----------------
typedef struct ScriptVecVTable {
  uint32_t (*len)(struct ScriptVecType *);
  struct ScriptType (*get)(struct ScriptVecType *, uint32_t, bool *);
} ScriptVecVTable;
typedef struct ScriptVecType {
  mol2_cursor_t cur;
  ScriptVecVTable *t;
} ScriptVecType;

typedef struct ScriptVecOptVTable {
  bool (*is_none)(struct ScriptVecOptType *);
  bool (*is_some)(struct ScriptVecOptType *);
  struct ScriptVecType (*unwrap)(struct ScriptVecOptType *);
} ScriptVecOptVTable;
typedef struct ScriptVecOptType {
  mol2_cursor_t cur;
  ScriptVecOptVTable *t;
} ScriptVecOptType;

typedef struct XudtWitnessInputVTable {
  struct ScriptOptType (*owner_script)(struct XudtWitnessInputType *);
  struct BytesOptType (*owner_signature)(struct XudtWitnessInputType *);
  struct ScriptVecOptType (*raw_extension_data)(struct XudtWitnessInputType *);
  struct BytesVecType (*extension_data)(struct XudtWitnessInputType *);
} XudtWitnessInputVTable;
typedef struct XudtWitnessInputType {
  mol2_cursor_t cur;
  XudtWitnessInputVTable *t;
} XudtWitnessInputType;

typedef struct RCRuleVTable {
  mol2_cursor_t (*smt_root)(struct RCRuleType *);
  uint8_t (*flags)(struct RCRuleType *);
} RCRuleVTable;
typedef struct RCRuleType {
  mol2_cursor_t cur;
  RCRuleVTable *t;
} RCRuleType;

typedef struct RCCellVecVTable {
  uint32_t (*len)(struct RCCellVecType *);
  mol2_cursor_t (*get)(struct RCCellVecType *, uint32_t, bool *);
} RCCellVecVTable;
typedef struct RCCellVecType {
  mol2_cursor_t cur;
  RCCellVecVTable *t;
} RCCellVecType;

typedef struct RCDataVTable {
  uint32_t (*item_id)(struct RCDataType *);
  struct RCRuleType (*as_RCRule)(struct RCDataType *);
  struct RCCellVecType (*as_RCCellVec)(struct RCDataType *);
} RCDataVTable;
typedef struct RCDataType {
  mol2_cursor_t cur;
  RCDataVTable *t;
} RCDataType;

typedef struct SmtProofVTable {
  uint32_t (*len)(struct SmtProofType *);
  uint8_t (*get)(struct SmtProofType *, uint32_t, bool *);
} SmtProofVTable;
typedef struct SmtProofType {
  mol2_cursor_t cur;
  SmtProofVTable *t;
} SmtProofType;

typedef struct SmtProofEntryVTable {
  uint8_t (*mask)(struct SmtProofEntryType *);
  mol2_cursor_t (*proof)(struct SmtProofEntryType *);
} SmtProofEntryVTable;
typedef struct SmtProofEntryType {
  mol2_cursor_t cur;
  SmtProofEntryVTable *t;
} SmtProofEntryType;

typedef struct SmtProofEntryVecVTable {
  uint32_t (*len)(struct SmtProofEntryVecType *);
  struct SmtProofEntryType (*get)(struct SmtProofEntryVecType *, uint32_t,
                                  bool *);
} SmtProofEntryVecVTable;
typedef struct SmtProofEntryVecType {
  mol2_cursor_t cur;
  SmtProofEntryVecVTable *t;
} SmtProofEntryVecType;

typedef struct SmtUpdateItemVTable {
  mol2_cursor_t (*key)(struct SmtUpdateItemType *);
  uint8_t (*packed_values)(struct SmtUpdateItemType *);
} SmtUpdateItemVTable;
typedef struct SmtUpdateItemType {
  mol2_cursor_t cur;
  SmtUpdateItemVTable *t;
} SmtUpdateItemType;

typedef struct SmtUpdateItemVecVTable {
  uint32_t (*len)(struct SmtUpdateItemVecType *);
  struct SmtUpdateItemType (*get)(struct SmtUpdateItemVecType *, uint32_t,
                                  bool *);
} SmtUpdateItemVecVTable;
typedef struct SmtUpdateItemVecType {
  mol2_cursor_t cur;
  SmtUpdateItemVecVTable *t;
} SmtUpdateItemVecType;

typedef struct SmtUpdateActionVTable {
  struct SmtUpdateItemVecType (*updates)(struct SmtUpdateActionType *);
  mol2_cursor_t (*proof)(struct SmtUpdateActionType *);
} SmtUpdateActionVTable;
typedef struct SmtUpdateActionType {
  mol2_cursor_t cur;
  SmtUpdateActionVTable *t;
} SmtUpdateActionType;

typedef struct XudtDataVTable {
  mol2_cursor_t (*lock)(struct XudtDataType *);
  struct BytesVecType (*data)(struct XudtDataType *);
} XudtDataVTable;
typedef struct XudtDataType {
  mol2_cursor_t cur;
  XudtDataVTable *t;
} XudtDataType;

#ifndef MOLECULEC_C2_DECLARATION_ONLY

// ----implementation-------------
struct ScriptVecType make_ScriptVec(mol2_cursor_t *cur) {
  ScriptVecType ret;
  ret.cur = *cur;
  ret.t = GetScriptVecVTable();
  return ret;
}
struct ScriptVecVTable *GetScriptVecVTable(void) {
  static ScriptVecVTable s_vtable;
  static int inited = 0;
  if (inited) return &s_vtable;
  s_vtable.len = ScriptVec_len_impl;
  s_vtable.get = ScriptVec_get_impl;
  return &s_vtable;
}
uint32_t ScriptVec_len_impl(ScriptVecType *this) {
  return mol2_dynvec_length(&this->cur);
}
ScriptType ScriptVec_get_impl(ScriptVecType *this, uint32_t index,
                              bool *existing) {
  ScriptType ret = {0};
  mol2_cursor_res_t res = mol2_dynvec_slice_by_index(&this->cur, index);
  if (res.errno != MOL2_OK) {
    *existing = false;
    return ret;
  } else {
    *existing = true;
  }
  ret.cur = res.cur;
  ret.t = GetScriptVTable();
  return ret;
}
struct ScriptVecOptType make_ScriptVecOpt(mol2_cursor_t *cur) {
  ScriptVecOptType ret;
  ret.cur = *cur;
  ret.t = GetScriptVecOptVTable();
  return ret;
}
struct ScriptVecOptVTable *GetScriptVecOptVTable(void) {
  static ScriptVecOptVTable s_vtable;
  static int inited = 0;
  if (inited) return &s_vtable;
  s_vtable.is_none = ScriptVecOpt_is_none_impl;
  s_vtable.is_some = ScriptVecOpt_is_some_impl;
  s_vtable.unwrap = ScriptVecOpt_unwrap_impl;
  return &s_vtable;
}
bool ScriptVecOpt_is_none_impl(ScriptVecOptType *this) {
  return mol2_option_is_none(&this->cur);
}
bool ScriptVecOpt_is_some_impl(ScriptVecOptType *this) {
  return !mol2_option_is_none(&this->cur);
}
ScriptVecType ScriptVecOpt_unwrap_impl(ScriptVecOptType *this) {
  ScriptVecType ret;
  mol2_cursor_t cur = this->cur;
  ret.cur = cur;
  ret.t = GetScriptVecVTable();
  return ret;
}
struct XudtWitnessInputType make_XudtWitnessInput(mol2_cursor_t *cur) {
  XudtWitnessInputType ret;
  ret.cur = *cur;
  ret.t = GetXudtWitnessInputVTable();
  return ret;
}
struct XudtWitnessInputVTable *GetXudtWitnessInputVTable(void) {
  static XudtWitnessInputVTable s_vtable;
  static int inited = 0;
  if (inited) return &s_vtable;
  s_vtable.owner_script = XudtWitnessInput_get_owner_script_impl;
  s_vtable.owner_signature = XudtWitnessInput_get_owner_signature_impl;
  s_vtable.raw_extension_data = XudtWitnessInput_get_raw_extension_data_impl;
  s_vtable.extension_data = XudtWitnessInput_get_extension_data_impl;
  return &s_vtable;
}
ScriptOptType XudtWitnessInput_get_owner_script_impl(
    XudtWitnessInputType *this) {
  ScriptOptType ret;
  mol2_cursor_t cur = mol2_table_slice_by_index(&this->cur, 0);
  ret.cur = cur;
  ret.t = GetScriptOptVTable();
  return ret;
}
BytesOptType XudtWitnessInput_get_owner_signature_impl(
    XudtWitnessInputType *this) {
  BytesOptType ret;
  mol2_cursor_t cur = mol2_table_slice_by_index(&this->cur, 1);
  ret.cur = cur;
  ret.t = GetBytesOptVTable();
  return ret;
}
ScriptVecOptType XudtWitnessInput_get_raw_extension_data_impl(
    XudtWitnessInputType *this) {
  ScriptVecOptType ret;
  mol2_cursor_t cur = mol2_table_slice_by_index(&this->cur, 2);
  ret.cur = cur;
  ret.t = GetScriptVecOptVTable();
  return ret;
}
BytesVecType XudtWitnessInput_get_extension_data_impl(
    XudtWitnessInputType *this) {
  BytesVecType ret;
  mol2_cursor_t cur = mol2_table_slice_by_index(&this->cur, 3);
  ret.cur = cur;
  ret.t = GetBytesVecVTable();
  return ret;
}
struct RCRuleType make_RCRule(mol2_cursor_t *cur) {
  RCRuleType ret;
  ret.cur = *cur;
  ret.t = GetRCRuleVTable();
  return ret;
}
struct RCRuleVTable *GetRCRuleVTable(void) {
  static RCRuleVTable s_vtable;
  static int inited = 0;
  if (inited) return &s_vtable;
  s_vtable.smt_root = RCRule_get_smt_root_impl;
  s_vtable.flags = RCRule_get_flags_impl;
  return &s_vtable;
}
mol2_cursor_t RCRule_get_smt_root_impl(RCRuleType *this) {
  mol2_cursor_t ret;
  mol2_cursor_t ret2 = mol2_slice_by_offset(&this->cur, 0, 32);
  ret = convert_to_array(&ret2);
  return ret;
}
uint8_t RCRule_get_flags_impl(RCRuleType *this) {
  uint8_t ret;
  mol2_cursor_t ret2 = mol2_slice_by_offset(&this->cur, 32, 1);
  ret = convert_to_Uint8(&ret2);
  return ret;
}
struct RCCellVecType make_RCCellVec(mol2_cursor_t *cur) {
  RCCellVecType ret;
  ret.cur = *cur;
  ret.t = GetRCCellVecVTable();
  return ret;
}
struct RCCellVecVTable *GetRCCellVecVTable(void) {
  static RCCellVecVTable s_vtable;
  static int inited = 0;
  if (inited) return &s_vtable;
  s_vtable.len = RCCellVec_len_impl;
  s_vtable.get = RCCellVec_get_impl;
  return &s_vtable;
}
uint32_t RCCellVec_len_impl(RCCellVecType *this) {
  return mol2_fixvec_length(&this->cur);
}
mol2_cursor_t RCCellVec_get_impl(RCCellVecType *this, uint32_t index,
                                 bool *existing) {
  mol2_cursor_t ret = {0};
  mol2_cursor_res_t res = mol2_fixvec_slice_by_index(&this->cur, 32, index);
  if (res.errno != MOL2_OK) {
    *existing = false;
    return ret;
  } else {
    *existing = true;
  }
  ret = convert_to_array(&res.cur);
  return ret;
}
struct RCDataType make_RCData(mol2_cursor_t *cur) {
  RCDataType ret;
  ret.cur = *cur;
  ret.t = GetRCDataVTable();
  return ret;
}
struct RCDataVTable *GetRCDataVTable(void) {
  static RCDataVTable s_vtable;
  static int inited = 0;
  if (inited) return &s_vtable;
  s_vtable.item_id = RCData_item_id_impl;
  s_vtable.as_RCRule = RCData_as_RCRule_impl;
  s_vtable.as_RCCellVec = RCData_as_RCCellVec_impl;
  return &s_vtable;
}
uint32_t RCData_item_id_impl(RCDataType *this) {
  return mol2_unpack_number(&this->cur);
}
RCRuleType RCData_as_RCRule_impl(RCDataType *this) {
  RCRuleType ret;
  mol2_union_t u = mol2_union_unpack(&this->cur);
  ret.cur = u.cursor;
  ret.t = GetRCRuleVTable();
  return ret;
}
RCCellVecType RCData_as_RCCellVec_impl(RCDataType *this) {
  RCCellVecType ret;
  mol2_union_t u = mol2_union_unpack(&this->cur);
  ret.cur = u.cursor;
  ret.t = GetRCCellVecVTable();
  return ret;
}
struct SmtProofType make_SmtProof(mol2_cursor_t *cur) {
  SmtProofType ret;
  ret.cur = *cur;
  ret.t = GetSmtProofVTable();
  return ret;
}
struct SmtProofVTable *GetSmtProofVTable(void) {
  static SmtProofVTable s_vtable;
  static int inited = 0;
  if (inited) return &s_vtable;
  s_vtable.len = SmtProof_len_impl;
  s_vtable.get = SmtProof_get_impl;
  return &s_vtable;
}
uint32_t SmtProof_len_impl(SmtProofType *this) {
  return mol2_fixvec_length(&this->cur);
}
uint8_t SmtProof_get_impl(SmtProofType *this, uint32_t index, bool *existing) {
  uint8_t ret = {0};
  mol2_cursor_res_t res = mol2_fixvec_slice_by_index(&this->cur, 1, index);
  if (res.errno != MOL2_OK) {
    *existing = false;
    return ret;
  } else {
    *existing = true;
  }
  ret = convert_to_Uint8(&res.cur);
  return ret;
}
struct SmtProofEntryType make_SmtProofEntry(mol2_cursor_t *cur) {
  SmtProofEntryType ret;
  ret.cur = *cur;
  ret.t = GetSmtProofEntryVTable();
  return ret;
}
struct SmtProofEntryVTable *GetSmtProofEntryVTable(void) {
  static SmtProofEntryVTable s_vtable;
  static int inited = 0;
  if (inited) return &s_vtable;
  s_vtable.mask = SmtProofEntry_get_mask_impl;
  s_vtable.proof = SmtProofEntry_get_proof_impl;
  return &s_vtable;
}
uint8_t SmtProofEntry_get_mask_impl(SmtProofEntryType *this) {
  uint8_t ret;
  mol2_cursor_t ret2 = mol2_table_slice_by_index(&this->cur, 0);
  ret = convert_to_Uint8(&ret2);
  return ret;
}
mol2_cursor_t SmtProofEntry_get_proof_impl(SmtProofEntryType *this) {
  mol2_cursor_t ret;
  mol2_cursor_t re2 = mol2_table_slice_by_index(&this->cur, 1);
  ret = convert_to_rawbytes(&re2);
  return ret;
}
struct SmtProofEntryVecType make_SmtProofEntryVec(mol2_cursor_t *cur) {
  SmtProofEntryVecType ret;
  ret.cur = *cur;
  ret.t = GetSmtProofEntryVecVTable();
  return ret;
}
struct SmtProofEntryVecVTable *GetSmtProofEntryVecVTable(void) {
  static SmtProofEntryVecVTable s_vtable;
  static int inited = 0;
  if (inited) return &s_vtable;
  s_vtable.len = SmtProofEntryVec_len_impl;
  s_vtable.get = SmtProofEntryVec_get_impl;
  return &s_vtable;
}
uint32_t SmtProofEntryVec_len_impl(SmtProofEntryVecType *this) {
  return mol2_dynvec_length(&this->cur);
}
SmtProofEntryType SmtProofEntryVec_get_impl(SmtProofEntryVecType *this,
                                            uint32_t index, bool *existing) {
  SmtProofEntryType ret = {0};
  mol2_cursor_res_t res = mol2_dynvec_slice_by_index(&this->cur, index);
  if (res.errno != MOL2_OK) {
    *existing = false;
    return ret;
  } else {
    *existing = true;
  }
  ret.cur = res.cur;
  ret.t = GetSmtProofEntryVTable();
  return ret;
}
struct SmtUpdateItemType make_SmtUpdateItem(mol2_cursor_t *cur) {
  SmtUpdateItemType ret;
  ret.cur = *cur;
  ret.t = GetSmtUpdateItemVTable();
  return ret;
}
struct SmtUpdateItemVTable *GetSmtUpdateItemVTable(void) {
  static SmtUpdateItemVTable s_vtable;
  static int inited = 0;
  if (inited) return &s_vtable;
  s_vtable.key = SmtUpdateItem_get_key_impl;
  s_vtable.packed_values = SmtUpdateItem_get_packed_values_impl;
  return &s_vtable;
}
mol2_cursor_t SmtUpdateItem_get_key_impl(SmtUpdateItemType *this) {
  mol2_cursor_t ret;
  mol2_cursor_t ret2 = mol2_slice_by_offset(&this->cur, 0, 32);
  ret = convert_to_array(&ret2);
  return ret;
}
uint8_t SmtUpdateItem_get_packed_values_impl(SmtUpdateItemType *this) {
  uint8_t ret;
  mol2_cursor_t ret2 = mol2_slice_by_offset(&this->cur, 32, 1);
  ret = convert_to_Uint8(&ret2);
  return ret;
}
struct SmtUpdateItemVecType make_SmtUpdateItemVec(mol2_cursor_t *cur) {
  SmtUpdateItemVecType ret;
  ret.cur = *cur;
  ret.t = GetSmtUpdateItemVecVTable();
  return ret;
}
struct SmtUpdateItemVecVTable *GetSmtUpdateItemVecVTable(void) {
  static SmtUpdateItemVecVTable s_vtable;
  static int inited = 0;
  if (inited) return &s_vtable;
  s_vtable.len = SmtUpdateItemVec_len_impl;
  s_vtable.get = SmtUpdateItemVec_get_impl;
  return &s_vtable;
}
uint32_t SmtUpdateItemVec_len_impl(SmtUpdateItemVecType *this) {
  return mol2_fixvec_length(&this->cur);
}
SmtUpdateItemType SmtUpdateItemVec_get_impl(SmtUpdateItemVecType *this,
                                            uint32_t index, bool *existing) {
  SmtUpdateItemType ret = {0};
  mol2_cursor_res_t res = mol2_fixvec_slice_by_index(&this->cur, 33, index);
  if (res.errno != MOL2_OK) {
    *existing = false;
    return ret;
  } else {
    *existing = true;
  }
  ret.cur = res.cur;
  ret.t = GetSmtUpdateItemVTable();
  return ret;
}
struct SmtUpdateActionType make_SmtUpdateAction(mol2_cursor_t *cur) {
  SmtUpdateActionType ret;
  ret.cur = *cur;
  ret.t = GetSmtUpdateActionVTable();
  return ret;
}
struct SmtUpdateActionVTable *GetSmtUpdateActionVTable(void) {
  static SmtUpdateActionVTable s_vtable;
  static int inited = 0;
  if (inited) return &s_vtable;
  s_vtable.updates = SmtUpdateAction_get_updates_impl;
  s_vtable.proof = SmtUpdateAction_get_proof_impl;
  return &s_vtable;
}
SmtUpdateItemVecType SmtUpdateAction_get_updates_impl(
    SmtUpdateActionType *this) {
  SmtUpdateItemVecType ret;
  mol2_cursor_t cur = mol2_table_slice_by_index(&this->cur, 0);
  ret.cur = cur;
  ret.t = GetSmtUpdateItemVecVTable();
  return ret;
}
mol2_cursor_t SmtUpdateAction_get_proof_impl(SmtUpdateActionType *this) {
  mol2_cursor_t ret;
  mol2_cursor_t re2 = mol2_table_slice_by_index(&this->cur, 1);
  ret = convert_to_rawbytes(&re2);
  return ret;
}
struct XudtDataType make_XudtData(mol2_cursor_t *cur) {
  XudtDataType ret;
  ret.cur = *cur;
  ret.t = GetXudtDataVTable();
  return ret;
}
struct XudtDataVTable *GetXudtDataVTable(void) {
  static XudtDataVTable s_vtable;
  static int inited = 0;
  if (inited) return &s_vtable;
  s_vtable.lock = XudtData_get_lock_impl;
  s_vtable.data = XudtData_get_data_impl;
  return &s_vtable;
}
mol2_cursor_t XudtData_get_lock_impl(XudtDataType *this) {
  mol2_cursor_t ret;
  mol2_cursor_t re2 = mol2_table_slice_by_index(&this->cur, 0);
  ret = convert_to_rawbytes(&re2);
  return ret;
}
BytesVecType XudtData_get_data_impl(XudtDataType *this) {
  BytesVecType ret;
  mol2_cursor_t cur = mol2_table_slice_by_index(&this->cur, 1);
  ret.cur = cur;
  ret.t = GetBytesVecVTable();
  return ret;
}
#endif  // MOLECULEC_C2_DECLARATION_ONLY

#ifdef __cplusplus
}
#endif /* __cplusplus */

#endif  // _BLOCKCHAIN_MOL2_API2_H_
