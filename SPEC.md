## Primary Types

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
vector
union
function

## Statements

if
for
break
continue
return

## Packages

### built-in packages

#### tx

#### debug

Support limited print function. Formatting is not support.

#### cell

## Example

```
import "tx"
import "cell"

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

  return true;
}
```
