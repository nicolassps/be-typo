package model

import uuid "github.com/satori/go.uuid"

type Player struct {
	Session string  `json:"session_id"`
	Name    string  `json:"name"`
	Avatar  string  `json:"avatar"`
	Rounds  []Round `json:"rounds"`
}

func NewPlayer(name string, avatar string) *Player {
	id := uuid.NewV4().String()

	return &Player{
		Session: id,
		Name:    name,
		Avatar:  avatar,
		Rounds:  []Round{},
	}
}
