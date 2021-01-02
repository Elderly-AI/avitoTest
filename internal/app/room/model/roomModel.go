//go:generate easyjson -all roomModel.go
package roomModel

// easyjson:json
type Room struct {
	RoomID      int    `json:"room_id" db:"room_id"`
	Created     string `json:"create" db:"created"`
	Cost        int    `json:"cost" db:"cost"`
	Description string `json:"description" db:"description"`
}

// easyjson:json
type RoomAdd struct {
	Cost        int    `json:"cost"`
	Description string `json:"description"`
}

type RoomOrder struct {
	Sort  string
	Order string
}

// easyjson:json
type RoomID struct {
	RoomID int `json:"room_id"`
}
