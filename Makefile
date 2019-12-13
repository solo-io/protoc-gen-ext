PHONY: protoc-gen-hash
protoc-gen-hash:
	go build -o _output/protoc-gen-hash protoc-gen-hash/main.go
