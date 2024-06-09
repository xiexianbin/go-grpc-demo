.PHONY: protoc
protoc:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	protoc --go_out=../../.. proto/*.proto
	protoc --go-grpc_out=../../.. proto/*.proto

.PHONY: tls
tls:
	openssl ecparam -genkey -name secp384r1 -out cmd/simple_tls/server.key
	openssl req -new -x509 -sha256 -key cmd/simple_tls/server.key -out cmd/simple_tls/server.crt -days 3650 -addext "subjectAltName = DNS:go-grpc-demo"

.PHONY: self-ca
self-ca:
	mkdir -p cmd/simple_ca/conf/{client,server}
	echo "create ca ..."
	openssl genrsa -out cmd/simple_ca/conf/ca.key 2048
	openssl req -new -x509 -days 7200 -key cmd/simple_ca/conf/ca.key -out cmd/simple_ca/conf/ca.pem

	echo "subjectAltName = @alt_names\n\n[alt_names]\nDNS.1 = go-grpc-demo" > cmd/simple_ca/conf/san.cnf

	echo "create server crt ..."
	openssl ecparam -genkey -name secp384r1 -out cmd/simple_ca/conf/server/server.key
	openssl req -new -key cmd/simple_ca/conf/server/server.key -out cmd/simple_ca/conf/server/server.csr # -addext "subjectAltName = DNS:go-grpc-demo"
	openssl x509 -req -sha256 -CA cmd/simple_ca/conf/ca.pem -CAkey cmd/simple_ca/conf/ca.key -CAcreateserial -days 3650 -in cmd/simple_ca/conf/server/server.csr -out cmd/simple_ca/conf/server/server.crt -extfile cmd/simple_ca/conf/san.cnf

	echo "create client crt ..."
	openssl ecparam -genkey -name secp384r1 -out cmd/simple_ca/conf/client/client.key
	openssl req -new -key cmd/simple_ca/conf/client/client.key -out cmd/simple_ca/conf/client/client.csr # -addext "subjectAltName = DNS:go-grpc-demo"
	openssl x509 -req -sha256 -CA cmd/simple_ca/conf/ca.pem -CAkey cmd/simple_ca/conf/ca.key -CAcreateserial -days 3650 -in cmd/simple_ca/conf/client/client.csr -out cmd/simple_ca/conf/client/client.crt -extfile cmd/simple_ca/conf/san.cnf
