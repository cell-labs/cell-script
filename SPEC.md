## Primitive Types

array []  
uint8  
uint16  
uint32  
uint64  
uint128  
uint256  
byte  
string    
bool  
table  

## Statements


## Control


## Example

Here is an example of a simple cell script.
```
import "tx"
import "cell"
import "debug"

function main() {

  cell[] inputs = tx.inputs();
  cell[] outputs = tx.outputs();
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

  byte[][] witness = tx.witness();
  for(byte[] w: witness) {
    debug.log("the witness data is", w);
  }
  
  return true;
}
```
