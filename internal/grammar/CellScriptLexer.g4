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

// Skip rules
WHITESPACE: [ \r\n\t]+ -> skip;

// Tokens
L_CURLY    : '{';
R_CURLY    : '}';
L_BRACKET  : '(';
R_BRACKET  : ')';
COMMA      : ',';
NUMBER     : [0-9]+;
IDENTIFIER : [a-zA-Z_][a-zA-Z0-9_]*;
STRING     : '"' IDENTIFIER '"';
TYPE       : 'int';
BOOL       : 'bool';
ADD        : '+';
SUB        : '-';
MUL        : '*';
DIV        : '/';

// End of statement
EOS: ([\r\n]+ | ';' | '/*' .*? '*/' | EOF);