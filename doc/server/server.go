package main

//var clients = make(map[net.Conn]bool)
//
//func main() {
//	listener, err := net.Listen("tcp", ":8080")
//	if err != nil {
//		log.Printf("服务端启动失败！错误信息：%v", err)
//		return
//	}
//	defer listener.Close()
//	log.Println("服务端已启动！")
//
//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			log.Printf("获取客户端连接失败！错误信息：%v", err)
//			continue
//		}
//		clients[conn] = true
//		log.Printf("新客户端 %v 已连接！", conn.RemoteAddr())
//		go handleConn(conn)
//	}
//}
//
//func handleConn(conn net.Conn) {
//	defer func() {
//		delete(clients, conn)
//		conn.Close()
//		log.Printf("客户端 %v 已断开连接！", conn.RemoteAddr())
//	}()
//	reader := bufio.NewReader(conn)
//	for {
//		msg, err := reader.ReadString('\n')
//		if err != nil {
//			log.Printf("读取客户端 %v 消息失败！错误信息：%v", conn.RemoteAddr(), err)
//			break
//		}
//		msg = strings.TrimSpace(msg)
//		if msg == "" {
//			continue
//		}
//		log.Printf("获取客户端 %s 消息：%s", conn.RemoteAddr(), msg)
//		broadcast(msg, conn)
//	}
//}
//
//func broadcast(msg string, sender net.Conn) {
//	for client := range clients {
//		if client == sender {
//			continue
//		}
//		_, err := client.Write([]byte(msg + "\n"))
//		if err != nil {
//			log.Printf("客户端 %s 的消息写入失败！错误信息：%v", sender.RemoteAddr(), err)
//			client.Close()
//			delete(clients, client)
//		}
//	}
//}
