// CellScriptParser.g4
// $antlr-format alignTrailingComments true, columnLimit 150, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine false, allowShortBlocksOnASingleLine true, alignSemicolons hanging, alignColons hanging

parser grammar CellScriptParser;

options {
    tokenVocab = CellScriptLexer;
}

sourceFile
    : importStmt functionStmt EOF
    ;

// import statement
importStmt
    : (importDecl eos)*
    ;

// function statement
functionStmt
    : ((functionDecl | declaration) eos)* EOF
    ;

eos
    : EOF
    | EOS
    ;

importDecl
    : IMPORT IDENTIFIER
    ;

declaration
    : TYPE IDENTIFIER
    ;

functionDecl
    : FUNC IDENTIFIER typeParameters? signature body?
    ;

typeParameters
    : L_BRACKET typeParameterDecl (COMMA typeParameterDecl)* R_BRACKET
    ;

signature
    : BOOL
    ;

body
    : L_CURLY expression? R_CURLY
    ;

typeParameterDecl
    : TYPE IDENTIFIER
    ;

expression
    : expression op = ('*' | '/') expression
    | expression op = ('+' | '-') expression
    | NUMBER
    ;