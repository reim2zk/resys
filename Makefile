build:
	go build -o bin/seq cmd/seq/main.go cmd/seq/io.go

seq:
	./bin/seq calc/seq/xor.test.json
