package usecase

import (
	"fmt"
	"testing"
	domain "typo/core/internal/domain"
	repository "typo/core/internal/repository"
)

func TestJoinMatchOnSuccess(t *testing.T) {
	repo := repository.NewRoomInMemoryRepository()
	newQuickMatchUseCase := NewQuickMatchUseCase(repo)
	uc := NewJoinMatchUseCase(repo)

	player1 := domain.NewPlayer("PLAYER_1", "AVATAR_1")

	room, _ := newQuickMatchUseCase.Match(player1)
	if room == nil || room.Code == "" {

		t.Errorf("Cannot created room")
	}

	player2 := domain.NewPlayer("PLAYER_2", "AVATAR_2")

	room2, _ := uc.Join(player2, room.Code)

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

func TestJoinMatchOnNotFound(t *testing.T) {
	repo := repository.NewRoomInMemoryRepository()
	uc := NewJoinMatchUseCase(repo)

	player := domain.NewPlayer("PLAYER_2", "AVATAR_2")

	room, err := uc.Join(player, "ANY_CODE")

	if room != nil || err == nil {
		t.Errorf("Room created, not expected")
	}

	if err.Error() != "room_not_found" {
		msg := fmt.Sprintf("Not expected error %s", err.Error())
		t.Errorf(msg)
	}
}
