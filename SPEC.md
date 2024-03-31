# Introduction

This is the reference manual for the Cell Script programming language.

Cell Script is a smart-contract oriented programming language, and focused in UTXO model rather than other cases. Currently, CKB is supported. Other UTXO chains will be supported in the future.

All the smart contracts are constructed from packages which are the easy to maintain by the community.

The first princple is keep Cell Script simple and easy to use.  

# Notation

TODO

# Source code representation

The source code is Unicode text encoded in UTF-8, and the source code file is name as `.cell`. 

Each code point is distinct, including uppercase and lowercase letters. There are implementation restrictions: the NUL character (U+0000) may be disallowed, and a UTF-8-encoded byte order mark (U+FEFF) may be ignored if it's the first code point in the source.

## Characters



```
newline        = /* the Unicode code point U+000A */ .
unicode_char   = /* an arbitrary Unicode code point except newline */ .
unicode_letter = /* a Unicode code point categorized as "Letter" */ .
unicode_digit  = /* a Unicode code point categorized as "Number, decimal digit" */ .
```

## Letters and digits

```
letter        = unicode_letter | "_" .
decimal_digit = "0" … "9" .
binary_digit  = "0" | "1" .
octal_digit   = "0" … "7" .
hex_digit     = "0" … "9" | "A" … "F" | "a" … "f" .
```

# Lexical elements

## Comments

There are two forms of comments:
1. Line comments start with the character sequence // and stop at the end of the line.
2. General comments start with the character sequence /* and stop with the first subsequent character sequence */.

A comment cannot start inside string literal, or inside a comment. A general comment containing no newlines acts like a space. Any other comment acts like a newline.

## Tokens

Tokens form the vocabulary of Cell Script. There are four classes: identifiers, keywords, operators and punctuation, and literals. White space, formed from spaces (U+0020), horizontal tabs (U+0009), carriage returns (U+000D), and newlines (U+000A), is ignored except as it separates tokens that would otherwise combine into a single token. Also, a newline or end of file may trigger the insertion of a semicolon. While breaking the input into tokens, the next token is the longest sequence of characters that form a valid token.

## Semicolons

The formal syntax uses semicolons ";" as terminators in a number of productions.

## Identifiers

Identifiers name program entities such as variables and types. An identifier is a sequence of one or more letters and digits. The first character in an identifier must be a letter.

## Keywords

```
bool
break
byte
continue
const
else
for
function
if
import
return
string
uint8  
uint16  
uint32  
uint64  
uint128  
uint256  
table
union
var
vector
```

## Operators and punctuation

The following character sequences represent operators and punctuation.

```
+    &     &&    ==    !=    (    )
-    |     ||    <     <=    [    ]
*    ^     >     ,     ;     {    }
/    <<    =     .     >=   
%    >>    !      
```

## Integer literals

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


## String literals

A string literal represents a string constant obtained from concatenating a sequence of characters. 

```
string_lit         = "`" { unicode_char | newline } "`" .
```

These examples all represent the same string:

```
"日本語"                                 // UTF-8 input text
"\u65e5\u672c\u8a9e"                    // the explicit Unicode code points
"\U000065e5\U0000672c\U00008a9e"        // the explicit Unicode code points
"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // the explicit UTF-8 bytes
```

# Constants

There are boolean constants, integer constants, and string constants. integer, and complex constants are collectively called numeric constants.

TODO

# Variables

A variable is a storage location for holding a value. The set of permissible values is determined by the variable's type.

A variable declaration or, for function parameters and results, the signature of a function declaration or function literal reserves storage for a named variable.

The static type (or just type) of a variable is the type given in its declaration, the type provided in the new call or composite literal, or the type of an element of a structured variable.

```
var v *T           // v has value nil, static type *T
```

# Types

A type determines a set of values together with operations and methods specific to those values. A type may be denoted by a type name, if it has one, which must be followed by type arguments if the type is generic. A type may also be specified using a type literal, which composes a type from existing types.

```
Type      = TypeName [ TypeArgs ] | TypeLit | "(" Type ")" .
TypeName  = identifier | QualifiedIdent .
TypeArgs  = "[" TypeList [ "," ] "]" .
TypeList  = Type { "," Type } .
TypeLit   = VectorType | TableType | FunctionType .
```

The language predeclares certain type names. Others are introduced with type declarations or type parameter lists. Composite types—vector, table, function types—may be constructed using type literals.

Predeclared types, defined types, and type parameters are called named types. An alias denotes a named type if the type given in the alias declaration is a named type.

# Primitive Types

Cell Script provides the following primitive types.

## Boolean types
A boolean type represents the set of Boolean truth values denoted by the predeclared constants true and false. The predeclared boolean type is bool; it is a defined type.

## Numeric types

Cell Sscript only has unsigned integers as numeric types. The predeclared architecture-independent numeric types are:

```
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
uint128     the set of all unsigned 128-bit integers (0 to 340282366920938463463374607431768211455)
uint256     the set of all unsigned 256-bit integers (0 to 115792089237316195423570985008687907853269984665640564039457584007913129639935)
```

## String types
A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes. The number of bytes is called the length of the string and is never negative. Strings are immutable: once created, it is impossible to change the contents of a string. The predeclared string type is string; it is a defined type.

The length of a string s can be discovered using the built-in function len. The length is a compile-time constant if the string is a constant. 

## Byte types

Byte is a simplify expression of uint8 for most cases.

## Union types

TODO

## Option types

TODO

## Function types
A function type denotes the set of all functions with the same parameter and result types.

```
FunctionType   = "function" Signature .
Signature      = Parameters [ Result ] .
Result         = Parameters | Type .
Parameters     = "(" [ ParameterList [ "," ] ] ")" .
ParameterList  = ParameterDecl { "," ParameterDecl } .
ParameterDecl  = [ IdentifierList ] [ "..." ] Type .
```
Within a list of parameters or results, the names (IdentifierList) must all be present. Each name stands for one item (parameter or result) of the specified type and all non-blank names in the signature must be unique. Parameter and result lists are always parenthesized except that if there is exactly one unnamed result it may be written as an unparenthesized type.

```
function()
function(x uint8) uint8
function(a, _ uint32, z uint64) bool
```

# Complex Types

## Vector types
An vector is a numbered sequence of elements of a single type, called the element type. The number of elements is called the length of the vector and is never negative.

```
VectorType   = "[" VectorLength "]" ElementType .
VectorLength = Expression .
ElementType = Type .
```

The length is part of the vector's type; it must evaluate to a non-negative constant representable by a value of type int. The length of vector a can be discovered using the built-in function size. The elements can be addressed by integer indices 0 through len(a)-1. Vector types are always one-dimensional but may be composed to form multi-dimensional types.


## Table types

A table is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField). Within a table, non-blank field names must be unique.

TODO


# Blocks

A block is a possibly empty sequence of declarations and statements within matching brace brackets.

```
Block = "{" StatementList "}" .
StatementList = { Statement ";" } .
```

# Declarations and scope

A declaration binds a non-blank identifier to a constant, type, type parameter, variable, function, label, or package. Every identifier in a program must be declared. No identifier may be declared twice in the same block, and no identifier may be declared in both the file and package block.

The blank identifier may be used like any other identifier in a declaration, but it does not introduce a binding and thus is not declared. In the package block, the identifier init may only be used for init function declarations, and like the blank identifier it does not introduce a new binding.


```
Declaration   = ConstDecl | TypeDecl | VarDecl .
TopLevelDecl  = Declaration | FunctionDecl | MethodDecl .
```

The scope of a declared identifier is the extent of source text in which the identifier denotes the specified constant, type, variable, function, label, or package.

An identifier declared in a block may be redeclared in an inner block. While the identifier of the inner declaration is in scope, it denotes the entity declared by the inner declaration.

The package clause is not a declaration; the package name does not appear in any scope. Its purpose is to identify the files belonging to the same package and to specify the default package name for import declarations.

## Label scopes

TODO

## Predeclared identifiers

TODO

## Constant declarations

TODO

## Type declarations

TODO

## Variable declarations

A variable declaration creates one or more variables, binds corresponding identifiers to them, and gives each a type and an initial value.

```
VarDecl     = "var" ( VarSpec | "(" { VarSpec ";" } ")" ) .
VarSpec     = IdentifierList ( Type [ "=" ExpressionList ] | "=" ExpressionList ) .
```

```
var i uin8
```

## Function declarations

A function declaration binds an identifier, the function name, to a function.

```
FunctionDecl = "function" FunctionName [ TypeParameters ] Signature [ FunctionBody ] .
FunctionName = identifier .
FunctionBody = Block .
```

If the function's signature declares result parameters, the function body's statement list must end in a terminating statement.



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

# Statements

if
for
break
continue
return


TODO

# Built-in functions

TODO

# Packages


# Program initialization and execution


# Errors


# Misc


## tx


## debug

Support limited print function. Formatting is not support.

## cell


# Appendix
