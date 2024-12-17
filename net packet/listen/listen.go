package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go test(wg)
	mess := "Hee"
	listen, err := net.Listen("tcp", ":4545")
	if err != nil {
		panic("err")
	}
	defer listen.Close()
	fmt.Print("serv listen")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Print("err")
			return
		}
		conn.Write([]byte(mess))
		conn.Close()
	}
}
func test(wt *sync.WaitGroup) {
	defer wt.Done()
	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	wt.Wait()

	io.Copy(os.Stdout, conn)
	fmt.Println("\nDone")
}
