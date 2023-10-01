package model

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

const MAX_ROUNDS_PER_ROOM = 5

type Room struct {
	Id           string     `json:"id"`
	Status       RoomStatus `json:"status"`
	CreationTime time.Time  `json:"creation_time"`
	Rounds       []Round    `json:"rounds"`
	Players      []Player   `json:"players"`
}

func CreateRoom() *Room {
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
		return false, errors.New("room_already_in_progress")
	}

	for _, p := range r.Players {
		if player.Session == p.Session {
			return false, errors.New("player_already_connected")
		}
	}

	r.Players = append(r.Players, *player)
	if len(r.Players) == 8 {
		r.StartGame()
	}

	return true, nil
}

func (r *Room) NewRound() (*Round, error) {
	if !(r.Status == Starting || r.Status == Started) {
		return nil, errors.New("new_round_unsuported_operation")
	}

	if len(r.Rounds) >= MAX_ROUNDS_PER_ROOM {
		return nil, errors.New("new_round_unsuported_operation")
	}

	if r.Status == Starting {
		r.Status = Started
	}

	newRound := r.createRoundDistinctLetter()
	r.Rounds = append(r.Rounds, *newRound)

	return newRound, nil
}

func (r *Room) createRoundDistinctLetter() *Round {
	round := NewRound(len(r.Rounds) + 1)

	if r.existLetter(round.Letter) {
		return r.createRoundDistinctLetter()
	}

	return round
}

func (r *Room) existLetter(letter string) bool {
	for _, round := range r.Rounds {
		if round.Letter == letter {
			return true
		}
	}

	return false
}
