// CellsScriptLexer.g4
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
MAIN     : 'main';

// Skip rules
WHITESPACE: [ \r\n\t]+ -> skip;

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

// Literal
LITERAL    : STRING_LIT | BOOL_LIT | NUMBER;
NUMBER     : [0-9]+;
STRING_LIT : '"' IDENTIFIER '"';
BOOL_LIT   : 'true' | 'false';

// End of statement
EOS: ([\r\n]+ | ';' | '/*' .*? '*/' | EOF);