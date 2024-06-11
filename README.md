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

## How to build (Ubuntu)

Install `golang 1.22`
```
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go
```

Install `antlr4`, `openjdk-21`
```
sudo apt install antlr4
sudo apt install openjdk-21-jdk
```

Install `llvm`
```
sudo bash -c "$(wget -O - https://apt.llvm.org/llvm.sh)"
```

Install `riscv-tools`

```
sudo apt update
sudo apt upgrade
sudo apt install autoconf automake autotools-dev curl python3 python3-pip libmpc-dev libmpfr-dev libgmp-dev gawk build-essential bison flex texinfo gperf libtool patchutils bc zlib1g-dev libexpat-dev ninja-build git cmake libglib2.0-dev libslirp-dev

mkdir risc-v
cd risc-v
git clone https://github.com/riscv-collab/riscv-gnu-toolchain.git

cd riscv-gnu-toolchain
./configure --prefix=/opt/riscv
make linux

echo 'export PATH="$PATH:/opt/riscv/bin"' >> ~/.bashrc

source ~/.bashrc
```

Install `spike`

```
sudo apt install device-tree-compiler libboost-regex-dev

cd ~/risc-v
git clone https://github.com/riscv-software-src/riscv-isa-sim.git

mkdir build
cd build
../configure --prefix=$RISCV
$ make
$ [sudo] make install
```


```
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

