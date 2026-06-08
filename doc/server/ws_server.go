package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var (
	//go:embed static/index.html
	indexHTML    embed.FS
	clients      = make(map[*websocket.Conn]bool)
	clientsMutex = sync.RWMutex{}
	upgrade      = websocket.Upgrader{
		CheckOrigin: func(res *http.Request) bool {
			return true
		},
	}
)

func main() {
	router := gin.Default()
	router.GET("/", getIndex)
	router.GET("/ws", initWebsocket)
	router.GET("/getOnline", getOnline)

	router.Run(":8080")
}

func initWebsocket(ctx *gin.Context) {
	conn, err := upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Printf("建立 Websocket 连接失败！错误信息：%v", err)
		return
	}

	clientsMutex.Lock()
	clients[conn] = true
	clientsMutex.Unlock()

	log.Printf("ID:%s 加入了聊天室！", conn.RemoteAddr())
	broadcast("有新客户端加入聊天室！")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("获取 ID:%s 发送的信息失败！错误信息：%v", conn.RemoteAddr(), err)
			break
		}
		broadcast(string(msg))
	}
}

func broadcast(msg string) {
	clientsMutex.RLock()
	for client := range clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Printf("ID:%s 发送消息失败！错误信息：%v", client.RemoteAddr(), err)
			delete(clients, client)
		}
	}
	clientsMutex.RUnlock()
}

func getIndex(ctx *gin.Context) {
	file, err := indexHTML.ReadFile("static/index.html")
	if err != nil {
		log.Printf("读取 index.html 文件失败！错误信息：%v", err)
		ctx.JSON(http.StatusInternalServerError, "Server error")
		return
	}
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", file)
}

func getOnline(ctx *gin.Context) {
	clientsMutex.RLock()
	i := len(clients)
	clientsMutex.RUnlock()
	ctx.JSON(http.StatusOK, gin.H{
		"code":   200,
		"online": i,
	})
}
