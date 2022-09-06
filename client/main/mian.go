package main

import (
	"easy-rpc/client"
	"fmt"
	"log"
	"sync"
)

func main() {
	log.SetFlags(0)
	var wg sync.WaitGroup
	clients, _ := client.Dial("tcp", "127.0.0.1:6900")
	defer func() {
		_ = clients.Close()
	}()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := fmt.Sprintf("geerpc req %d", i)
			var reply string
			if err := clients.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Println("reply", reply)
		}(i)
	}
	wg.Wait()
}
