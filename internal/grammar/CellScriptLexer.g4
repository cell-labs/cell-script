/*
 * A CellScript grammar for ANTLR 4 from the Language Specification https://cellsript.io/spec
 */

// $antlr-format alignTrailingComments true, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine true, allowShortBlocksOnASingleLine true, minEmptyLines 0, alignSemicolons ownLine
// $antlr-format alignColons trailing, singleLineOverrulesHangingColon true, alignLexerCommands true, alignLabels true, alignTrailers true

lexer grammar CellScriptLexer;

// Keywords
PACKAGE  : 'package';
IMPORT   : 'import';
FUNC     : 'function';
IF       : 'if';
ELSE     : 'else';
FOR      : 'for';
CONTINUE : 'continue';
BREAK    : 'break';
RETURN   : 'return';
TYPEDEF  : 'type';
VAR      : 'var';
CONST    : 'const';

// Types
TYPE: TYPE_TABLE | TYPE_VECTOR | TYPE_UNION | TYPE_OPTION | TYPE_SLICE;
// Primitive
TYPE_BOOL    : 'bool';
TYPE_UINT8   : 'uint8';
TYPE_UINT16  : 'uint16';
TYPE_UINT32  : 'uint32';
TYPE_UINT64  : 'uint64';
TYPE_UINT128 : 'uint128';
TYPE_UINT256 : 'uint256';
// Sequence
TYPE_VECTOR : 'vector';
TYPE_SLICE  : '[' TYPE ']';
// User-defined
TYPE_TABLE  : 'table';
TYPE_UNION  : 'union';
TYPE_OPTION : 'option';
// Function

// Identifiers
IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;

// Tokens
// Delimeters
L_CURLY      : '{';
R_CURLY      : '}';
L_SQURAE     : '[';
R_SQUARE     : ']';
L_PARENTHESE : '(';
R_PARENTHESE : ')';
// Puctuations
ADD        : '+';
SUB        : '-';
MUL        : '*';
DIV        : '/';
MOD        : '%';
XOR        : '^';
AND        : '&';
OR         : '|';
NOT        : '!';
ANDAND     : '&&';
OROR       : '||';
EQUAL      : '=';
EQUALEQUAL : '==';
NOTEQUAL   : '!=';
GT         : '>';
LT         : '<';
GE         : '>=';
LE         : '<=';
UNDERSCORE : '_';
DOT        : '.';
DOTDOT     : '..';
COMMA      : ',';
SEMI       : ';';
COLON      : ':'; // type seperator in varDecl

// Literal
LITERAL    : STRING_LIT | BOOL_LIT | NUMBER;
NUMBER     : [0-9]+;
STRING_LIT : '"' IDENTIFIER '"';
BOOL_LIT   : 'true' | 'false';

// String literals
RAW_STRING_LIT         : '`' ~'`'* '`';
INTERPRETED_STRING_LIT : '"' (~["\\] | ESCAPED_VALUE)* '"';

// Hidden tokens
WS           : [ \t\n\r]+                          -> skip;
COMMENT      : '/*' .*? '*/'                       -> skip;
EOS          : [\r\n]+ | ';' | '/*' .*? '*/' | EOF -> skip;
LINE_COMMENT : '//' ~[\r\n]*                       -> skip;

// Fragments
fragment ESCAPED_VALUE:
    '\\' ([abfnrtv\\'"] | OCTAL_DIGIT OCTAL_DIGIT OCTAL_DIGIT | 'x' HEX_DIGIT HEX_DIGIT)
;

fragment DECIMALS: [0-9] ('_'? [0-9])*;

fragment OCTAL_DIGIT: [0-7];

fragment HEX_DIGIT: [0-9a-fA-F];

fragment BIN_DIGIT: [01];

fragment LETTER: [a-zA-Z] | '_';