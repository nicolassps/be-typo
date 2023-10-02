package usecase

import (
	"errors"
	domain "typo/core/internal/domain"
	presenter "typo/core/internal/presenter"
	repository "typo/core/internal/repository"
)

type QuickMatchPort interface {
	Match(player *domain.Player) (*domain.Room, error)
}

type QuickMatchUseCase struct {
	repository repository.RoomRepository
}

func NewQuickMatchUseCase(repo repository.RoomRepository) *QuickMatchUseCase {
	return &QuickMatchUseCase{
		repository: repo,
	}
}

func (q QuickMatchUseCase) Match(player *domain.Player) (*presenter.RoomResponse, error) {
	rooms := q.repository.RetrieveCreatedRooms()

	for _, room := range rooms {
		if room.CanJoin() {
			_, err := room.Join(player)

			if err != nil {
				return nil, err
			}

			err = q.repository.Update(&room)

			if err != nil {
				return nil, errors.New("store_room_error")
			}

			return presenter.RoomResponseFromEntity(&room), nil
		}
	}

	newRoom := domain.CreateRoom()
	_, err := newRoom.Join(player)

	if err != nil {
		return nil, err
	}

	err = q.repository.Store(newRoom)

	if err != nil {
		return nil, errors.New("store_room_error")
	}

	return presenter.RoomResponseFromEntity(newRoom), nil
}
