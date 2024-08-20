# Introduction

This is the reference manual for the Cell Script programming language.

Cell Script is a smart-contract oriented programming language, and focused in UTXO model rather than other cases. Currently, CKB is supported. Other UTXO chains will be supported in the future.

All the smart contracts are constructed from packages which are the easy to maintain by the community.

The first princple is keep Cell Script simple and easy to use.

# Notation

The syntax is specified using a [variant](https://en.wikipedia.org/wiki/Wirth_syntax_notation) of Extended Backus-Naur Form (EBNF):

```ebnf
Syntax      = { Production } ;
Production  = production_name "=" [ Expression ] "." ;
Expression  = Term { "|" Term } ;
Term        = Factor { Factor } ;
Factor      = production_name | token [ "…" token ] | Group | Option | Repetition ;
Group       = "(" Expression ")" ;
Option      = "[" Expression "]" ;
Repetition  = "{" Expression "}" ;
```

Productions are expressions constructed from terms and the following operators, in increasing precedence:

```
|   alternation
()  grouping
[]  option (0 or 1 times)
{}  repetition (0 to n times)
```

# Lexical elements

## Source code representation

The source code is Unicode code points encoded in UTF-8.

## Comments

There are two forms of comments:

1. Line comments start with the character sequence `//` and stop at the end of the line.
2. Block comments start with the character sequence `/*` and stop with the first subsequent character sequence `*/`.

A comment cannot start inside string literal, or inside a comment. A general comment containing no newlines acts like a space. Any other comment acts like a newline.

## Keywords

```c
break
bool
const
continue
else
extern
for
func
if
import
package
return
range
table
var
```

## Identifiers

Identifiers name program entities such as variables and types. An identifier is a sequence of one or more letters and digits. The first character in an identifier must be a letter in regex expression `[a-zA-Z_]`.

## Whitespace

Whitespace is defined as following Unicode code points:

```
U+0009 (horizontal tab, '\t')
U+000A (line feed, '\n')
U+000D (carriage return, '\r')
U+0020 (space, ' ')
```

All forms of whitespace serve only to separate tokens in the grammar, and have no semantic significance.
It means all whitespace can be replaced with other whitespace in a program.

## Tokens

Tokens form the vocabulary of Cell Script. There are four classes:

- [Keywords](#keywords)
- [Identifiers](#identifiers)
- [Operators and punctuation](#operators-and-punctuation)
- [Literals](#literal-expressions)

While breaking the input into tokens, the next token is the longest sequence of characters that form a valid token.

### Operators and punctuation

The following character sequences represent operators and punctuation.

```text
+    &     &&    ==    !=    (    )
-    |     ||    <     <=    [    ]
*    ^     >     ,     ;     {    }
/    <<    =     .     >=   
%    >>    !     :=
```

The formal syntax uses semicolons ";" as terminators in a number of productions.

# Module and source files

```ebnf
Module = Definition* ;
```

Source files has the extension `.cell`.
Each source file contains a sequence of zero or more [Declaration](#declarations) definitions.

## Prelude

TODO

## Main Function

A Module contains a main function can be compiled to an executable.

If a main function is present, it must take no arguments, and its return type must be `int64`.

# Declarations

```ebnf
Declaration = Package
            | Import
            | Constant
            | Variable
            | Function
            | Table;
```

## Packages

```ebnf
Package = package Identifier ;
```

## Imports

```ebnf
Import = import Identifier ;
```

## Constants

```ebnf
Constant = const Identifier Type (= Expression)? ;
```

## Variables

```ebnf
Variable = var Identifier Type (= Expression)? ;
```

## Functions

```ebnf
Function = func Identifier ( FunctionParameters? ) FunctionReturnType (BlockExpression)? ;
```

## Tables

```ebnf
Table = table Identifier { TableFields? } ;
TableFields = TableField (, TableField)* ,? ;
TableField = Identifier Type ;
```

## Unions(TODO)

# Expressions

## Literal Expressions

### Integer literals

Cell Script only supports unsigned integer to simplify the language itself.

An integer literal is a sequence of digits representing an integer constant. An optional prefix sets a non-decimal base: 0b for binary, 0x for hexadecimal. In hexadecimal literals, letters a through f and A through F represent values 10 through 15.

For readability, an underscore character _ may appear after a base prefix or between successive digits; such underscores do not change the literal's value.

```
int_lit        = binary_lit | hex_lit .
binary_lit     = "0b" [ "_" ] binary_digits .
hex_lit        = "0x" [ "_" ] hex_digits .

binary_digits  = binary_digit { [ "_" ] binary_digit } .
hex_digits     = hex_digit { [ "_" ] hex_digit } .
```

```
42
4_2
0600
0_600
0xBadFace
0xBad_Face
0x_67_7a_2f_cc_40_c6
170141183460469231731687303715884105727
170_141183_460469_231731_687303_715884_105727

_42         // an identifier, not an integer literal
42_         // invalid: _ must separate successive digits
4__2        // invalid: only one _ at a time
0_xBadFace  // invalid: _ must separate successive digits
```

### String literals

A string literal represents a string constant obtained from concatenating a sequence of characters.

```
string_lit         = " { unicode_char | newline } " .
```

These examples all represent the same string:

```
"日本語"                                 // UTF-8 input text
"\u65e5\u672c\u8a9e"                    // the explicit Unicode code points
"\U000065e5\U0000672c\U00008a9e"        // the explicit Unicode code points
"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // the explicit UTF-8 bytes
```

## Expression statements

An expression statement is one that evaluates an expression and ignores its result. As a rule, an expression statement's purpose is to trigger the effects of evaluating its expression.

# Types

A type determines a set of values together with operations and methods specific to those values. A type may be denoted by a type name, if it has one, which must be followed by type arguments if the type is generic. A type may also be specified using a type literal, which composes a type from existing types.

```
Type      = TypeName [ TypeArgs ] | TypeLit | "(" Type ")" .
TypeName  = identifier | QualifiedIdent .
TypeArgs  = "[" TypeList [ "," ] "]" .
TypeList  = Type { "," Type } .
TypeLit   = SliceType | TableType | FunctionType .
```

The language predeclares certain type names. Others are introduced with type declarations or type parameter lists. Composite types—slice, table, function types—may be constructed using type literals.

Predeclared types, defined types, and type parameters are called named types. An alias denotes a named type if the type given in the alias declaration is a named type.

# Primitive Types

Cell Script provides the following primitive types.

## Boolean types

A boolean type represents the set of Boolean truth values denoted by the predeclared constants true and false. The predeclared boolean type is `bool`; it is a defined type.

## Numeric types

Cell Sscript only has unsigned integers as numeric types. The predeclared architecture-independent numeric types are:

```
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
uint128     the set of all unsigned 128-bit integers (0 to 340282366920938463463374607431768211455)
uint256     (tbd)the set of all unsigned 256-bit integers (0 to 115792089237316195423570985008687907853269984665640564039457584007913129639935)
```

## String types

A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes. The number of bytes is called the length of the string and is never negative. Strings are immutable: once created, it is impossible to change the contents of a string. The predeclared string type is string; it is a defined type.

The length of a string s can be discovered using the built-in function len. The length is a compile-time constant if the string is a constant.

## Byte types

Byte is a simplify expression of uint8 for most cases.

### Reference types

A reference type denotes the set of all references to [variables](/##Variables) of a given type, called the *base type* of the reference.

```
ReferenceType = "*" BaseType . 
The syntax is specified using a variant of Extended Backus-Naur Form (EBNF):BaseType    = Type .
```

## Union types

TODO

## Option types

TODO

## Function types

A function type denotes the set of all functions with the same parameter and result types.

```
FunctionType   = "func" Signature .
Signature      = Parameters [ Result ] .
Result         = Parameters | Type .
Parameters     = "(" [ ParameterList [ "," ] ] ")" .
ParameterList  = ParameterDecl { "," ParameterDecl } .
ParameterDecl  = [ IdentifierList ] [ "..." ] Type .
```

Within a list of parameters or results, the names (IdentifierList) must all be present. Each name stands for one item (parameter or result) of the specified type and all non-blank names in the signature must be unique. Parameter and result lists are always parenthesized except that if there is exactly one unnamed result it may be written as an unparenthesized type.

```
func()
func(x uint8) uint8
func(a, _ uint32, z uint64) bool
```

# Complex Types

## Array types

An array is an ordered sequence of elements, all of the same type. This type is referred to as the element type. The number of elements in an array, known as its length, is always a non-negative integer.

Array Type Definition:

```
ArrayType   = "[" ArrayLength "]" ElementType .
ArrayLength = Expression .
ElementType = Type .
```

The length is an integral part of the array's type specification. It must be a non-negative constant that can be represented by an int. You can determine an array's length using the built-in len() function.

Array elements are accessed using zero-based indexing, with valid indices ranging from 0 to len(a)-1, where 'a' is the array.

While arrays are inherently one-dimensional, they can be nested to create multi-dimensional structures.

Examples:

```cellscript
[32]byte
[2*N] table { x uint64, y uint64 }
[3][5]int
```

An array type T cannot contain elements of type T, nor can it contain T as a component, either directly or indirectly, if the containing types are limited to arrays or tables.

```cellscript
// invalid array types
type (
	T1 [10]T1                 // element type of T1 is T1
	T2 [10]table{ f T2 }     // T2 contains T2 as component of a table
	T3 [10]T4                 // T3 contains T3 as component of a table in T4
	T4 table{ f T3 }         // T4 contains T4 as component of array T3 in a table
)

// valid array types
type (
	T5 [10]*T5                // T5 contains T5 as component of a pointer
	T6 [10]func() T6          // T6 contains T6 as component of a function type
	T7 [10]table{ f []T7 }   // T7 contains T7 as component of a slice in a table
)
```

## Slice types

An slice is a numbered sequence of elements of a single type, called the element type. The number of elements is called the length of the slice and is never negative.

```
SliceType   = "[" SliceLength "]" ElementType .
SliceLength = Expression .
ElementType = Type .
```

The length is part of the slice's type; it must evaluate to a non-negative constant representable by a value of type int. The length of slice a can be discovered using the built-in function size. The elements can be addressed by integer indices 0 through len(a)-1. Slice types are always one-dimensional but may be composed to form multi-dimensional types.

## Struct

TODO

## Table types

A table is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField). Within a table, non-blank field names must be unique.

```ebnf
Table = table Identifier { TableFields? } ;
TableFields = TableField (, TableField)* ,? ;
TableField = Identifier Type ;
```

# Blocks

A block is a possibly empty sequence of declarations and statements within matching brace brackets.

```
Block = "{" StatementList "}" .
StatementList = { Statement ";" } .
```

# Declarations and scope

A declaration binds a non-blank identifier to a constant, type, type parameter, variable, function, or package. Every identifier in a program must be declared. No identifier may be declared twice in the same block, and no identifier may be declared in both the file and package block.

The blank identifier may be used like any other identifier in a declaration, but it does not introduce a binding and thus is not declared. In the package block, the identifier init may only be used for init function declarations, and like the blank identifier it does not introduce a new binding.

```
Declaration   = ConstDecl | TypeDecl | VarDecl .
TopLevelDecl  = Declaration | FunctionDecl | MethodDecl .
```

The scope of a declared identifier is the extent of source text in which the identifier denotes the specified constant, type, variable, function, or package.

An identifier declared in a block may be redeclared in an inner block. While the identifier of the inner declaration is in scope, it denotes the entity declared by the inner declaration.

The package clause is not a declaration; the package name does not appear in any scope. Its purpose is to identify the files belonging to the same package and to specify the default package name for import declarations.

## Predeclared identifiers

The following identifiers are implicitly declared in the universe block

```
Types:
 bool byte string
 uint8 uint16 uint32 uint64 uint128 uint256

Constants:
 true false

Functions:
 append cap len
```

## Constant declarations

```
const a, b, c = 3, 4, "foo" // a = 3, b = 4, c = "foo", untyped integer and string constants
```

## Type declarations

A type declaration binds an identifier, the type name, to a type. Type declarations come in two forms: alias declarations and type definitions.

```
TypeDecl = "type" ( TypeSpec | "(" { TypeSpec ";" } ")" ) .
TypeSpec = AliasDecl | TypeDef .
```

## Variable declarations

A variable declaration creates one or more variables, binds corresponding identifiers to them, and gives each a type and an initial value.

```
VarDecl     = "var" ( VarSpec | "(" { VarSpec ";" } ")" ) .
VarSpec     = IdentifierList ( Type [ "=" ExpressionList ] | "=" ExpressionList ) .
```

```
var i uint8
var i uint16
var i uint32
var i uint64
var i uint128
var i uint256

var i byte

```

A short variable declaration uses the syntax:

```
ShortVarDecl = IdentifierList ":=" ExpressionList .
```

```
i, j := 0, 10
f := func() int { return 7 }
```

## Function declarations

A function declaration binds an identifier, the function name, to a function.

```
FunctionDecl = "func" FunctionName [ TypeParameters ] Signature [ FunctionBody ] .
FunctionName = identifier .
FunctionBody = Block .
```

If the function's signature declares result parameters, the function body's statement list must end in a terminating statement.

```
func min(x uint8, y uint8) uint8{
 if x < y {
  return x
 }
 return y
}
```

# Expressions

An expression specifies the computation of a value by applying operators and functions to operands.

## Operands

Operands denote the elementary values in an expression. An operand may be a literal, a (possibly qualified) non-blank identifier denoting a constant, variable, or function, or a parenthesized expression.

```
Operand     = Literal | OperandName [ TypeArgs ] | "(" Expression ")" .
Literal     = BasicLit | CompositeLit | FunctionLit .
BasicLit    = int_lit | byte_lit | string_lit .
OperandName = identifier | QualifiedIdent .
```

## Qualified identifiers

A qualified identifier is an identifier qualified with a package name prefix. Both the package name and the identifier must not be blank.

```
QualifiedIdent = PackageName "." identifier .
```

A qualified identifier accesses an identifier in a different package, which must be imported. The identifier must be exported and declared in the package block of that package.

```
math.Sin // denotes the Sin function in package math
```

## Index expressions

A primary expression of the form

```
a[x]
```

# Statements

## Assignment statements

An assignment replaces the current value stored in a variable with a new value specified by an expression. An assignment statement may assign a single value to a single variable, or multiple values to a matching number of variables.

```
Assignment = ExpressionList assign_op ExpressionList .

assign_op = [ add_op | mul_op ] "=" .
```

Each left-hand side operand must be addressable, a map index expression, or (for = assignments only) the blank identifier. Operands may be parenthesized.

```cellscript
x = 1
a[i] = 23
```

```cellscript
a, b = b, a  // exchange a and b

x := []int{1, 2, 3}
i := 0
i, x[i] = 1, 2  // set i = 1, x[0] = 2

i = 0
x[i], i = 2, 1  // set x[0] = 2, i = 1

x[0], x[0] = 1, 2  // set x[0] = 1, then x[0] = 2 (so x[0] == 2 at end)

x[1], x[3] = 4, 5  // set x[1] = 4, then panic setting x[3] = 5.

type Point table { x, y int }
var p Point
x[2], p.x = 6, 7  // set x[2] = 6, then panic setting p.x = 7

i = 2
x = []int{3, 5, 7}
for i, x[i] = range x {  // set i, x[2] = 0, x[0]
 break
}
// after this loop, i == 0 and x is []int{3, 5, 3}
```

## If statements

"If" statements specify the conditional execution of two branches according to the value of a boolean expression. If the expression evaluates to true, the "if" branch is executed, otherwise, if present, the "else" branch is executed.

```
IfStmt = "if" [ SimpleStmt ";" ] Expression Block [ "else" ( IfStmt | Block ) ] .
```

```cellscript
if x > max {
 x = max
}
```

## For statements

A "for" statement specifies repeated execution of a block. There are three forms: The iteration may be controlled by a single condition, a "for" clause, or a "range" clause.

```cellscript
for a < b {
 a *= 2
}
```

```cellscript
for i := 0; i < 10; i++ {
 f(i)
}
```

```cellscript
var a [10]string
for i, s := range a {
 // type of i is int
 // type of s is string
 // s == a[i]
 g(i, s)
}

```

## Break statements

A "break" statement terminates execution of the innermost "for" statement within the same function.

```
BreakStmt = "break".
```

```cellscript
for i = 0; i < n; i++ {
  for j = 0; j < m; j++ {
   if a[i][j] == nil {
    state = Error
    break 
   } else if a[i][j] == item {
    state = Found
    break
   }
  }
 }
```

## Continue statements

A "continue" statement begins the next iteration of the innermost enclosing "for" loop by advancing control to the end of the loop block. The "for" loop must be within the same function.

```
ContinueStmt = "continue" .
```

```cellscript
rows := [[1,2,3],[4,5,6]]
for y, row := range rows {
  for x, data := range row {
   if data % 2 == 0 {
    continue
   }
   row[x] = data + bias(x, y)
  }
 }
```

## Return statements

A "return" statement in a function F terminates the execution of F, and optionally provides one or more result values. Any functions deferred by F are executed before F returns to its caller.

```
ReturnStmt = "return" [ ExpressionList ] .
```

In a function without a result type, a "return" statement must not specify any result values.

```
func noResult() {
 return
}
```

# Built-in functions

Built-in functions are predeclared. They are called like any other function but some of them accept a type instead of an expression as the first argument.

The built-in functions do not have standard Go types, so they can only appear in call expressions; they cannot be used as function values.

TODO

## Append

TODO

## Cap

TODO

## Len

The built-in functions len take arguments of various types and return a result of type int. The implementation guarantees that the result always fits into an int.

```
Call      Argument type    Result

len(s)    string type      string length in bytes
          [n]T, *[n]T      array length (== n)
          []T              slice length

```

# Names

## tx

TODO

## debug

Support limited print function. Formatting is not support.

## cell

TODO

# Appendix

TODO
