package model

import (
	"testing"
)

func TestFunctionalRoomAtFirstRound(t *testing.T) {
	room := CreateRoom()

	if room.Status != Created {
		t.Errorf("Room is not on created status")
	}

	player1 := NewPlayer("PLAYER_1", "avatar")
	player2 := NewPlayer("PLAYER_2", "avatar")

	room.Join(player1)

	started, err := room.StartGame()

	if len(room.Players) != 1 {
		t.Errorf("Room not join the player 1")
	}

	if started || err == nil || room.Status != Created {
		t.Errorf("Room started with 1 player")
	}

	room.Join(player2)

	if len(room.Players) != 2 {
		t.Errorf("Room not join the player 2")
	}

	started, err = room.StartGame()

	if !started || err != nil || room.Status != Starting {
		t.Errorf("Room not started with 2 players")
	}

	round, err := room.NewRound()

	if round == nil || err != nil || len(room.Rounds) != 1 {
		t.Errorf("Room unable to create a round")
	}

	round, err = room.NewRound()

	if round == nil || err != nil || len(room.Rounds) != 2 {
		t.Errorf("Room unable to create a second round")
	}

	room.NewRound()
	room.NewRound()
	round, err = room.NewRound()

	if round == nil || err != nil || len(room.Rounds) != 5 {
		t.Errorf("Room unable to create the last round")
	}

	round, err = room.NewRound()

	if round != nil || err == nil || len(room.Rounds) > 5 {
		t.Errorf("Room create more than max rounds granted")
	}
}
