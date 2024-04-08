#!/bin/sh

antlr4 -Dlanguage=Go -no-visitor -package lex  -o ../lex  CellScriptLexer.g4
antlr4 -Dlanguage=Go -no-visitor -package parse -o ../parse CellScriptParser.g4 \
       -lib .antlr -visitor