package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateAndGetCheckUpTime(t *testing.T) Checkuptime {
	now := time.Now().UTC().Truncate(time.Second)

	arg := CreateCheckUpTimeParams{
		Morning: sql.NullTime{Time: now.Add(8 * time.Hour), Valid: true},
		Evening: sql.NullTime{Time: now.Add(16 * time.Hour), Valid: true},
		Night:   sql.NullTime{Time: now.Add(22 * time.Hour), Valid: true},
	}

	result, err := testQueries.CreateCheckUpTime(context.Background(), arg)
	require.NoError(t, err)

	id, err := result.LastInsertId()
	require.NoError(t, err)

	record, err := testQueries.GetCheckUpTime(context.Background(), id)
	require.NoError(t, err)

	require.True(t, record.Morning.Valid)
	require.True(t, record.Evening.Valid)
	require.True(t, record.Night.Valid)

	require.WithinDuration(t, arg.Morning.Time, record.Morning.Time, time.Second)
	require.WithinDuration(t, arg.Evening.Time, record.Evening.Time, time.Second)
	require.WithinDuration(t, arg.Night.Time, record.Night.Time, time.Second)

	return record
}

func TestCreateAndGetCheckUpTime(t *testing.T) {
	CreateAndGetCheckUpTime(t)
}

func TestListCheckUpTimes(t *testing.T) {
	for i := 0; i < 3; i++ {
		CreateAndGetCheckUpTime(t)
	}

	records, err := testQueries.ListCheckUpTimes(context.Background(), ListCheckUpTimesParams{
		Limit:  2,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Len(t, records, 2)
}

func TestUpdateCheckUpTime(t *testing.T) {
	old := CreateAndGetCheckUpTime(t)

	newArg := UpdateCheckUpTimeParams{
		ID:      old.ID,
		Morning: sql.NullTime{Time: time.Now().Add(9 * time.Hour), Valid: true},
		Evening: sql.NullTime{Time: time.Now().Add(17 * time.Hour), Valid: true},
		Night:   sql.NullTime{Time: time.Now().Add(23 * time.Hour), Valid: true},
	}

	err := testQueries.UpdateCheckUpTime(context.Background(), newArg)
	require.NoError(t, err)

	updated, err := testQueries.GetCheckUpTime(context.Background(), old.ID)
	require.NoError(t, err)

	require.WithinDuration(t, newArg.Morning.Time, updated.Morning.Time, time.Second)
	require.WithinDuration(t, newArg.Evening.Time, updated.Evening.Time, time.Second)
	require.WithinDuration(t, newArg.Night.Time, updated.Night.Time, time.Second)
}

func TestDeleteCheckUpTime(t *testing.T) {
	record := CreateAndGetCheckUpTime(t)

	err := testQueries.DeleteCheckUpTime(context.Background(), record.ID)
	require.NoError(t, err)

	deleted, err := testQueries.GetCheckUpTime(context.Background(), record.ID)
	require.Error(t, err)
	require.Empty(t, deleted)
}
