package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//建立socket，监听端口
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:1026")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}
		handleConn(tcpConn)
	}
}

var db *sql.DB

func handleConn(conn *net.TCPConn) {
	buffer := make([]byte, 2048)
	len, _ := conn.Read(buffer)
	fmt.Println(string(buffer))
	fmt.Println("接收到消息")
	sendToNext(buffer[0:len])
}

func sendToNext(buffer []byte) {
	fmt.Println("发送到达client的" + string(buffer))
	str := string(buffer)
	str = strings.Replace(str, " ", "", -1)
	for j := 0; j < len(str); j++ {
		if str[j] == 0 {
			str = str[:j]
			break
		}
	}
	strs := strings.Split(str, "-")

	db, err := sql.Open("mysql", "root:qwertyuiop81@tcp(127.0.0.1:3306)/BINGYAN_2?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	ques2 := "select " + strs[0] + " from " + strs[1] + ";"
	fmt.Println(ques2)

	ques := "select author from Users_text;"
	fmt.Println(ques)
	fmt.Println(len(ques2))
	fmt.Println(len(ques))
	fmt.Println(string(ques) == ques2)
	rows, err := db.Query(ques2)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	var get []string
	for rows.Next() {
		var p string
		rows.Scan(&p)
		get = append(get, p)
	}
	gotten := strings.Join(get, "-")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8078")
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
	_, err = conn.Write([]byte(gotten))
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
