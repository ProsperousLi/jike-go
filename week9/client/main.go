package main

import (
	//"bufio"
	"fmt"
	"net"
	//"os"

	"protocol"
)

// 客户端 fix length
func client_tcp_fix_length(conn net.Conn) {
	fmt.Println("client, fix length")
	sendByte := make([]byte, 1024)
	sendMsg := "{\"test01\":1,\"test02\",2}"
	for i := 0; i < 1000; i++ {
		tempByte := []byte(sendMsg)
		for j := 0; j < len(tempByte) && j < 1024; j++ {
			sendByte[j] = tempByte[j]
		}
		_, err := conn.Write(sendByte)
		if err != nil {
			fmt.Println(err, ",err index=", i)
			return
		}
		fmt.Println("send over once")
	}
}

// 客户端 delimiter based
func client_tcp_delimiter(conn net.Conn) {
	fmt.Println("client, delimiter based")
	var sendMsgs string
	sendMsg := "{\"test01\":1,\"test02\",2}\n"
	for i := 0; i < 1000; i++ {
		sendMsgs += sendMsg
		_, err := conn.Write([]byte(sendMsgs))
		if err != nil {
			fmt.Println(err, ",err index=", i)
			return
		}
		fmt.Println("send over once")
	}
}

// 客户端 length field based frame decoder
func client_tcp_frame_decoder(conn net.Conn) {
	fmt.Println("client, length field based frame decoder")
	for i := 0; i < 1000; i++ {
		sendMsg := "{\"test01\":1,\"test02\",2}"
		_, err := conn.Write(protocol.Packet([]byte(sendMsg)))
		if err != nil {
			fmt.Println(err, ",err index=", i)
			return
		}
		fmt.Println("send over once")
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//client_tcp_delimiter(conn)
	client_tcp_delimiter(conn)
}
