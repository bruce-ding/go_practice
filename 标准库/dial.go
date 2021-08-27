package main

import (
	"fmt"
	"io"
	"net"
	"time"
)


// Dialer is the interface for creating a specialized net.Conn.
//type Dialer interface {
//	Dial(network, addr string) (c net.Conn, err error)
//}

func dial() {
	d := &net.Dialer{Timeout: time.Second}
	//dialer := Dialer(d)
	//if srv.Network == "tcp+tls" {
	//	dialer = &tls.Dialer{NetDialer: d, Config: srv.TLS}
	//}

	var err error
	var conn net.Conn

	conn, err = d.Dial("tcp", "localhost:3000")
	if err != nil {
		fmt.Println(err)
		return
	}

	if x, ok := conn.(*net.TCPConn); ok {
		err = x.SetKeepAlive(true)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	bytes, err := io.ReadAll(conn)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bytes))

	//fmt.Println(conn.Read())
	//
	//r := bufio.NewReader(conn)
	//w := bufio.NewWriter(conn)
	//
	//line, err := readString(r)
	//if err != nil {
	//	conn.Close()
	//	return nil, err
	//}
}

func main() {
	dial()
}
