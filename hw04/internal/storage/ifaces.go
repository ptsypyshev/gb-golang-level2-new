package storage

import (
	"context"

	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/models"
)

// UserStor is a storage for User model (provides CRUD operations).
type UserStor interface {
	CreateUser(ctx context.Context, u models.User) (int, error)
	ReadUser(ctx context.Context, id int) (models.User, error)
	UpdateUser(ctx context.Context, u models.User) error
	DeleteUser(ctx context.Context, id int) error
}

// FriendshipStor is a storage for Friendship model (provides some of CRUD operations).
type FriendshipStor interface {
	CreateFriendship(ctx context.Context, u models.Friendship) error
	ReadFriends(ctx context.Context, id int) ([]models.User, error)
}
