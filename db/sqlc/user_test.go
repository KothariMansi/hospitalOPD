package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/KothariMansi/hospitalOPD/db/util"
	"github.com/stretchr/testify/require"
)

func CreateAndGetUser(t *testing.T) User {
	arg := CreateUserParams{
		Name:     util.RandomName(),
		Password: util.RandomPassword(),
		State:    util.RandomState(),
		City:     util.RandomCity(),
		Gender:   UserGender(util.RandomGender()),
		Age:      sql.NullInt32{Int32: util.RandomAge(), Valid: true},
	}

	result, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)

	id, err := result.LastInsertId()
	require.NoError(t, err)

	user, err := testQueries.GetUser(context.Background(), id)
	require.NoError(t, err)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.State, user.State)
	require.Equal(t, arg.City, user.City)
	require.Equal(t, arg.Gender, user.Gender)
	require.Equal(t, arg.Age, user.Age)

	return user
}

func TestCreateAndGetUser(t *testing.T) {
	CreateAndGetUser(t)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 3; i++ {
		CreateAndGetUser(t)
	}

	users, err := testQueries.ListUsers(context.Background(), ListUsersParams{
		Limit:  2,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Len(t, users, 2)
}

func TestUpdateUser(t *testing.T) {
	user := CreateAndGetUser(t)

	arg := UpdateUserParams{
		ID:       user.ID,
		Name:     util.RandomName(),
		Password: util.RandomPassword(),
		State:    util.RandomState(),
		City:     util.RandomCity(),
		Gender:   UserGender(util.RandomGender()),
		Age:      sql.NullInt32{Int32: util.RandomAge(), Valid: true},
	}

	err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)

	updated, err := testQueries.GetUser(context.Background(), user.ID)
	require.NoError(t, err)
	require.Equal(t, arg.Name, updated.Name)
	require.Equal(t, arg.City, updated.City)
	require.Equal(t, arg.State, updated.State)
	require.Equal(t, arg.Gender, updated.Gender)
}

func TestDeleteUser(t *testing.T) {
	user := CreateAndGetUser(t)

	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)

	deleted, err := testQueries.GetUser(context.Background(), user.ID)
	require.Error(t, err)
	require.Empty(t, deleted)
}
