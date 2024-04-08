#!/bin/sh

# $PWD is supposed to be internal/grammar
antlr4 -Dlanguage=Go -no-visitor -package lex  -o ../lex    CellScriptLexer.g4
antlr4 -Dlanguage=Go -visitor    -package parse -o ../parse CellScriptParser.g4 \
       -lib .antlr
cp ../walker.go ../parse
