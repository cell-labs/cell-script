clean:
	rm -rf parser
antlr:
	go generate ./...
dev:
	go run main.go