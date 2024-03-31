clean:
	rm -rf parser
	rm -rf lexer
grammar: antlr
antlr:
	go generate ./...
dev:
	go run main.go
build:
	go build -o cell
tests:
	./cell || \
	./cell test/cases/hi.cell