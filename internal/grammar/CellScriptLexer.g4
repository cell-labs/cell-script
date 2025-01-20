/*
 * A CellScript grammar for ANTLR 4 from the Language Specification https://cellsript.io/spec
 */

// $antlr-format alignTrailingComments true, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine true, allowShortBlocksOnASingleLine true, minEmptyLines 0, alignSemicolons ownLine
// $antlr-format alignColons trailing, singleLineOverrulesHangingColon true, alignLexerCommands true, alignLabels true, alignTrailers true

lexer grammar CellScriptLexer;

// Keywords
FUNC     : 'function';
PACKAGE  : 'package';
IF       : 'if';
ELSE     : 'else';
FOR      : 'for';
CONTINUE : 'continue';
BREAK    : 'break';
IMPORT   : 'import';
RETURN   : 'return';
VECTOR   : 'vector';
TABLE    : 'table';
UNION    : 'union';
VAR      : 'var';

// Tokens
L_CURLY    : '{';
R_CURLY    : '}';
L_BRACKET  : '(';
R_BRACKET  : ')';
COMMA      : ',';
IDENTIFIER : [a-zA-Z_][a-zA-Z0-9_]*;
TYPE       : TYPE_INT | TYPE_BOOL;
TYPE_INT   : 'int';
TYPE_BOOL  : 'bool';
ADD        : '+';
SUB        : '-';
MUL        : '*';
DIV        : '/';
MOD        : '%';
ASSIGN     : '=';

// Literal
LITERAL    : STRING_LIT | BOOL_LIT | NUMBER;
NUMBER     : [0-9]+;
STRING_LIT : '"' IDENTIFIER '"';
BOOL_LIT   : 'true' | 'false';

// String literals
RAW_STRING_LIT         : '`' ~'`'* '`';
INTERPRETED_STRING_LIT : '"' (~["\\] | ESCAPED_VALUE)* '"';

// Hidden tokens
WS           : [ \t]+        -> skip;
COMMENT      : '/*' .*? '*/' -> skip;
TERMINATOR   : [\r\n]+       -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

// Fragments
fragment ESCAPED_VALUE:
    '\\' (
        'u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
        | 'U' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
        | [abfnrtv\\'"]
        | OCTAL_DIGIT OCTAL_DIGIT OCTAL_DIGIT
        | 'x' HEX_DIGIT HEX_DIGIT
    )
;

fragment DECIMALS: [0-9] ('_'? [0-9])*;

fragment OCTAL_DIGIT: [0-7];

fragment HEX_DIGIT: [0-9a-fA-F];

fragment BIN_DIGIT: [01];

fragment LETTER: [a-zA-Z] | '_';

// Treat whitespace as normal
WS_DEFAULT: [ \t]+ -> skip;
// Ignore any comments that only span one line
COMMENT_DEFAULT      : '/*' ~[\r\n]*? '*/' -> skip;
LINE_COMMENT_DEFAULT : '//' ~[\r\n]*       -> skip;
// Emit an EOS token for any newlines, semicolon, multiline comments or the EOF and 
//return to normal lexing
EOS: ([\r\n]+ | ';' | '/*' .*? '*/' | EOF);