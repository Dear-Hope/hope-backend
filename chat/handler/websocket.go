package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func (ths *handler) upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := ths.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	return ws, nil
}

func (ths *handler) ServeChatWS(c *gin.Context) {
	conn, err := ths.upgrade(c.Writer, c.Request)
	if err != nil {
		log.Println(err)
	}

	client := &Client{
		Conn: conn,
		Pool: ths.pool,
		svc:  ths.svc,
	}

	ths.pool.Register <- client
	client.Read()
}
