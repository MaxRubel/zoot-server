package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MaxRubel/zoot-server/types"
)

var rooms types.Rooms

func writeCORS(w http.ResponseWriter, r *http.Request) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
        return false
    } else { 
		return true
	}
}

func makeNewRoom(w http.ResponseWriter, r *http.Request) {

	ok:= writeCORS(w, r)

	if ok {
		var room types.Room

		err := json.NewDecoder(r.Body).Decode(&room)

		if err != nil {
			fmt.Println("Error decoding JSON Request body:", err)
			return
		}

		rooms = append(rooms, room)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(room)

		fmt.Println("created new room:", room.ID)
	}

}

func main() {
	http.HandleFunc("/newRoom", makeNewRoom)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
