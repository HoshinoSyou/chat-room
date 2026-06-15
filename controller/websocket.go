package controller

import (
	"chat-room/service"
	"chat-room/util/response"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	clients      = make(map[uint]*websocket.Conn)
	clientsMutex = sync.RWMutex{}
	upgrade      = websocket.Upgrader{
		CheckOrigin: func(req *http.Request) bool {
			return true
		},
	}
)

func InitWebsocket(ctx *gin.Context) {
	uidRaw, exists := ctx.Get("uid")
	if !exists {
		log.Printf("获取用户登录状态失败！")
		response.Error(ctx, "获取用户登录状态失败！", errors.New("获取用户登录状态失败！未获取到用户信息！"))
		return
	}
	uid := uidRaw.(uint)
	//username := ctx.GetString("username")
	log.Printf("用户 %d 尝试建立 websocket 连接", uid)

	// 升级 ws 协议
	conn, err := upgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Printf("升级 Websocket 连接失败！错误信息：%v", err)
		return
	}

	defer conn.Close()

	clientsMutex.Lock()
	oldConn, exists := clients[uid]
	if exists {
		oldConn.Close()
	}
	clients[uid] = conn
	clientsMutex.Unlock()

	defer func() {
		clientsMutex.Lock()
		delete(clients, uid)
		clientsMutex.Unlock()
		log.Printf("用户 %d 断开 WebSocket 连接！", uid)
	}()

	conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				err := conn.WriteMessage(websocket.PingMessage, nil)
				if err != nil {
					log.Printf("发送 Ping 包失败！错误信息：%v", err)
					return
				}
			}
		}
	}()

	for {
		//msg := models.Message{
		//	Model:      gorm.Model{},
		//	FromUserId: uid,
		//}
		var msgJSON struct {
			TargetUsername string `json:"targetUsername"`
			Content        string `json:"content"`
		}
		err := conn.ReadJSON(&msgJSON)
		if err != nil {
			log.Printf("读取用户 %d 发送的消息内容失败！错误信息：%v", uid, err)
			break
		}
		res, msg, err := service.CreateMessage(uid, msgJSON.TargetUsername, msgJSON.Content)
		if !res {
			log.Printf("用户 %d 发送的消息保存至 mysql 数据库失败！错误信息：%v", uid, err)
			conn.WriteJSON(gin.H{"error": "发送消息失败！"})
			continue
		}

		clientsMutex.RLock()
		targetConn, e := clients[msg.ToUserId]
		clientsMutex.RUnlock()
		if e {
			err := targetConn.WriteJSON(gin.H{
				"type":         "private_message",
				"message_id":   msg.ID,
				"time":         msg.CreatedAt,
				"from_user_id": uid,
				"content":      msg.Content,
			})
			if err != nil {
				log.Printf("将用户 %d 的消息推送至用户 %d 失败！错误信息：%v", uid, msg.ToUserId, err)
			}
		} else {
			log.Printf("将用户 %d 发送给用户 %d 的消息保存成功！", uid, msg.ToUserId)
		}
		conn.WriteJSON(gin.H{
			"status":     "ok",
			"message_id": msg.ID,
		})
	}
}
