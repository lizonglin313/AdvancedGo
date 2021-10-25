package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	if os.Args[1] == "1" {
		listener, err := net.Listen("tcp", "localhost:8123")
		if err != nil {
			log.Fatal(err)
		}
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err) // e.g., connection aborted
				continue
			}
			//handleConnSend(conn)		// 同一时间只能开一个端口接收
			go handleConnSend(conn)		 // 可以同时开多个

		}
	} else {
		conn, err := net.Dial("tcp", "localhost:8123")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		go mustCopy(os.Stdout, conn)
		mustCopy(conn, os.Stdin)
	}

}
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}


func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

func handleConnSend(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func handleConnRecive(conn net.Conn) {
	defer conn.Close()
	br := bufio.NewReader(conn)
	for {
		data, err := br.ReadString('\n')
		if err != nil {
			log.Print(err)
			continue
		}
		//if err == io.EOF{
		//	break
		//}
		fmt.Printf("%s", data)
		fmt.Fprintf(conn, "OK\n")
	}

}
