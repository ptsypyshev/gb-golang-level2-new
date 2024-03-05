package friends

import (
	"context"

	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/models"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/repo"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/storage"
)

var _ repo.FriendshipRepo = (*friends)(nil)

// friends is an implementation of FriendshipRepo interface.
type friends struct {
	storage storage.FriendshipStor
}

// New is a constructor for friends.
func New(s storage.FriendshipStor) *friends {
	return &friends{
		storage: s,
	}
}

// Create implements repo.FriendshipRepo.
func (f *friends) Create(ctx context.Context, fr models.Friendship) error {
	return f.storage.CreateFriendship(ctx, fr)
}

// Read implements repo.FriendshipRepo.
func (f *friends) Read(ctx context.Context, id int) ([]models.User, error) {
	return f.storage.ReadFriends(ctx, id)
}
