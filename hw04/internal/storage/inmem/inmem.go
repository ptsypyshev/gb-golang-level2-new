package inmem

import (
	"context"
	"sync"

	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/errtype"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/models"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/storage"
)

var _ storage.UserStor = (*MemStor)(nil)
var _ storage.FriendshipStor = (*MemStor)(nil)

type FriendsMap map[int]struct{}

// MemStor is an in-memory storage. It implements UserStor and FriendshipStor interfaces.
type MemStor struct {
	users      map[int]models.User
	friends    map[int]FriendsMap
	lastUserID int
	mu         sync.RWMutex
}

// New is a constructor for MemStor.
func New() *MemStor {
	return &MemStor{
		users:   make(map[int]models.User),
		friends: make(map[int]FriendsMap),
	}
}

// CreateUser implements storage.UserStor.
func (m *MemStor) CreateUser(ctx context.Context, u models.User) (int, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.lastUserID++
	u.ID = m.lastUserID
	m.users[u.ID] = u
	return u.ID, nil
}

// ReadUser implements storage.UserStor.
func (m *MemStor) ReadUser(ctx context.Context, id int) (models.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	u, ok := m.users[id]
	if !ok {
		return models.User{}, errtype.ErrUserNotFound
	}
	return u, nil
}

// UpdateUser implements storage.UserStor.
func (m *MemStor) UpdateUser(ctx context.Context, u models.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.users[u.ID]
	if !ok {
		return errtype.ErrUserNotFound
	}
	m.users[u.ID] = u
	return nil
}

// DeleteUser implements storage.UserStor.
func (m *MemStor) DeleteUser(ctx context.Context, id int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.users, id)
	return nil
}

// CreateFriendship implements storage.FriendshipStor.
func (m *MemStor) CreateFriendship(ctx context.Context, fr models.Friendship) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.addToFriend(fr.SourceID, fr.TargetID)
	m.addToFriend(fr.TargetID, fr.SourceID)
	return nil
}

// ReadFriends implements storage.FriendshipStor.
func (m *MemStor) ReadFriends(ctx context.Context, id int) ([]models.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	frMap, ok := m.friends[id]
	if !ok {
		return nil, errtype.ErrFriendsNotFound
	}
	res := make([]models.User, 0, len(frMap))
	for k := range frMap {
		u, err := m.ReadUser(ctx, k)
		if err != nil {
			continue
		}
		res = append(res, u)
	}
	if len(res) == 0 {
		return nil, errtype.ErrFriendsNotFound
	}
	return res, nil
}

// addToFriend creates friendship betweet source and target.
func (m *MemStor) addToFriend(sourceID, targetID int) {
	_, ok := m.friends[sourceID]
	if !ok {
		m.friends[sourceID] = FriendsMap{targetID: struct{}{}}
	} else {
		m.friends[sourceID][targetID] = struct{}{}
	}
}
