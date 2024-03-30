# Cell Script

Cell Script is a newly designed language for smart-contract programming on the UTXO chain, which currently supports CKB.

## Example

Here is an example of a simple cell script.
```
import "tx"
import "cell"
import "debug"

function main() {

  vector<cell> inputs = tx.inputs();
  vector<cell> outputs = tx.outputs();
  if(inputs.size() < outputs.size()) {
    return false;
  }

  for(cell input: inputs) {
    if(input.capacity < 100) {
      return true;
    }
  }

  uint8 idx = tx.scriptIndex("script hash");
  debug.log("find the script hash at cell idx");

  vector<vector<byte>> witness = tx.witness();
  for(vector<byte> w: witness) {
    debug.log("the witness data is", w);
  }
  
  return true;
}
```

The ideas and specs can be found [here](./SPEC.md). 

The internal discussions can be found [here](./DISCUSSION.md). 

