package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Printf("连接服务端失败！错误信息：%v", err)
		return
	}
	defer conn.Close()
	log.Println("已连接至聊天室！输入消息以回车键结束，输入 exit 可退出聊天室。")

	go readMsg(conn)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := strings.TrimSpace(scanner.Text())
		if msg == "exit" {
			log.Printf("客户端 %s 已退出聊天室！", conn.RemoteAddr())
			return
		}
		if msg == "" {
			continue
		}
		_, err := conn.Write([]byte(msg + "\n"))
		if err != nil {
			log.Printf("发送消息失败！错误信息：%v", err)
			return
		}
	}
}

func readMsg(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("\n客户端 %s 接收消息失败！已断开连接！错误信息：%v", conn.RemoteAddr(), err)
			os.Exit(0)
		}
		log.Printf("\n收到广播消息：%s", msg)
		log.Printf("> ")
	}
}
