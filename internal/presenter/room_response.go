package presenter

import (
	domain "typo/core/internal/domain"
)

type RoomResponse struct {
	Code         string          `json:"code"`
	TotalPlayers int             `json:"total_players"`
	Status       string          `json:"status"`
	Rounds       []RoundResponse `json:"rounds"`
}

func RoomResponseFromEntity(room *domain.Room) *RoomResponse {
	rounds := []RoundResponse{}

	for _, round := range room.Rounds {
		rounds = append(rounds, *RoundResponseFromEntity(round))
	}

	return &RoomResponse{
		Code:         room.FriendlyCode(),
		TotalPlayers: len(room.Players),
		Status:       string(room.Status),
		Rounds:       rounds,
	}
}
