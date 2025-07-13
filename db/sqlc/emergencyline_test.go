package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/KothariMansi/hospitalOPD/db/util"
	"github.com/stretchr/testify/require"
)

func CreateAndGetEmergencyLine(t *testing.T) Emergencyline {
	client := CreateAndGetClient(t)
	doctor := CreateAndGetDoctor(t)

	arg := CreateEmergencyLineParams{
		RegTime:     sql.NullTime{Time: time.Now(), Valid: true},
		TokenNumber: util.RandomInt(101, 999),
		ClientID:    client.ID,
		DoctorID:    doctor.ID,
		Ischecked:   false,
		CheckedTime: sql.NullTime{Valid: false},
	}

	result, err := testQueries.CreateEmergencyLine(context.Background(), arg)
	require.NoError(t, err)

	id, err := result.LastInsertId()
	require.NoError(t, err)

	row, err := testQueries.GetEmergencyLine(context.Background(), id)
	require.NoError(t, err)

	require.Equal(t, arg.ClientID, row.ClientID)
	require.Equal(t, arg.DoctorID, row.DoctorID)
	require.Equal(t, arg.TokenNumber, row.TokenNumber)
	require.False(t, row.Ischecked)

	return row
}

func TestCreateAndGetEmergencyLine(t *testing.T) {
	CreateAndGetEmergencyLine(t)
}

func TestListEmergencyLines(t *testing.T) {
	for i := 0; i < 3; i++ {
		CreateAndGetEmergencyLine(t)
	}

	list, err := testQueries.ListEmergencyLines(context.Background(), ListEmergencyLinesParams{
		Limit:  2,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Len(t, list, 2)
}

func TestUpdateEmergencyLine(t *testing.T) {
	row := CreateAndGetEmergencyLine(t)
	client := CreateAndGetClient(t)
	doctor := CreateAndGetDoctor(t)

	arg := UpdateEmergencyLineParams{
		ID:          row.ID,
		RegTime:     sql.NullTime{Time: time.Now(), Valid: true},
		TokenNumber: util.RandomInt(200, 999),
		ClientID:    client.ID,
		DoctorID:    doctor.ID,
		Ischecked:   true,
		CheckedTime: sql.NullTime{Time: time.Now(), Valid: true},
	}

	err := testQueries.UpdateEmergencyLine(context.Background(), arg)
	require.NoError(t, err)

	updated, err := testQueries.GetEmergencyLine(context.Background(), row.ID)
	require.NoError(t, err)
	require.True(t, updated.Ischecked)
}

func TestDeleteEmergencyLine(t *testing.T) {
	row := CreateAndGetEmergencyLine(t)

	err := testQueries.DeleteEmergencyLine(context.Background(), row.ID)
	require.NoError(t, err)

	_, err = testQueries.GetEmergencyLine(context.Background(), row.ID)
	require.Error(t, err)
}

func TestCountEmergencyLines(t *testing.T) {
	client := CreateAndGetClient(t)
	doctor := CreateAndGetDoctor(t)

	// Insert some emergency records
	for i := 0; i < 2; i++ {
		_, err := testQueries.CreateEmergencyLine(context.Background(), CreateEmergencyLineParams{
			RegTime:     sql.NullTime{Time: time.Now().Add(-time.Duration(i) * time.Hour), Valid: true},
			TokenNumber: 200 + int64(i),
			ClientID:    client.ID,
			DoctorID:    doctor.ID,
			Ischecked:   false,
			CheckedTime: sql.NullTime{Valid: false},
		})
		require.NoError(t, err)
	}

	count, err := testQueries.CountEmergencyLines(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, count, int64(2))
}

func TestListEmergencyWithDetails(t *testing.T) {
	client := CreateAndGetClient(t)
	doctor := CreateAndGetDoctor(t)

	// Insert records for join
	for i := 0; i < 3; i++ {
		_, err := testQueries.CreateEmergencyLine(context.Background(), CreateEmergencyLineParams{
			RegTime:     sql.NullTime{Time: time.Now().Add(-time.Duration(i) * time.Hour), Valid: true},
			TokenNumber: 300 + int64(i),
			ClientID:    client.ID,
			DoctorID:    doctor.ID,
			Ischecked:   false,
			CheckedTime: sql.NullTime{Valid: false},
		})
		require.NoError(t, err)
	}

	records, err := testQueries.ListEmergencyWithDetails(context.Background(), ListEmergencyWithDetailsParams{
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(records), 3)

	for _, r := range records {
		require.NotZero(t, r.ID)
		require.NotEmpty(t, r.ClientName)
		require.NotEmpty(t, r.DoctorName)
	}
}

func TestSearchEmergencyByDate(t *testing.T) {
	// Create some records
	for i := 0; i < 3; i++ {
		CreateAndGetEmergencyLine(t)
	}

	// Time range
	start := time.Now().Add(-1 * time.Hour)
	end := time.Now().Add(1 * time.Hour)

	// Search
	emergencies, err := testQueries.SearchEmergencyByDate(context.Background(), SearchEmergencyByDateParams{
		FromRegTime: sql.NullTime{Time: start, Valid: true},
		ToRegTime:   sql.NullTime{Time: end, Valid: true},
		Limit:       5,
		Offset:      0,
	})
	require.NoError(t, err)
	require.NotEmpty(t, emergencies)

	for _, e := range emergencies {
		require.WithinDuration(t, start, e.RegTime.Time, time.Hour*2)
		require.NotEmpty(t, e.ClientName)
		require.NotEmpty(t, e.DoctorName)
	}
}
