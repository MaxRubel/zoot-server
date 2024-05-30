package ws

import (
	"fmt"
	"log"

	"github.com/MaxRubel/zoot-server/data"
	"github.com/gorilla/websocket"
)

func BroadcastMessage(message string, roomId string) {
	var num int
	for i, room := range data.Rooms {
		if room.ID == roomId {
			fmt.Println("broadcasting...")
			for j, client := range room.Clients {

				if client.WsConn == nil {
					num--
					log.Printf("Client %s has a nil WebSocket connection\n", client.ID)
					continue
				}
				err := client.WsConn.WriteMessage(websocket.TextMessage, []byte(message))
				num++
				if err != nil {
					log.Printf("Error writing message to client %s: %v\n", client.ID, err)
					// Optionally, you might want to remove the client from the room's list
					data.Rooms[i].Clients = append(room.Clients[:j], room.Clients[j+1:]...)
				}
			}
			break
		}
	}
	fmt.Printf("Broadcasted to %v clients", num)
}

// func removeClient(clients []*types.Client, clientId string) []*types.Client {
// 	var newClients []*types.Client
// 	for _, client := range clients {
// 		if client.ID != clientId {
// 			newClients = append(newClients, client)
// 		}
// 	}
// 	return newClients
// }
