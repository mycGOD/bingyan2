package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	//建立socket，监听端口
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:1024")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}
		go handleConn(tcpConn)
	}
}
func handleConn(conn *net.TCPConn) {
	buffer := make([]byte, 2048)
	conn.Read(buffer)
	fmt.Println(string(buffer))
	fmt.Println("接收到消息")
	go sendToClient(buffer, 0)

}
func sendToClient(buffer []byte, no int) {
	fmt.Println("发送至client的" + string(buffer))
	//var clients []string
	//clients = []string{"127.0.0.1:1025", "127.0.0.1:1026", "127.0.0.1:1027"}
	var client string
	str := string(buffer)
	for j := 0; j < len(str); j++ {
		if str[j] == 0 {
			str = str[:j]
			break
		}
	}
	spl := strings.Split(str, "-")
	if spl[1] != "Users_commend" && spl[1] != "Users_announce" && no != 3 {
		client = "127.0.0.1:1027"
	} else if spl[1] != "Users_text" && spl[1] != "Users_friends" && no != 1 {
		client = "127.0.0.1:1025"
	} else {
		client = "127.0.0.1:1026"
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp4", client)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(-1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(-1)
	}
	_, err = conn.Write(buffer)
	if err != nil {
		fmt.Println("Error to send message because of ", err.Error())
		os.Exit(-1)
	}
	fmt.Println("函数已返回")
}

//处理连接

func Log(v ...interface{}) {
	log.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
