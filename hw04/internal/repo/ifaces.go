package repo

import (
	"context"

	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/models"
)

// UserRepo is a repository for User model (provides CRUD operations).
type UserRepo interface {
	Create(ctx context.Context, u models.User) (int, error)
	Read(ctx context.Context, id int) (models.User, error)
	Update(ctx context.Context, u models.User) error
	Delete(ctx context.Context, id int) error
}

// FriendshipRepo is a repository for Friendship model (provides some of CRUD operations).
type FriendshipRepo interface {
	Create(ctx context.Context, u models.Friendship) error
	Read(ctx context.Context, id int) ([]models.User, error)
}
