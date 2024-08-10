package pserver

import (
	"bufio"
	"fmt"
	"net"
)

func Serv() error {
	fmt.Print("starting server\n\n")
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		return err
	}

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')

	fmt.Print(status)
	return err
}
