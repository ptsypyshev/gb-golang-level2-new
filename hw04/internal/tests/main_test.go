package tests

import (
	"context"
	"database/sql"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/app"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/repo/friends"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/repo/users"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/storage/pgdb"
	"github.com/ptsypyshev/gb-golang-level2-new/hw04/internal/usecase/friendship"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	user1 = `{"user":{"id":1,"name":"ivan","age":20}}
`
	badUser = `{"err":"bad user ID"}
`
	id4Response = `{"id":4}
`
	updatedID4Response = `{"msg":"user with id 4 is successfully updated"}
`
)

type IntegrationTestSuite struct {
	suite.Suite
	db   *sql.DB
	stor *pgdb.PostgresStore
}

func (s *IntegrationTestSuite) SetupSuite() {
}

func (s *IntegrationTestSuite) TearDownSuite() {
}

func (s *IntegrationTestSuite) SetupTest() {
}

func TestIntegrationSuite(t *testing.T) {
	s := new(IntegrationTestSuite)

	pgPool, pgRes := Start()
	defer Stop(pgPool, pgRes)

	stor, err := pgdb.New("postgres://postgres:postgres@localhost:5432/friends?sslmode=disable")
	if err != nil {
		log.Panicf("cannot init DB: %s", err)
	}
	defer stor.Close()

	s.db = stor.GetDB()
	s.stor = stor

	suite.Run(t, s)
}

func (s *IntegrationTestSuite) createData(t *testing.T) {
	t.Helper()
	time.Sleep(3 * time.Second)

	_, err := s.db.Exec(`
	CREATE TABLE users (
		id          SERIAL PRIMARY KEY,
		name        varchar    NOT NULL,
		age         INT    NOT NULL,
		created_at  timestamp NOT NULL DEFAULT NOW(),
		updated_at  timestamp NOT NULL DEFAULT NOW()
	)`)
	assert.NoError(t, err)

	_, err = s.db.Exec(`
	CREATE TABLE friendships (
		id        SERIAL PRIMARY KEY,
		source_id INT    NOT NULL,
		target_id INT    NOT NULL
	)`)
	assert.NoError(t, err)

	_, err = s.db.Exec("INSERT INTO users (name, age) VALUES ('ivan', 20), ('petr', 21), ('alex', 33)")
	assert.NoError(t, err)
}

func (s *IntegrationTestSuite) TestHandlers() {
	t := s.T()
	s.createData(t)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	uRepo := users.New(s.stor)
	fRepo := friends.New(s.stor)
	fUsecase := friendship.New(uRepo, fRepo)

	a := app.New(uRepo, fUsecase)
	a.SetupRoutes()

	go a.Run(ctx)
	time.Sleep(time.Second)

	t.Run("Create User", func(t *testing.T) {
		var client http.Client

		reqBody := `{"name": "Pavel", "age": 39}`
		req, err := http.NewRequest(http.MethodPost, "http://localhost:8000/users", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		assert.NoError(t, err)

		resp, err := client.Do(req)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		assert.NoError(t, err)

		defer resp.Body.Close()

		resBody, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.Equal(t, id4Response, string(resBody))
	})

	t.Run("Read User", func(t *testing.T) {
		var client http.Client

		req, err := http.NewRequest(http.MethodGet, "http://localhost:8000/users/1", nil)
		req.Header.Set("Content-Type", "application/json")

		assert.NoError(t, err)

		resp, err := client.Do(req)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NoError(t, err)

		defer resp.Body.Close()

		resBody, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.Equal(t, user1, string(resBody))
	})

	t.Run("Read User Bad", func(t *testing.T) {
		var client http.Client

		req, err := http.NewRequest(http.MethodGet, "http://localhost:8000/users/hi", nil)
		req.Header.Set("Content-Type", "application/json")

		assert.NoError(t, err)

		resp, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		defer resp.Body.Close()

		resBody, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.Equal(t, badUser, string(resBody))
	})

	t.Run("Update User", func(t *testing.T) {
		var client http.Client

		reqBody := `{"id":4, "name": "Admin", "age": 30}`
		req, err := http.NewRequest(http.MethodPut, "http://localhost:8000/users", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		assert.NoError(t, err)

		resp, err := client.Do(req)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NoError(t, err)

		defer resp.Body.Close()

		resBody, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.Equal(t, updatedID4Response, string(resBody))
	})

	t.Run("Delete User", func(t *testing.T) {
		var client http.Client

		reqBody := `{"target_id":4}`
		req, err := http.NewRequest(http.MethodDelete, "http://localhost:8000/users", strings.NewReader(reqBody))
		req.Header.Set("Content-Type", "application/json")

		assert.NoError(t, err)

		resp, err := client.Do(req)
		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
		assert.NoError(t, err)
	})
}
