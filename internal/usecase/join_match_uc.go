package usecase

import (
	"errors"
	domain "typo/core/internal/domain"
	presenter "typo/core/internal/presenter"
	repository "typo/core/internal/repository"
)

type JoinMatchPort interface {
	Join(player *domain.Player, code string) (*domain.Room, error)
}

type JoinMatchUseCase struct {
	repository repository.RoomRepository
}

func NewJoinMatchUseCase(repo repository.RoomRepository) *JoinMatchUseCase {
	return &JoinMatchUseCase{
		repository: repo,
	}
}

func (q JoinMatchUseCase) Join(player *domain.Player, code string) (*presenter.RoomResponse, error) {
	room, err := q.repository.RetrieveRoom(code)

	if err != nil {
		return nil, err
	}

	if !room.CanJoin() {
		return nil, errors.New("room_is_closed")
	}

	_, err = room.Join(player)

	if err != nil {
		return nil, err
	}

	err = q.repository.Update(room)

	if err != nil {
		return nil, errors.New("store_room_error")
	}

	return presenter.RoomResponseFromEntity(room), nil

}
