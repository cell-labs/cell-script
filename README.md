# Cell Script

Cell Script is a newly designed language for smart-contract programming on the UTXO chain, which currently supports CKB.

The ideas and specs can be found [here](./SPEC.md). 

The internal discussion can be found [here](./DISCUSSION.md). 

How to build (MacOS)
```
brew install go@1.22
brew install antlr@4
brew install openjdk@21
make build

cell <file>.cell
```