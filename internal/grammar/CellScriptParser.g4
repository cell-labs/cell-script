// CellScriptParser.g4
// $antlr-format alignTrailingComments true, columnLimit 150, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine false, allowShortBlocksOnASingleLine true, alignSemicolons hanging, alignColons hanging

parser grammar CellScriptParser;

options {
    tokenVocab = CellScriptLexer;
}

sourceFile
    : importStmt declaration EOF
    ;

// import statement
importStmt
    : (importDecl eos)*
    ;

// declaration
declaration
    : ((functionDecl | varDecl) eos)* EOF
    ;

eos
    : EOF
    | EOS
    ;

importDecl
    : IMPORT IDENTIFIER
    ;

varDecl
    : VAR TYPE IDENTIFIER
    ;

functionDecl
    : FUNC IDENTIFIER typeParameters? signature block?
    ;

typeParameters
    : L_BRACKET typeParameterDecl? (COMMA typeParameterDecl)* R_BRACKET
    ;

signature
    : TYPE
    ;

block
    : L_CURLY expression? R_CURLY
    ;

typeParameterDecl
    : TYPE IDENTIFIER
    ;

expression
    : arithmeticExpr
    | returnExpr
    ;

arithmeticExpr
    : arithmeticExpr op = ('+' | '-') arithmeticExpr
    | arithmeticExpr op = ('*' | '/') arithmeticExpr
    | IDENTIFIER
    ;

returnExpr
    : RETURN (IDENTIFIER)
    ;

statement
    : declaration
    | simpleStmt
    | returnStmt
    | breakStmt
    | continueStmt
    | block
    | ifStmt
    | forStmt
    ;

simpleStmt
    : assignment
    | expressionStmt
    ;

expressionStmt
    : expression
    ;

assignment
    : expression assign_op expression
    ;

assign_op
    : (ADD | SUB | MUL | DIV | MOD)? ASSIGN
    ;

returnStmt
    : RETURN expression? eos
    ;

breakStmt
    : BREAK IDENTIFIER?
    ;

continueStmt
    : CONTINUE IDENTIFIER?
    ;

ifStmt
    : IF (expression | eos expression | simpleStmt eos expression) block (ELSE (ifStmt | block))?
    ;

forStmt
    : FOR (expression? | forClause) block
    ;

forClause
    : initStmt = simpleStmt? eos expression? eos postStmt = simpleStmt?
    ;