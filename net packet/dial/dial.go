package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	httprequest := "GET / HTTP/1.1\n" +
		"Host: golang.org\n\n" //сообщение для отправки
	conn, err := net.Dial("tcp", "golang.org:80") // диал для отправки запроса
	if err != nil {                               //если ошибка не пуста
		panic("sd")
	}
	defer conn.Close() // в конце закрываем отправки запросов

	if _, err = conn.Write([]byte(httprequest)); err != nil { //обязательно отправляем только срез байт
		fmt.Print(err)
		return
	}
	io.Copy(os.Stdout, conn)
	fmt.Print("done")

}
