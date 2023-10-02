package usecase

import (
	"testing"
	domain "typo/core/internal/domain"
	"typo/core/internal/repository"
)

func TestQuickMatchOnSuccess(t *testing.T) {
	repo := repository.NewRoomInMemoryRepository()
	uc := NewQuickMatchUseCase(repo)

	player1 := domain.NewPlayer("PLAYER_1", "AVATAR_1")

	room, _ := uc.Match(player1)

	if room == nil || room.Code == "" {
		t.Errorf("Cannot created room")
	}

	player2 := domain.NewPlayer("PLAYER_2", "AVATAR_2")

	room2, _ := uc.Match(player2)

	if room == nil || room.Code == "" {
		t.Errorf("Cannot find room to second player")
	}

	if room.Code != room2.Code {
		t.Errorf("The rooms finded are different")
	}

	if room2.TotalPlayers != 2 {
		t.Errorf("The two players cant join on the same room")
	}
}
