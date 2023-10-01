package model

import (
	random "typo/core/internal/pkg/random"
)

type Round struct {
	Number int    `json:"number"`
	Letter string `json:"letter"`
	Words  []Word `json:"words"`
}

func NewRound(number int) *Round {
	letter := random.RandomLetter()

	return &Round{
		Number: number,
		Letter: letter,
		Words:  []Word{},
	}
}
