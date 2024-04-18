# Cell Script

Cell Script is a newly designed language for smart-contract programming on the UTXO chain, which currently supports CKB.

## Example

Here is an example of a simple cell script.
```
import "debug"
import "tx"
import "cell"

// main is the entry point of every cell script
function main() {
    tx.scriptVerify()
    var ins = tx.inputs()
    var outs = tx.outputs()

    var in_sum uint64
    var out_sum uint64

    for _, input := range ins {
        in_sum += input
        if in_sum < input {
            debug.Printf("input overflow")
            return 1
        }
    }

    for _, output := range outs {
        out_sum += output
        if out_sum < output {
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
export PATH="/opt/homebrew/bin:$PATH"
export PATH="/opt/homebrew/opt/llvm@16/bin:$PATH"
brew install --cask spike
make build
source install.sh

cell <file>.cell
cargo install --git https://github.com/nervosnetwork/ckb-standalone-debugger ckb-debugger
ckb-debugger --bin <file>
```
## To develop xUDT

ckb-c-stdlib use molecule 0.7.1
```
cargo install moleculec@0.7.1 --locked
```

## How to Deploy?
```
git clone git@github.com:cell-labs/cell-cli.git
cd cell-cli
npm install
npm install -g .
# 1、Copy the cell bin file to the Cell-Cli folder: exmple: helloworld
# 2、Enter your CKB private key in cell.config.js
cell-cli deploy ./helloworld
```

