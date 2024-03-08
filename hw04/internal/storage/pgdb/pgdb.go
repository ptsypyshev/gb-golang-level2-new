package pgdb

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/errtype"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/models"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/storage"
)

var _ storage.UserStor = (*PostgresStore)(nil)
var _ storage.FriendshipStor = (*PostgresStore)(nil)

// PostgresStore is a PostgreSQL-based storage. It implements UserStor and FriendshipStor interfaces.
type PostgresStore struct {
	db *sql.DB
}

// New is a constructor for PostgresStore.
func New(connStr string) (*PostgresStore, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

// GetDB returns connection to DB.
func (s *PostgresStore) GetDB() *sql.DB {
	return s.db
}

// Close releases connection to DB.
func (s *PostgresStore) Close() {
	s.db.Close()
}

// CreateUser implements storage.UserStor.
func (s *PostgresStore) CreateUser(ctx context.Context, u models.User) (int, error) {
	query := "INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id"
	var id int
	err := s.db.QueryRowContext(ctx, query, u.Name, u.Age).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// ReadUser implements storage.UserStor.
func (s *PostgresStore) ReadUser(ctx context.Context, id int) (models.User, error) {
	var user models.User
	query := "SELECT id, name, age FROM users WHERE id = $1"
	err := s.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// UpdateUser implements storage.UserStor.
func (s *PostgresStore) UpdateUser(ctx context.Context, u models.User) error {
	switch {
	case u.Name != "" && u.Age != 0:
		query := "UPDATE users SET name = $1, age = $2 WHERE id = $3"
		_, err := s.db.ExecContext(ctx, query, u.Name, u.Age, u.ID)
		if err != nil {
			return err
		}
	case u.Name != "":
		query := "UPDATE users SET name = $1 WHERE id = $2"
		_, err := s.db.ExecContext(ctx, query, u.Name, u.ID)
		if err != nil {
			return err
		}
	case u.Name != "":
		query := "UPDATE users SET age = $1 WHERE id = $2"
		_, err := s.db.ExecContext(ctx, query, u.Age, u.ID)
		if err != nil {
			return err
		}
	default:
		return errtype.ErrBadRequest
	}

	return nil
}

// DeleteUser implements storage.UserStor.
func (s *PostgresStore) DeleteUser(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

// CreateFriendship implements storage.FriendshipStor.
func (s *PostgresStore) CreateFriendship(ctx context.Context, fr models.Friendship) error {
	query := "INSERT INTO friendships (source_id, target_id) VALUES ($1, $2),($2, $1)"
	_, err := s.db.ExecContext(ctx, query, fr.SourceID, fr.TargetID)
	if err != nil {
		return err
	}
	return nil
}

// ReadFriends implements storage.FriendshipStor.
func (s *PostgresStore) ReadFriends(ctx context.Context, id int) ([]models.User, error) {
	var friends []models.User
	query := "SELECT u.id, u.name, u.age FROM friendships f INNER JOIN users u ON f.target_id = u.id WHERE f.source_id = $1"
	rows, err := s.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var friend models.User
		err := rows.Scan(&friend.ID, &friend.Name, &friend.Age)
		if err != nil {
			return nil, err
		}
		friends = append(friends, friend)
	}
	return friends, nil
}
