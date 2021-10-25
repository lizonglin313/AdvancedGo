package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	message  = make(chan string)
)

// broadcaster
// @Desc: 	监听全局的 message entering leaving 事件
// @Notice:
func broadcaster() {
	clients := make(map[client]bool)

	for {
		select {
		case msg := <-message: // 如果有消息
			for cli := range clients {
				cli <- msg // 把消息给 client
			}
		case cli := <-entering: // 如果有新来的客户端
			clients[cli] = true
		case cli := <-leaving: // 如果有客户端离开
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn)  {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	message <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		message <- who + ": " + input.Text()
	}

	leaving <- ch
	message <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string)  {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
