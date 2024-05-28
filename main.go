package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MaxRubel/zoot-server/data"
	"github.com/MaxRubel/zoot-server/types"
	ws "github.com/MaxRubel/zoot-server/websockets"
)

func writeCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
}


func newClient(w http.ResponseWriter, r *http.Request) {
    
	writeCORS(w)
    if r.Method != http.MethodPost {
        return
    }

    var cRequest types.JoinRequest
    err := json.NewDecoder(r.Body).Decode(&cRequest)

    if err != nil {
        fmt.Println("Error decoding JSON client request:", err)
        return
    }

    var client types.Client
    client.ID = cRequest.ClientId
	
	for i := range data.Rooms {
		if data.Rooms[i].ID == cRequest.RoomId {
			data.Rooms[i].Clients = append(data.Rooms[i].Clients, &client)
			fmt.Println("Adding client to room:", client.ID)
			break
		}
	}
	
}

func serveRooms(w http.ResponseWriter, r *http.Request){
	writeCORS(w)

	if r.Method != http.MethodGet {
		return
	}

	var roomIds []string
	for _, r := range data.Rooms {
		roomIds = append(roomIds, r.ID)
	}

	roomsJ, err := json.Marshal(roomIds)

	if err != nil {
		fmt.Println("Error encoding room id array into json")
		return
	}

	fmt.Println("Serving rooms...")
	w.Write(roomsJ)
}

func makeNewRoom(w http.ResponseWriter, r *http.Request) {

	writeCORS(w)

	if r.Method != http.MethodPost {
		return
	}

		var room types.Room

		err := json.NewDecoder(r.Body).Decode(&room)

		if err != nil {
			fmt.Println("Error decoding JSON Request body:", err)
			return
		}

		data.Rooms = append(data.Rooms, room)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(room)

		fmt.Println("created new room:", room.ID)
}

func main() {

	http.HandleFunc("/newRoom", makeNewRoom)
	http.HandleFunc("/rooms", serveRooms)
	http.HandleFunc("/joinRoom", newClient)
	http.HandleFunc("/ws", ws.WsHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
