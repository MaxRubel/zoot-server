package types

import (
	"github.com/gorilla/websocket"
)

type Room struct {
	ID           string `json:"id"`
	Participants int    `json:"participants"`
	Clients      []*Client
}

type Rooms []Room

type Client struct {
	ID     string
	WsConn *websocket.Conn
}

type JoinRequest struct {
	RoomId   string `json:"roomId"`
	ClientId string `json:"clientId"`
}

var Clients = make(map[string]*websocket.Conn)
