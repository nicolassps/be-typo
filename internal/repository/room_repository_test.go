package repository

import (
	"testing"
	domain "typo/core/internal/domain"
)

func TestRoomRepositoryStore(t *testing.T) {
	repo := NewRoomInMemoryRepository()

	room := domain.CreateRoom()
	repo.Store(room)

	if repo.memory[room.FriendlyCode()] == nil {
		t.Errorf("Repository cannot store a room")
	}
}

func TestRoomRepositoryStoreOnDuplicateKey(t *testing.T) {
	repo := NewRoomInMemoryRepository()

	room := domain.CreateRoom()
	repo.Store(room)

	if repo.memory[room.FriendlyCode()] == nil {
		t.Errorf("Repository cannot store a room")
	}

	err := repo.Store(room)

	if err.Error() != "duplicated_entry" {
		t.Errorf("Repository stored a duplicated entry")
	}
}

func TestRoomRepositoryUpdate(t *testing.T) {
	repo := NewRoomInMemoryRepository()

	room := domain.CreateRoom()
	repo.Store(room)

	if repo.memory[room.FriendlyCode()] == nil {
		t.Errorf("Repository cannot store a room")
	}

	player := domain.NewPlayer("PLAYER", "AVATAR")

	room.Join(player)
	repo.Update(room)

	rooms := repo.RetrieveCreatedRooms()

	if rooms[0].Id != room.Id {
		t.Errorf("Repository cannot retriave a created room")
	}

	if len(rooms[0].Players) != 1 {
		t.Errorf("Repository cannot update with new player at room")
	}
}
