package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/KothariMansi/hospitalOPD/db/util"
	"github.com/stretchr/testify/require"
)

func CreateAndGetClient(t *testing.T) Client {
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
	CreateAndGetClient(t)
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
	client := CreateAndGetClient(t)

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
	client := CreateAndGetClient(t)

	err := testQueries.DeleteClient(context.Background(), client.ID)
	require.NoError(t, err)

	deleted, err := testQueries.GetClient(context.Background(), client.ID)
	require.Error(t, err)
	require.Empty(t, deleted)
}

func TestCountClients(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateAndGetClient(t)
	}

	count, err := testQueries.CountClients(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, count, int64(5))
}

func TestSearchClientsByName(t *testing.T) {
	// Create unique pattern-based clients
	targetName := "TestName_" + util.RandomString(3)
	for i := 0; i < 3; i++ {
		_, err := testQueries.CreateClient(context.Background(), CreateClientParams{
			Name:  fmt.Sprintf("%s_%d", targetName, i),
			State: util.RandomState(),
			City:  util.RandomCity(),
			Age:   util.RandomAge(),
		})
		require.NoError(t, err)
	}

	// Add some noise
	for i := 0; i < 2; i++ {
		CreateAndGetClient(t)
	}

	clients, err := testQueries.SearchClientsByName(context.Background(), SearchClientsByNameParams{
		Name:   targetName + "%",
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(clients), 3)
	for _, client := range clients {
		require.Contains(t, client.Name, targetName)
	}
}

func TestListClientsByLocation(t *testing.T) {
	state := "TestState_" + util.RandomString(4)
	city := "TestCity_" + util.RandomString(4)

	// Add matching clients
	for i := 0; i < 3; i++ {
		_, err := testQueries.CreateClient(context.Background(), CreateClientParams{
			Name:  util.RandomName(),
			State: state,
			City:  city,
			Age:   util.RandomAge(),
		})
		require.NoError(t, err)
	}

	// Add noise
	for i := 0; i < 2; i++ {
		CreateAndGetClient(t)
	}

	clients, err := testQueries.ListClientsByLocation(context.Background(), ListClientsByLocationParams{
		City:   city,
		State:  state,
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(clients), 3)
	for _, client := range clients {
		require.Equal(t, city, client.City)
		require.Equal(t, state, client.State)
	}
}
