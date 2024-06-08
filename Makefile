.PHONY: protoc
protoc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	protoc --go_out=../../.. proto/*.proto
	protoc --go-grpc_out=../../.. proto/*.proto

.PHONY: tls
tls:
	openssl ecparam -genkey -name secp384r1 -out cmd/tls/server.key
	openssl req -new -x509 -sha256 -key cmd/tls/server.key -out cmd/tls/server.crt -days 3650 -addext "subjectAltName = DNS:go-grpc-demo"
