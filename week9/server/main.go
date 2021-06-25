package main

import (
	"bufio"
	"fmt"
	"net"

	"protocol"
)

// 服务端 fix length
func server_tcp_fix_length(conn net.Conn) {
	fmt.Println("server, fix length")
	const (
		BYTE_LENGTH = 1024
	)

	for {
		var buf = make([]byte, BYTE_LENGTH)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("client data :", string(buf))
	}
}

// 服务端 delimiter based
func server_tcp_delimiter(conn net.Conn) {
	fmt.Println("server, delimiter based")

	reader := bufio.NewReader(conn)
	for {
		slice, err := reader.ReadSlice('\n')
		if err != nil {
			continue
		}
		fmt.Printf("%s", slice)
	}
}

// 服务端 length field based frame decoder
func server_tcp_frame_decoder(conn net.Conn) {
	fmt.Println("server, length field based frame decoder")
	var buf = make([]byte, 0)
	var readerChannel = make(chan []byte, 16)
	go func() {
		select {
		case data := <-readerChannel:
			fmt.Println("channel=", string(data))
		}
	}()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}

		protocol.Unpack(append(buf, buffer[:n]...), readerChannel)
	}
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := lis.Accept()
		defer conn.Close()

		if err != nil {
			fmt.Println(err)
			return
		}
		go server_tcp_delimiter(conn)
	}

}
