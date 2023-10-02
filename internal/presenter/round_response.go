package presenter

import (
	domain "typo/core/internal/domain"
)

type RoundResponse struct {
	Number int    `json:"number"`
	Letter string `json:"letter"`
}

func RoundResponseFromEntity(round domain.Round) *RoundResponse {
	return &RoundResponse{
		Number: round.Number,
		Letter: round.Letter,
	}
}
