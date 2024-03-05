package usecase

import (
	"context"

	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/models"
)

// Friendship is a usecase for Friendship model (provides operations with friends).
type Friendship interface {
	CreateFriendship(ctx context.Context, fr models.Friendship) error
	ReadFriends(ctx context.Context, id int) ([]models.User, error)
}
