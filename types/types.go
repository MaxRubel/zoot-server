package types

type Room struct {
	ID           string `json:"id"`
	Participants int    `json:"participants"`
}

type Rooms []Room
