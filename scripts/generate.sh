#!/bin/sh

antlr4 -Dlanguage=Go -no-visitor -package lexer  -o ../lexer  CellScriptLexer.g4
antlr4 -Dlanguage=Go -no-visitor -package parser -o ../parser CellScriptParser.g4 \
       -lib .antlr -visitor