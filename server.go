/*
Copyright [2022] [xiexianbin.cn]

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
    "log"
    "net"
    "net/rpc"
    "net/rpc/jsonrpc"
    "os"

    "github.com/xiexianbin/go-rpc-demo/objects"
)

func main() {
    // Register RPC publishes
    err := rpc.Register(&objects.Calculate{})
    if err != nil {
        log.Println(err.Error())
        os.Exit(-1)
    }

    // Listener
    addr := ":8000"
    listener, err := net.Listen("tcp", addr)
    if err != nil {
        log.Println(err)
        os.Exit(-1)
    }
    log.Println("rpc server listen on", listener.Addr())
    defer func(listener net.Listener) {
        err := listener.Close()
        if err != nil {
            log.Println(err.Error())
            os.Exit(-1)
        }
        log.Println("rpc server closed.")
    }(listener)

    // listen actions
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println(err)
            return
        }
        go func(conn net.Conn) {
            log.Println("get connect from", conn.RemoteAddr())
            defer func(conn net.Conn) {
                _ = conn.Close()
                log.Printf("client %s closed.\n", conn.RemoteAddr())
            }(conn)

            // jsonrpc handler Client action
            jsonrpc.ServeConn(conn)
        }(conn)
    }
}
