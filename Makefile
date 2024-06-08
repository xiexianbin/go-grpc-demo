.PHONY: protoc
protoc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	protoc --go_out=../../.. proto/*.proto
	protoc --go-grpc_out=../../.. proto/*.proto
