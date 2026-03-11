package sqlstore_test

import (
	"testing"

	"github.com/mqksrwq/REST-API/internal/app/model"
	"github.com/mqksrwq/REST-API/internal/app/store"
	"github.com/mqksrwq/REST-API/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, dataBaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, dataBaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	email := "user@example.org"

	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser(t)
	u.Email = email
	err = s.User().Create(u)
	if err != nil {
		return
	}

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
