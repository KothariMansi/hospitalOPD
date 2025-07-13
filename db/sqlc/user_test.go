package db

import (
	"context"
	"database/sql"
	"fmt"
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

func TestCountUsers(t *testing.T) {
	for i := 0; i < 3; i++ {
		CreateAndGetUser(t)
	}

	count, err := testQueries.CountUsers(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, count, int64(3))
}

func TestSearchUsersByName(t *testing.T) {
	targetName := "TestUser_" + util.RandomString(4)
	for i := 0; i < 3; i++ {
		_, err := testQueries.CreateUser(context.Background(), CreateUserParams{
			Name:     fmt.Sprintf("%s_%d", targetName, i),
			Password: util.RandomPassword(),
			State:    util.RandomState(),
			City:     util.RandomCity(),
			Gender:   "MALE",
			Age:      util.SqlNullAge(util.RandomAge()),
		})
		require.NoError(t, err)
	}

	// Add noise
	for i := 0; i < 2; i++ {
		CreateAndGetUser(t)
	}

	users, err := testQueries.SearchUsersByName(context.Background(), SearchUsersByNameParams{
		Name:   targetName + "%",
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(users), 3)
	for _, user := range users {
		require.Contains(t, user.Name, targetName)
	}
}

func TestListUsersByGenderAndCity(t *testing.T) {
	city := "TestCity_" + util.RandomString(4)
	gender := "FEMALE"

	for i := 0; i < 3; i++ {
		_, err := testQueries.CreateUser(context.Background(), CreateUserParams{
			Name:     util.RandomName(),
			Password: util.RandomPassword(),
			State:    util.RandomState(),
			City:     city,
			Gender:   UserGender(gender),
			Age:      util.SqlNullAge(util.RandomAge()),
		})
		require.NoError(t, err)
	}

	// Add noise
	for i := 0; i < 2; i++ {
		CreateAndGetUser(t)
	}

	users, err := testQueries.ListUsersByGenderAndCity(context.Background(), ListUsersByGenderAndCityParams{
		Gender: UserGender(gender),
		City:   city,
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(users), 3)
	for _, user := range users {
		require.Equal(t, gender, string(user.Gender))
		require.Equal(t, city, user.City)
	}
}
