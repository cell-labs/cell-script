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

The internal discussion can be found [here](./DISCUSSION.md). 

## How to build (MacOS)

```
brew install go@1.22
brew install antlr@4
brew install openjdk@21
brew install llvm@16
brew tap riscv-software-src/riscv
brew install riscv-tools
export PATH=/opt/homebrew/bin:$PATH
brew install --cask spike
make build

cell <file>.cell
cargo install --git https://github.com/nervosnetwork/ckb-standalone-debugger ckb-debugger
ckb-debugger --bin <file>
```
