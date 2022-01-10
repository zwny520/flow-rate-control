package main

import (
	"fmt"
	"net"
	"time"
)

func Connect() net.Conn {
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("err : ", err)
		time.Sleep(1 * time.Second)
		Connect()
	}

	if(!SendMessage(conn)){
		fmt.Println("err : ", err)
		time.Sleep(1 * time.Second)
		Connect()
	}

	return conn
}

func SendMessage(conn net.Conn) bool {
	for  {
		_, err := conn.Write([]byte("client send request!")) // 发送数据
		if err != nil {
			break
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			break
		}
		fmt.Println(time.Now().Format("15:04:05")+ " recv server respone:"+string(buf[:n]))
	}
	return false
}
// TCP 客户端
func main() {
	conn := Connect()
	defer conn.Close() // 关闭TCP连接
}
