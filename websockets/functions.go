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

func WsHandler(w http.ResponseWriter, r *http.Request) {

    conn, err := upgrader.Upgrade(w, r, nil)

    if err != nil {
        return
    }

	for {
		_, incoming, err := conn.ReadMessage()

		split := strings.Split(string(incoming), "&")

		clientId :=  split[1]
		roomId := split[0]

		fmt.Println("roomId:", roomId)

		addWsToClient(clientId, roomId, conn)

		if err != nil {
			log.Println(err)
			break
		}
	}

    defer conn.Close()

}

// func HandleConnections(w http.ResponseWriter, r *http.Request) {
// 	// Upgrade initial GET request to a WebSocket
// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer ws.Close()

// 	// Register new client
// 	clients[ws] = true

// 	for {
// 		var msg Message
// 		// Read in a new message as JSON and map it to a Message object
// 		err := ws.ReadJSON(&msg)
// 		if err != nil {
// 			log.Printf("error: %v", err)
// 			delete(clients, ws)
// 			break
// 		}
// 		// Send the received message to the broadcast channel
// 		broadcast <- msg
// 	}
// }

// func HandleMessages() {
// 	for {
// 		// Grab the next message from the broadcast channel
// 		msg := <-broadcast
// 		// Send it out to every client that is currently connected
// 		for client := range clients {
// 			err := client.WriteJSON(msg)
// 			if err != nil {
// 				log.Printf("error: %v", err)
// 				client.Close()
// 				delete(clients, client)
// 			}
// 		}
// 	}
// }

// // // Initialize the map
// // var roomWebsockets = make(map[string][]*websocket.Conn)

// // // Function to add a WebSocket connection to a room
// // func addWebSocketToRoom(roomID string, ws *websocket.Conn) {
// //     roomWebsockets[roomID] = append(roomWebsockets[roomID], ws)
// // }

// // // Function to remove a WebSocket connection from a room
// // func removeWebSocketFromRoom(roomID string, ws *websocket.Conn) {
// //     var updatedWS []*websocket.Conn
// //     for _, conn := range roomWebsockets[roomID] {
// //         if conn != ws {
// //             updatedWS = append(updatedWS, conn)
// //         }
// //     }
// //     roomWebsockets[roomID] = updatedWS
// // }

// // func iterateWebSocketConnections(roomID string) {
	
// //     for _, ws := range roomWebsockets[roomID] {

// //     }
// // }