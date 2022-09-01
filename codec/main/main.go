package main

import (
	//"easy-rpc/codec"
	"easy-rpc/codec/geerpc"
	//"encoding/json"
	//"fmt"
	"log"
	"net"
	//"time"
)

func StartServer(addr chan string) {
	l, err := net.Listen("tcp", "127.0.0.1:6900")
	if err != nil {
		log.Fatal("network error:", err)
	}
	//addr <- l.Addr().String()
	log.Println("start rpc server on", l.Addr())
	geerpc.Accept(l)
}

func init() {
	log.SetFlags(log.Lshortfile | log.Llongfile)
}
func main() {
	addr := make(chan string)
	StartServer(addr)

	//conn, _ := net.Dial("tcp", <-addr)
	//defer func() { _ = conn.Close() }()
	//
	//time.Sleep(time.Second)
	//
	//_ = json.NewEncoder(conn).Encode(geerpc.DefaultOption)
	//cc := codec.NewGobCodec(conn)
	//
	//for i := 0; i < 3; i++ {
	//	h := &codec.Header{
	//		ServiceMethod: "Foo.Sum",
	//		Seq:           uint64(i),
	//	}
	//	_ = cc.Write(h, fmt.Sprintf("geerpc req %d", h.Seq))
	//	_ = cc.ReadHeader(h)
	//	var reply string
	//	_ = cc.ReadBody(&reply)
	//	log.Println("reply:", reply)
	//}
}
