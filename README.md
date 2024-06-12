# Cell Script

Cell Script is a newly designed language for smart-contract programming on the UTXO chain, which currently supports CKB.

## Example

Here is an example of a simple cell script.
```
import "debug"
import "tx"
import "cell"

function main() {
    if tx.isOwnerMode() {
        return 0
    }

    in_sum, out_sum := 0, 0
    ins := tx.inputs()
    if len(ins) == int32(0) {
        return 1
    }
    for input := range tx.inputs() {
        in_sum += input
    }

    for output := range tx.outputs() {
        out_sum += output
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
sudo mkdir /opt/riscv
./configure --prefix=/opt/riscv
sudo make linux

echo 'export PATH="$PATH:/opt/riscv/bin"' >> ~/.bashrc

source ~/.bashrc
```

Install `spike`

```
sudo apt install device-tree-compiler libboost-all-dev

cd ~/risc-v
git clone https://github.com/riscv-software-src/riscv-isa-sim.git

cd riscv-isa-sim

mkdir build
cd build

../configure --prefix=/opt/riscv
make
make install
```


Build & Install compiler in the root folder
```
make build
source install.sh
```



## How to compile cellscript program

You can write your program in file with `.cell`. Here for example, we say `hello.cell`

```
cell hello.cell
```

You will get the output file also known as elf file named `hello` in the root folder.

To run the elf file, we need `ckb-debugger` the default debugger for ckb programs.

Run the following command to install `ckb-debugger`, suppose you already know what is cargo.

```
cargo install --git https://github.com/nervosnetwork/ckb-standalone-debugger ckb-debugger
```

Run below command to debug & run the `hello` program.
```
ckb-debugger --bin hello
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

