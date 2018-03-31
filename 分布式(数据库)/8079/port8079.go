package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

/*func findFromDatabase(what string, which string, destory int) string {

}*/
func main() {
	//建立socket，监听端口
	netListen, err := net.Listen("tcp", "localhost:8079")
	CheckError(err)
	defer netListen.Close()
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:1024")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(-1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(-1)
	}
	conn.Write([]byte("person-Users_commend"))
	fmt.Println("8079已经寄出去")
	//--------------------------------------------------
	conn.Close()
	listenSock, err := net.Listen("tcp", "localhost:8078")
	CheckError(err)
	defer listenSock.Close()

	for {
		new_conn, err := listenSock.Accept()
		if err != nil {
			continue
		}

		recvConnMsg(new_conn)
		return
	}

}
func recvConnMsg(conn net.Conn) {
	//  var buf [50]byte
	buf := make([]byte, 50)

	defer conn.Close()
	var all string
	for {
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println(all)
			fmt.Println("conn closed")
			return
		}

		//fmt.Println("recv msg:", buf[0:n])
		all += string(buf[0:n])
		//fmt.Println("recv msg:", string(buf[0:n]))
	}
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
