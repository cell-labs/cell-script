## Primary Types

array []  
uint8  
uint16  
uint32  
uint64  
uint128  
uint256  
byte  
string  
struct  
bool  

## Statements


## Control


## Example

```
import "tx"
import "cell"

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

  return true;
}
```
