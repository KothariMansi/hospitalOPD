package db

import (
	"context"
	"testing"

	"github.com/KothariMansi/hospitalOPD/db/util"
	"github.com/stretchr/testify/require"
)

func createRandomClient(t *testing.T) Client {
	arg := CreateClientParams{
		Name:  util.RandomName(),
		State: util.RandomState(),
		City:  util.RandomCity(),
		Age:   util.RandomAge(),
	}
	result, err := testQueries.CreateClient(context.Background(), arg)
	require.NoError(t, err)
	id, err := result.LastInsertId()
	require.NoError(t, err)
	client, err := testQueries.GetClient(context.Background(), id)
	require.NoError(t, err)
	require.Equal(t, arg.Name, client.Name)
	require.Equal(t, arg.State, client.State)
	require.Equal(t, arg.City, client.City)
	require.Equal(t, arg.Age, client.Age)
	return client
}

func TestCreateAndGetClient(t *testing.T) {
	createRandomClient(t)
}

func TestListClients(t *testing.T) {
	for i := 0; i < 3; i++ {
		_, err := testQueries.CreateClient(context.Background(), CreateClientParams{
			Name:  util.RandomName(),
			State: util.RandomState(),
			City:  util.RandomCity(),
			Age:   util.RandomAge(),
		})
		require.NoError(t, err)
	}

	clients, err := testQueries.ListClients(context.Background(), ListClientsParams{
		Limit:  2,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Len(t, clients, 2)
}

func TestUpdateClient(t *testing.T) {
	client := createRandomClient(t)

	newAge := util.RandomAge()
	newCity := util.RandomCity()
	newState := util.RandomState()

	err := testQueries.UpdateClient(context.Background(), UpdateClientParams{
		ID:    client.ID,
		Name:  client.Name,
		State: newState,
		City:  newCity,
		Age:   int32(newAge),
	})
	require.NoError(t, err)

	updated, err := testQueries.GetClient(context.Background(), client.ID)
	require.NoError(t, err)
	require.Equal(t, newState, updated.State)
	require.Equal(t, newCity, updated.City)
	require.Equal(t, newAge, updated.Age)
}

func TestDeleteClient(t *testing.T) {
	client := createRandomClient(t)

	err := testQueries.DeleteClient(context.Background(), client.ID)
	require.NoError(t, err)

	deleted, err := testQueries.GetClient(context.Background(), client.ID)
	require.Error(t, err)
	require.Empty(t, deleted)
}
