package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/errtype"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/models"
)

// CreateUser is a handler for GET /users/{user_id} route.
func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	var (
		user          models.User
		createdUserID int
		err           error
	)

	err = render.DecodeJSON(r.Body, &user)
	if err != nil || user.Name == "" || user.Age == 0 || user.ID != 0 {
		a.makeResponse(w, r, nil, errtype.ErrBadRequest, 0)
		return
	}

	createdUserID, err = a.users.Create(r.Context(), user)
	a.makeResponse(w, r, map[string]any{"id": createdUserID}, err, http.StatusCreated)
}

// ReadUser is a handler for POST /users route.
func (a *App) ReadUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "user_id")
	if id == "" {
		a.makeResponse(w, r, nil, errtype.ErrBadRequest, 0)
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		a.makeResponse(w, r, nil, errtype.ErrBadUserID, 0)
		return
	}

	user, err := a.users.Read(r.Context(), userID)
	a.makeResponse(w, r, map[string]any{"user": user}, err, http.StatusOK)
}

// UpdateUser is a handler for PUT /users route.
func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := render.DecodeJSON(r.Body, &user)
	if err != nil {
		a.makeResponse(w, r, nil, errtype.ErrBadRequest, 0)
		return
	}

	err = a.users.Update(r.Context(), user)
	if err != nil {
		a.makeResponse(w, r, nil, err, 0)
		return
	}

	a.makeResponse(w, r, map[string]any{"msg": fmt.Sprintf("user with id %d is successfully updated", user.ID)}, err, http.StatusOK)
}

// DeleteUser is a handler for DELETE /users route.
func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var (
		userReq DeleteUserReq
		err     error
	)

	err = render.DecodeJSON(r.Body, &userReq)
	if err != nil {
		a.makeResponse(w, r, nil, errtype.ErrBadRequest, 0)
		return
	}

	if userReq.ID == 0 {
		a.makeResponse(w, r, nil, errtype.ErrBadRequest, 0)
		return
	}

	err = a.users.Delete(r.Context(), userReq.ID)

	a.makeResponse(w, r, nil, err, http.StatusNoContent)
}

// CreateFriendship is a handler for POST /friends route.
func (a *App) CreateFriendship(w http.ResponseWriter, r *http.Request) {
	var (
		sourceUser, targetUser models.User
		fr                     models.Friendship
		err                    error
	)

	err = render.DecodeJSON(r.Body, &fr)
	if err != nil {
		a.makeResponse(w, r, nil, errtype.ErrBadRequest, 0)
		return
	}

	err = a.friendship.CreateFriendship(r.Context(), fr)
	if err != nil {
		a.makeResponse(w, r, nil, err, 0)
		return
	}

	sourceUser, err = a.users.Read(r.Context(), fr.SourceID)
	if err != nil {
		a.makeResponse(w, r, nil, err, 0)
		return
	}

	targetUser, err = a.users.Read(r.Context(), fr.TargetID)
	if err != nil {
		a.makeResponse(w, r, nil, err, 0)
		return
	}
	a.makeResponse(w, r, map[string]any{"msg": fmt.Sprintf("users %s and %s are friends now", sourceUser.Name, targetUser.Name)}, err, http.StatusOK)
}

// GetUserFriends is a handler for GET /friends/{user_id} route.
func (a *App) GetUserFriends(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "user_id")
	if id == "" {
		a.makeResponse(w, r, nil, errtype.ErrEmptyUserID, 0)
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		a.makeResponse(w, r, nil, errtype.ErrBadUserID, 0)
		return
	}

	friends, err := a.friendship.ReadFriends(r.Context(), userID)
	a.makeResponse(w, r, map[string]any{"friends": friends}, err, http.StatusOK)
}

// makeResponse is an util method which renders a response and sets its status.
func (a *App) makeResponse(w http.ResponseWriter, r *http.Request, res map[string]any, err error, okStatus int) {
	switch err {
	case nil:
		render.Status(r, okStatus)
		render.JSON(w, r, res)
		return
	case errtype.ErrBadRequest, errtype.ErrBadUserID, errtype.ErrEmptyUserID:
		render.Status(r, http.StatusBadRequest)
	case errtype.ErrUserNotFound, errtype.ErrFriendsNotFound:
		render.Status(r, http.StatusNotFound)
	default:
		render.Status(r, http.StatusInternalServerError)
	}
	render.JSON(w, r, map[string]any{"err": err.Error()})
}
