package handler

import (
	"HOPE-backend/v1/models"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
	svc  models.ChatService
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var response models.Response

		var req models.NewChatRequest
		json.Unmarshal(p, &req)

		chat, err := c.svc.NewChat(req)
		if err != nil {
			log.Println(err)
			response.Error = err.Error()
		}

		response.Result = chat

		c.Pool.Broadcast <- response
		fmt.Printf("Message Received: %+v\n", response)
	}
}
