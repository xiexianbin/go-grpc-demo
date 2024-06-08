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

## creat TSL cert(option)

install and use [xca](https://github.com/x-ca/go-ca) to create tsl cert.

```
# 生成根证书
xca -create-ca true \
  -root-cert x-ca/ca/root-ca.crt \
  -root-key x-ca/ca/root-ca/private/root-ca.key \
  -tls-cert x-ca/ca/tls-ca.crt \
  -tls-key x-ca/ca/tls-ca/private/tls-ca.key

# 生成 server 证书
xca -cn server \
  --domains "localhost" \
  --ips 127.0.0.1 \
  -tls-cert x-ca/ca/tls-ca.crt \
  -tls-key x-ca/ca/tls-ca/private/tls-ca.key

# 生成 client 证书
xca -cn client \
  --domains "localhost" \
  --ips 127.0.0.1 \
  -tls-cert x-ca/ca/tls-ca.crt \
  -tls-key x-ca/ca/tls-ca/private/tls-ca.key
```

## start server

```
# no tsl
go run server.go

# tsl
go run server.go -ca-crt ./x-ca/ca/root-ca.crt -server-crt ./x-ca/certs/server/server.bundle.crt -server-key ./x-ca/certs/server/server.key
```

## run client

```
# no tsl
go run client.go

# tsl
go run client.go -ca-crt ./x-ca/ca/root-ca.crt -client-crt ./x-ca/certs/client/client.bundle.crt -client-key ./x-ca/certs/client/client.key
```
