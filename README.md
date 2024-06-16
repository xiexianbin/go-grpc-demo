# go-grpc-demo

golang grpc demo

## Usage

### simple

```
# server
$ cd cmd/simple/server
$ go run ./main.go
2024/06/08 22:49:09 listen at 0.0.0.0:8000

# client
$ cd cmd/simple/client
$ go run main.go
2024/06/08 22:49:23 sum: 3
```

### stream server

```
# server
$ cd cmd/stream_server/server
$ go run ./main.go
2024/06/09 00:11:02 listen at 0.0.0.0:8000
2024/06/09 00:11:07 StreamService.Record resp pt: name:"callRecord"
2024/06/09 00:11:07 StreamService.Record resp pt: name:"callRecord"
2024/06/09 00:11:07 StreamService.Record resp pt: name:"callRecord"
2024/06/09 00:11:07 StreamService.Record resp pt: name:"callRecord"
2024/06/09 00:11:07 StreamService.Record resp pt: name:"callRecord"
2024/06/09 00:11:07 StreamService.Record resp pt: name:"callRecord"
2024/06/09 00:11:07 StreamService.Record resp pt: name:"callRecord"
2024/06/09 00:11:07 StreamService.Record resp pt: name:"callRecord"
2024/06/09 00:11:07 StreamService.Record resp pt: name:"callRecord"
2024/06/09 00:11:07 StreamService.Record resp pt: name:"callRecord"
2024/06/09 00:11:07 StreamService.Route pt: name:"callRoute"
2024/06/09 00:11:07 StreamService.Route pt: name:"callRoute"
2024/06/09 00:11:07 StreamService.Route pt: name:"callRoute"
2024/06/09 00:11:07 StreamService.Route pt: name:"callRoute"
2024/06/09 00:11:07 StreamService.Route pt: name:"callRoute"
2024/06/09 00:11:07 StreamService.Route pt: name:"callRoute"
2024/06/09 00:11:07 StreamService.Route pt: name:"callRoute"
2024/06/09 00:11:07 StreamService.Route pt: name:"callRoute"
2024/06/09 00:11:07 StreamService.Route pt: name:"callRoute"
2024/06/09 00:11:07 StreamService.Route pt: name:"callRoute"

# client
$ cd cmd/stream_server/client
$ go run ./main.go
2024/06/09 00:11:07 callList resp pt: name:"callList"
2024/06/09 00:11:07 callList resp pt: name:"callList" value:1
2024/06/09 00:11:07 callList resp pt: name:"callList" value:2
2024/06/09 00:11:07 callList resp pt: name:"callList" value:3
2024/06/09 00:11:07 callList resp pt: name:"callList" value:4
2024/06/09 00:11:07 callList resp pt: name:"callList" value:5
2024/06/09 00:11:07 callList resp pt: name:"callList" value:6
2024/06/09 00:11:07 callList resp pt: name:"callList" value:7
2024/06/09 00:11:07 callList resp pt: name:"callList" value:8
2024/06/09 00:11:07 callList resp pt: name:"callList" value:9
2024/06/09 00:11:07 callRecord resp pt: name:"gRPC Stream Server: Record" value:-1
2024/06/09 00:11:07 callRoute resp pt: name:"gPRC StreamService: Route"
2024/06/09 00:11:07 callRoute resp pt: name:"gPRC StreamService: Route" value:1
2024/06/09 00:11:07 callRoute resp pt: name:"gPRC StreamService: Route" value:2
2024/06/09 00:11:07 callRoute resp pt: name:"gPRC StreamService: Route" value:3
2024/06/09 00:11:07 callRoute resp pt: name:"gPRC StreamService: Route" value:4
2024/06/09 00:11:07 callRoute resp pt: name:"gPRC StreamService: Route" value:5
2024/06/09 00:11:07 callRoute resp pt: name:"gPRC StreamService: Route" value:6
2024/06/09 00:11:07 callRoute resp pt: name:"gPRC StreamService: Route" value:7
2024/06/09 00:11:07 callRoute resp pt: name:"gPRC StreamService: Route" value:8
2024/06/09 00:11:07 callRoute resp pt: name:"gPRC StreamService: Route" value:9
```

### tls

```
# generate tls key and cert to cmd/simple_tls/server.{key, crt}
$ make tls
...
Country Name (2 letter code) [AU]:
State or Province Name (full name) [Some-State]:
Locality Name (eg, city) []:
Organization Name (eg, company) [Internet Widgits Pty Ltd]:
Organizational Unit Name (eg, section) []:
Common Name (e.g. server FQDN or YOUR name) []:go-grpc-demo
Email Address []:

# server
$ cd cmd/simple_tls/server
$ go run ./main.go --help
  -help
        show help message
  -server-crt string
        server crt file path
  -server-key string
        server key file path
$ go run main.go -server-crt ../server.crt -server-key ../server.key

# client
$ cd cmd/simple_tls/client
$ go run main.go --help
  -client-crt string
        client crt file path
  -help
        show help message
$ go run main.go -client-crt ../server.crt
2024/06/09 00:50:11 version: Version:"v0.1.0"
2024/06/09 00:50:11 sum: Result:3
2024/06/09 00:50:11 diff: Result:-1
2024/06/09 00:50:11 fileContent: content:"..."
```

### ca & tls

```

$ tree cmd/simple_ca/conf
cmd/simple_ca/conf
├── ca.key
├── ca.pem
├── ca.srl
├── client
│   ├── client.crt
│   ├── client.csr
│   └── client.key
└── server
    ├── server.crt
    ├── server.csr
    └── server.key

3 directories, 9 files
```

#### by self-ca

```
make self-ca
mkdir -p cmd/simple_ca/conf/{client,server}
echo "create ca ..."
create ca ...
openssl genrsa -out cmd/simple_ca/conf/ca.key 2048
openssl req -new -x509 -days 7200 -key cmd/simple_ca/conf/ca.key -out cmd/simple_ca/conf/ca.pem
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:
State or Province Name (full name) [Some-State]:
Locality Name (eg, city) []:
Organization Name (eg, company) [Internet Widgits Pty Ltd]:
Organizational Unit Name (eg, section) []:
Common Name (e.g. server FQDN or YOUR name) []:
Email Address []:
echo "subjectAltName = @alt_names\n\n[alt_names]\nDNS.1 = go-grpc-demo" > cmd/simple_ca/conf/san.cnf
echo "create server crt ..."
create server crt ...
openssl ecparam -genkey -name secp384r1 -out cmd/simple_ca/conf/server/server.key
openssl req -new -key cmd/simple_ca/conf/server/server.key -out cmd/simple_ca/conf/server/server.csr # -addext "subjectAltName = DNS:go-grpc-demo"
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:
State or Province Name (full name) [Some-State]:
Locality Name (eg, city) []:
Organization Name (eg, company) [Internet Widgits Pty Ltd]:
Organizational Unit Name (eg, section) []:
Common Name (e.g. server FQDN or YOUR name) []:go-grpc-demo
Email Address []:

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:
An optional company name []:
openssl x509 -req -sha256 -CA cmd/simple_ca/conf/ca.pem -CAkey cmd/simple_ca/conf/ca.key -CAcreateserial -days 3650 -in cmd/simple_ca/conf/server/server.csr -out cmd/simple_ca/conf/server/server.crt -extfile cmd/simple_ca/conf/san.cnf
Certificate request self-signature ok
subject=C=AU, ST=Some-State, O=Internet Widgits Pty Ltd, CN=go-grpc-demo
echo "create client crt ..."
create client crt ...
openssl ecparam -genkey -name secp384r1 -out cmd/simple_ca/conf/client/client.key
openssl req -new -key cmd/simple_ca/conf/client/client.key -out cmd/simple_ca/conf/client/client.csr # -addext "subjectAltName = DNS:go-grpc-demo"
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:
State or Province Name (full name) [Some-State]:
Locality Name (eg, city) []:
Organization Name (eg, company) [Internet Widgits Pty Ltd]:
Organizational Unit Name (eg, section) []:
Common Name (e.g. server FQDN or YOUR name) []:go-grpc-demo
Email Address []:

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:
An optional company name []:
openssl x509 -req -sha256 -CA cmd/simple_ca/conf/ca.pem -CAkey cmd/simple_ca/conf/ca.key -CAcreateserial -days 3650 -in cmd/simple_ca/conf/client/client.csr -out cmd/simple_ca/conf/client/client.crt -extfile cmd/simple_ca/conf/san.cnf
Certificate request self-signature ok
subject=C=AU, ST=Some-State, O=Internet Widgits Pty Ltd, CN=go-grpc-demo
```

#### by xca

- creat TSL cert(option)

install and use [xca](https://github.com/x-ca/go-ca) to create tsl cert.

```
# 生成根证书
xca -create-ca true \
  -root-cert x-ca/simple_ca/root-ca.crt \
  -root-key x-ca/simple_ca/root-ca/private/root-ca.key \
  -tls-cert x-ca/simple_ca/tls-ca.crt \
  -tls-key x-ca/simple_ca/tls-ca/private/tls-ca.key

# 生成 server 证书
xca -cn server \
  --domains "localhost" \
  --ips 127.0.0.1 \
  -tls-cert x-ca/simple_ca/tls-ca.crt \
  -tls-key x-ca/simple_ca/tls-ca/private/tls-ca.key

# 生成 client 证书
xca -cn client \
  --domains "localhost" \
  --ips 127.0.0.1 \
  -tls-cert x-ca/simple_ca/tls-ca.crt \
  -tls-key x-ca/simple_ca/tls-ca/private/tls-ca.key
```

#### start server

```
# self-ca
$ cd cmd/simple_ca/server
$ go run ./main.go --help
  -ca-crt string
    	ca crt file path
  -help
    	show help message
  -server-crt string
    	server crt file path
  -server-key string
    	server key file path
$ go run ./main.go -ca-crt ../conf/ca.pem -server-crt ../conf/server/server.crt -server-key ../conf/server/server.key
2024/06/09 11:59:13 grpc server listen on [::]:8000

# xca
go run server.go -ca-crt ./x-ca/simple_ca/root-ca.crt -server-crt ./x-ca/certs/server/server.bundle.crt -server-key ./x-ca/certs/server/server.key
```

#### start client

```
# self-ca
$ cd cmd/simple_ca/client
$ go run ./main.go --help
  -ca-crt string
    	ca crt file path
  -client-crt string
    	client crt file path
  -client-key string
    	client key file path
  -help
    	show help message
$ go run ./main.go -ca-crt ../conf/ca.pem -client-crt ../conf/client/client.crt -client-key ../conf/client/client.key
2024/06/09 13:01:05 version: Version:"v0.1.0"
2024/06/09 13:01:05 sum: Result:3
2024/06/09 13:01:05 diff: Result:-1
2024/06/09 13:01:05 fileContent: content:"..."

# tsl
go run client.go -ca-crt ./x-ca/simple_ca/root-ca.crt -client-crt ./x-ca/certs/client/client.bundle.crt -client-key ./x-ca/certs/client/client.key
```

### simple_interceptor

```
$ cd cmd/simple_interceptor/server
$ go run main.go
2024/06/09 19:48:49 listen at 0.0.0.0:8000
2024/06/09 19:49:05 gRPC method: /proto.DemoService/Sum, req: Nums:1 Nums:2
2024/06/09 19:49:05 gRPC method: /proto.DemoService/Sum, resp: Result:3

# client
$ cd cmd/simple_interceptor/client
$ go run main.go
2024/06/09 19:49:05 sum: 3
```

### simple_http: gRPC & http Server

```
# generate tls key and cert to cmd/simple_tls/server.{key, crt}
$ make tls
...
Country Name (2 letter code) [AU]:
State or Province Name (full name) [Some-State]:
Locality Name (eg, city) []:
Organization Name (eg, company) [Internet Widgits Pty Ltd]:
Organizational Unit Name (eg, section) []:
Common Name (e.g. server FQDN or YOUR name) []:go-grpc-demo
Email Address []:

$ mv cmd/simple_tls/server.{key, crt} cmd/simple_http

# server
$ cd cmd/simple_http/server
$ go run ./main.go --help
  -help
        show help message
  -server-crt string
        server crt file path
  -server-key string
        server key file path
$ go run main.go -server-crt ../server.crt -server-key ../server.key
2024/06/09 20:29:32 application/grpc

# client
$ cd cmd/simple_http/client
$ go run main.go --help
  -client-crt string
        client crt file path
  -client-key string
        client key file path
  -help
        show help message
$ go run main.go -client-crt ../server.crt
2024/06/09 20:29:32 sum: 3

$ curl 127.0.0.1:8000
hello word!
```

### simple_auth

```
# server
$ cd cmd/simple_auth/server
$ go run ./main.go
2024/06/09 20:49:12 listen at 0.0.0.0:8000
token: {foo bar}

# client
$ cd cmd/simple_auth/client
$ go run main.go
2024/06/09 20:48:44 rpc error: code = Unauthenticated desc = bad key or secret
exit status 1
$ go run main.go
2024/06/09 20:49:15 sum: 3
```

### simple_deadline

```
# server
$ cd cmd/simple_deadline/server
$ go run ./main.go
2024/06/09 20:59:31 listen at 0.0.0.0:8000

# client
$ cd cmd/simple_deadline/client
$ go run ./main.go
2024/06/09 20:59:50 err: 4: context deadline exceeded
exit status 1
```

### simple_jaeger

- start jaeger

```
docker run --rm --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 4317:4317 \
  -p 4318:4318 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.57
```

```
# server
$ cd cmd/simple_jaeger/server
$ go run main.go
2024/06/10 20:10:43 listen at 0.0.0.0:8000

# client
$ cd cmd/simple_jaeger/client
$ go run main.go
2024/06/10 20:11:30 sum: 3
```

- open http://localhost:16686/search and search `Service: go-grpc-demo` get trace detail

## F&Q

### certificate relies on legacy Common Name field, use SANs instead

```
1. code = Unavailable desc = connection error: desc = "transport: authentication handshake failed: tls: failed to verify certificate: x509: certificate is not valid for any names, but wanted to match go-grpc-demo"

2. code = Unavailable desc = connection error: desc = "transport: authentication handshake failed: tls: failed to verify certificate: x509: certificate relies on legacy Common Name field, use SANs instead
```

- add `-addext "subjectAltName = DNS:go-grpc-demo"` when run `openssl req` to generate crt

### code = Unavailable desc = connection error: desc = "error reading server preface: http2: frame too large"

use TLS
