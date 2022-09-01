package main

import (
	"easy-rpc/codec"
	"easy-rpc/codec/geerpc"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:6900")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = conn.Close() }()
	_ = json.NewEncoder(conn).Encode(geerpc.DefaultOption)
	cc := codec.NewGobCodec(conn)

	for i := 0; i < 3; i++ {
		h := &codec.Header{
			ServiceMethod: "Foo.Sum",
			Seq:           uint64(i),
		}
		_ = cc.Write(h, fmt.Sprintf("geerpc req %d", h.Seq))
		_ = cc.ReadHeader(h)
		var reply string
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)
	}
}
