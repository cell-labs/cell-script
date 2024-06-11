# CellScript

CellScript is as simple as Go language. If you know basic syntax of Go, it is easy to migrate to CellScript.

## Status and aimed timeline

It is currently alpha, experimental. We expect CellScript to reach *beta-preview* in 06/2024.

## Main advantages

- Simplify the development process of CKB.
- Support UTXO model natively.

## Overview

A CellScript file consists of the following parts:

* Package declaration
* Import packages
* Functions
* Statements and expressions

Look at the following code, to understand it better:

### Example

```solidity
package main
import "debug"

function main() {
  debug.Println("Hello World!")
}
```

## Statements

`debug.Println("Hello World!")` is a statement.

## Comments

A comment is a text that is ignored upon execution.

Comments can be used to explain the code, and to make it more readable.

Comments can also be used to prevent code execution when testing an alternative code.

CellScript supports single-line or multi-line comments.

### Single-line Comments

```
// This is a comment
package main
import ("fmt")

function main() {
  // This is a comment
  fmt.Println("Hello World!")
}
```

### Multi-line Comments

Multi-line comments start with `/*` and ends with `*/`. Any text between `/*` and `*/` will be ignored by the compiler.

## Variables

### Variable Types

* unsigned integer
* string
* bool
* table
* vector
* pointer(TEMPRORY)
* union(TODO)

### Variable Naming Rules

A variable can have a short name (like x and y) or a more descriptive name (age, price, carname, etc.).

Variable naming rules:

* A variable name must start with a letter or an underscore character (_)
* A variable name cannot start with a digit
* A variable name can only contain alpha-numeric characters and underscores (`a-z, A-Z`, `0-9`, and `_` )
* Variable names are case-sensitive (age, Age and AGE are three different variables)
* There is no limit on the length of the variable name
* A variable name cannot contain spaces
* The variable name cannot be any CellScript keywords

### Declaring (Creating) Variable

#### `var` keyword

```
var variablename type = value
```

#### `:=` sign

```
variablename := value
```

#### Variable Declaration With Initial Value

```
package main
import ("fmt")

func main() {
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
}
```

#### Variable Declaration Without Initial Value

In Go, all variables are initialized. So, if you declare a variable without an initial value, its value will be set to the default value of its type:

```
package main
import ("fmt")

function main() {
  var a string
  var b int64
  var c bool

  fmt.Println(a) // "0"
  fmt.Println(b) // 0
  fmt.Println(c) // false
}
```

## Constants

If a variable should have a fixed value that cannot be changed, you can use the `const` keyword.

The `const` keyword declares the variable as "constant", which means that it is  **unchangeable and read-only** .

`const CONSTNAME type = value`

### Declaring a Constant

```
package main
import ("fmt")

const PI = 3.14

function main() {
  fmt.Println(PI)
}
```

## Output Functions

There are three functions to output text:

* `Println()`
* `Printf()`

```
package main
import ("debug")

function main() {
  var i,j string = "Hello","World"

  debug.Print(i)
  debug.Print(j)
}
```

## Array types

Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.

```
var array_name = [length]datatype{values} // here length is defined
array_name := [length]datatype{values} // here length is defined
```

## Vector types

An vector is a numbered sequence of elements of a single type, called the element type. The number of elements is called the length of the vector and is never negative.

```
vector_name := []datatype{values}
```

## Operators

Operators are used to perform operations on variables and values.

## Conditions

Conditional statements are used to perform different actions based on different conditions.

A condition can be either `true` or `false`.

CellScript supports the usual comprison operators from mathematics:

* Less than `<`
* Less than or equal `<=`
* Greater than `>`
* Greater than or equal `>=`
* Equal to `==`
* Not equal to `!=`

Additionally,CellScript supports the usual logical operators:

* Logical AND `&&`
* Logical OR `||`
* Logical NOT `!`

## Functions

### Create a Function

To create (often referred to as declare) a function, do the following:

* Use the `function` keyword.
* Specify a name for the function, followed by parentheses ().
* Finally, add code that defines what the function should do, inside curly braces {}.

```
function FunctionName() {
  // code to be executed
}
```

### Call a Function

Functions are not executed immediately. They are "saved for later use", and will be executed when they are called.

```
package main
import ("debug")

function myMessage() {
  debug.Println("I just got executed!")
}

function main() {
  myMessage() // call the function
}
```

## For Loops

Loops are handy if you want to run the same code over and over again, each time with a different value.

Each execution of a loop is called an  **iteration** .

The `for` loop can take up to three statements:

```
for statement1; statement2; statement3 {
   // code to be executed for each iteration
}
```

Example

```
package main
import ("fmt")

func main() {
  for i:=0; i < 5; i++ {
    fmt.Println(i)
  }
}
```

## Table types

A table is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField). Within a table, non-blank field names must be unique.

To declare a structure in CellScript, use the `type` and `struct` keywords:

```
type struct_name struct {
  member1 datatype;
  member2 datatype;
  member3 datatype;
  ...
}
```

Example

```
type Person struct {
  name string
  age int
  job string
  salary int
}
```
