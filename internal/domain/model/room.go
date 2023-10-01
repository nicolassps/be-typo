package model

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Room struct {
	Id           string     `json:"id"`
	Status       RoomStatus `json:"status"`
	CreationTime time.Time  `json:"creation_time"`
	Rounds       []Round    `json:"rounds"`
	Players      []Player   `json:"players"`
}

func CreateRoom(r Room) *Room {
	id := uuid.NewV4()

	return &Room{
		Id:           id.String(),
		Status:       Created,
		CreationTime: time.Now(),
		Rounds:       []Round{},
		Players:      []Player{},
	}
}

func (r *Room) StartGame() (bool, error) {
	if r.Status != Created || len(r.Players) <= 1 {
		return false, errors.New("invalid_start_operation")
	}

	r.Status = Starting
	return true, nil
}

func (r *Room) Join(player *Player) (bool, error) {
	if r.Status != Created || len(r.Players) >= 8 {
		return false, errors.New("Room in progress, can't join it")
	}

	for _, p := range r.Players {
		if player.Session == p.Session {
			return false, errors.New("Player already joined in room")
		}
	}

	r.Players = append(r.Players, *player)
	if len(r.Players) == 8 {
		r.StartGame()
	}

	return true, nil
}
