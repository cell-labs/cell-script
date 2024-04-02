# Cell Script

Cell Script is a newly designed language for smart-contract programming on the UTXO chain, which currently supports CKB.

## Example

Here is an example of a simple cell script.
```
//package main
import "debug"
import "tx"
import "cell"

// main is the entry point of every cell script
function main() {
    var inputs := tx.inputs()
    var outputs := tx.outputs()

    var in_sum, out_sum uint128

    for _, input := range inputs {
        in_sum += input.data.as(uint128)
        if in_sum < input.data.as(uint128) {
            debug.Printf("input overflow")
            return 1
        }
    }

    for _, output := range outputs {
        out_sum += output.data.as(uint128)
        if out_sum < input.data.as(uint128) {
            debug.Printf("output overflow")
            return 1
        }
    }

    if in_sum < out_sum {
        debug.Printf("Invalid Amount")
        return 1
    }
    
    return 0
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
source install.sh

cell <file>.cell
cargo install --git https://github.com/nervosnetwork/ckb-standalone-debugger ckb-debugger
ckb-debugger --bin <file>
```
