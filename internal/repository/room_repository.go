package repository

import (
	"errors"
	domain "typo/core/internal/domain"
)

type RoomRepository interface {
	RetrieveCreatedRooms() []domain.Room
	RetrieveRoom(code string) (*domain.Room, error)
	Store(room *domain.Room) error
	Update(room *domain.Room) error
}

type RoomInMemoryRepository struct {
	memory map[string]*domain.Room
}

func NewRoomInMemoryRepository() *RoomInMemoryRepository {
	return &RoomInMemoryRepository{
		memory: make(map[string]*domain.Room),
	}
}

func (r RoomInMemoryRepository) RetrieveCreatedRooms() []domain.Room {
	result := []domain.Room{}

	for _, room := range r.memory {
		if room.Status == domain.Created {
			result = append(result, *room)
		}
	}

	return result
}

func (r RoomInMemoryRepository) RetrieveRoom(code string) (*domain.Room, error) {
	for _, room := range r.memory {
		if room.FriendlyCode() == code {
			return room, nil
		}
	}

	return nil, errors.New("room_not_found")
}

func (r RoomInMemoryRepository) Store(room *domain.Room) error {
	if r.memory[room.FriendlyCode()] != nil {
		return errors.New("duplicated_entry")
	}

	r.memory[room.FriendlyCode()] = room
	return nil
}

func (r RoomInMemoryRepository) Update(room *domain.Room) error {
	if r.memory[room.FriendlyCode()] == nil {
		return errors.New("entry_not_found")
	}

	r.memory[room.FriendlyCode()] = room
	return nil
}
