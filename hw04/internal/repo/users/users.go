package users

import (
	"context"

	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/models"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/repo"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/storage"
)

var _ repo.UserRepo = (*users)(nil)

// users is an implementation of UserRepo interface.
type users struct {
	storage storage.UserStor
}

// New is a constructor for users.
func New(s storage.UserStor) *users {
	return &users{
		storage: s,
	}
}

// Create implements repo.UserRepo.
func (u *users) Create(ctx context.Context, user models.User) (int, error) {
	return u.storage.CreateUser(ctx, user)
}

// Read implements repo.UserRepo.
func (u *users) Read(ctx context.Context, id int) (models.User, error) {
	return u.storage.ReadUser(ctx, id)
}

// Update implements repo.UserRepo.
func (u *users) Update(ctx context.Context, user models.User) error {
	origUser, err := u.Read(ctx, user.ID)
	if err != nil {
		return err
	}

	if user.Name == "" {
		user.Name = origUser.Name
	}

	if user.Age == 0 {
		user.Age = origUser.Age
	}

	return u.storage.UpdateUser(ctx, user)
}

// Delete implements repo.UserRepo.
func (u *users) Delete(ctx context.Context, id int) error {
	_, err := u.storage.ReadUser(ctx, id)
	if err != nil {
		return err
	}
	return u.storage.DeleteUser(ctx, id)
}
