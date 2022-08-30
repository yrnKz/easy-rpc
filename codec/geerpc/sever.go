package geerpc

import (
	"easy-rpc/codec"
	"encoding/json"
	"io"
	"log"
	"net"
	"reflect"
	"sync"
)

const MagicNumber = 0x3bef5c

type Option struct{
	MagicNumber int   
	CodecType codec.Type
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber ,
	CodecType: codec.GobType,
}


type Server struct{}

func NewServer() * Server{
	return &Server{}
}

var DefaultServer = NewServer()

func (server * Server) Accept (lis net.Listener){
	for {
		conn,err := lis.Accept()
		if err != nil{
			log.Println("rpc server: accept error:", err)
			return
		}
		go server.ServeConn(conn)
	}
}

func Accept(lis net.Listener) {
	DefaultServer.Accept(lis)
}

func (server *Server) ServeConn(conn io.ReadWriteCloser){
	defer func ()  {
		_  =  conn.Close()
	}()
	var opt Option
	if err := json.NewDecoder(conn).Decode(&opt);err!=nil{
		log.Println("rpc server: options error: ", err)
		return
	}
	if opt.MagicNumber != MagicNumber {
		log.Printf("rpc server: invalid magic number %x", opt.MagicNumber)
		return
	}
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil{
		log.Printf("rpc server: invalid codec type %s", opt.CodecType)
		return
	}
	server.serveCodec(f(conn))
}

func (server *Server) serveCodec(cc codec.Codec) {
	sending := new(sync.Mutex) 
	wg := new(sync.WaitGroup)

	for{
		req,err := server.readRequest(cc)
		if err != nil{
			if req  ==nil{
				break
			}
			req.Error
		}
		
	}
}

type request struct {
	h          *codec.Header // header of request
	argv, replyv reflect.Value // argv and replyv of request
}

func (server *Server) readRequest(cc codec.Codec) (*codec.Header, error) {
	return nil,nil
}