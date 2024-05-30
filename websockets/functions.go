package ws

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/MaxRubel/zoot-server/data"
	"github.com/gorilla/websocket"
)

type Message struct {
	Content string `json:"roomId"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var Clients = make(map[string]*websocket.Conn)

func addWsToClient(cId string, rId string, conn *websocket.Conn) {
	fmt.Println("Searching for client:", cId, "in room:", rId)

	for i := range data.Rooms {
		if data.Rooms[i].ID == rId {
			fmt.Println("Found room:", rId)

			for x := range data.Rooms[i].Clients {
				client := data.Rooms[i].Clients[x]

				if client.ID == cId {
					client.WsConn = conn
					fmt.Println("Added websocket to client:", client.ID)
					return
				}
			}

			fmt.Println("Client not found in room:", rId)
			return
		}
	}
}

func checkHowManyClients(roomId string) {
	for i := range data.Rooms {
		if data.Rooms[i].ID == roomId {
			fmt.Printf("there are %v clients in this room", len(data.Rooms[i].Clients))
		}
	}
}

func WsHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return
	}

	for {
		_, incoming, _ := conn.ReadMessage()

		split := strings.Split(string(incoming), "&")

		clientId := split[2]
		roomId := split[1]
		messageType := split[0]

		if messageType == "0" {
			BroadcastMessage("testing server!", roomId)
			if err != nil {
				log.Println("Error writing message:", err)
				break
			}
		}
		if messageType == "1" {
			addWsToClient(clientId, roomId, conn)
			checkHowManyClients(roomId)

			if err != nil {
				log.Println(err)
				break
			}
		}

	}
	defer conn.Close()
}
