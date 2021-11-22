package handler

import (
	"HOPE-backend/models"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func (ths *handler) reader(conn *websocket.Conn) error {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return err
		}

		var req models.NewChatRequest
		json.Unmarshal(p, &req)

		chat, err := ths.svc.NewChat(req)
		if err != nil {
			return err
		}

		msg, _ := json.Marshal(chat)

		if err := conn.WriteMessage(messageType, msg); err != nil {
			return err
		}

	}
}

func (ths *handler) ServeChatWS(c *gin.Context) {
	ws, err := ths.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
	}

	ths.reader(ws)
}
