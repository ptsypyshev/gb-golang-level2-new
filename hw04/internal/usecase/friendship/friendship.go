package friendship

import (
	"context"

	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/models"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/repo"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/usecase"
)

var _ usecase.Friendship = (*friends)(nil)

// friends is an implementation of Friendship interface.
type friends struct {
	fRepo repo.FriendshipRepo
	uRepo repo.UserRepo
}

// New is a constructor for friends.
func New(u repo.UserRepo, f repo.FriendshipRepo) *friends {
	return &friends{
		uRepo: u,
		fRepo: f,
	}
}

// Create implements repo.FriendshipRepo.
func (f *friends) CreateFriendship(ctx context.Context, fr models.Friendship) error {
	_, err := f.uRepo.Read(ctx, fr.SourceID)
	if err != nil {
		return err
	}

	_, err = f.uRepo.Read(ctx, fr.TargetID)
	if err != nil {
		return err
	}

	return f.fRepo.Create(ctx, fr)
}

// Read implements repo.FriendshipRepo.
func (f *friends) ReadFriends(ctx context.Context, id int) ([]models.User, error) {
	friends, err := f.fRepo.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	return friends, nil
}
