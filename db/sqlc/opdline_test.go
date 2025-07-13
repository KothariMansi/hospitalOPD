package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/KothariMansi/hospitalOPD/db/util"
	"github.com/stretchr/testify/require"
)

func CreateAndGetOPDLine(t *testing.T) Opdline {
	client := CreateAndGetClient(t)
	doctor := CreateAndGetDoctor(t)

	arg := CreateOPDLineParams{
		RegTime:     sql.NullTime{Time: time.Now(), Valid: true},
		TokenNumber: int32(util.RandomInt(1, 100)),
		ClientID:    client.ID,
		DoctorID:    doctor.ID,
		Ischecked:   false,
		CheckedTime: sql.NullTime{Valid: false}, // not checked yet
	}

	result, err := testQueries.CreateOPDLine(context.Background(), arg)
	require.NoError(t, err)

	id, err := result.LastInsertId()
	require.NoError(t, err)

	record, err := testQueries.GetOPDLine(context.Background(), id)
	require.NoError(t, err)
	require.Equal(t, arg.TokenNumber, record.TokenNumber)
	require.Equal(t, arg.ClientID, record.ClientID)
	require.Equal(t, arg.DoctorID, record.DoctorID)
	require.False(t, record.Ischecked)

	return record
}

func TestCreateAndGetOPDLine(t *testing.T) {
	CreateAndGetOPDLine(t)
}

func TestListOPDLines(t *testing.T) {
	for i := 0; i < 3; i++ {
		CreateAndGetOPDLine(t)
	}

	lines, err := testQueries.ListOPDLines(context.Background(), ListOPDLinesParams{
		Limit:  2,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Len(t, lines, 2)
}

func TestUpdateOPDLine(t *testing.T) {
	line := CreateAndGetOPDLine(t)
	client := CreateAndGetClient(t)
	doctor := CreateAndGetDoctor(t)

	arg := UpdateOPDLineParams{
		ID:          line.ID,
		RegTime:     sql.NullTime{Time: time.Now(), Valid: true},
		TokenNumber: int32(util.RandomInt(100, 200)),
		ClientID:    client.ID,
		DoctorID:    doctor.ID,
		Ischecked:   true,
		CheckedTime: sql.NullTime{Time: time.Now(), Valid: true},
	}

	err := testQueries.UpdateOPDLine(context.Background(), arg)
	require.NoError(t, err)

	updated, err := testQueries.GetOPDLine(context.Background(), line.ID)
	require.NoError(t, err)
	require.Equal(t, arg.Ischecked, updated.Ischecked)
	require.Equal(t, arg.TokenNumber, updated.TokenNumber)
	require.Equal(t, arg.ClientID, updated.ClientID)
	require.Equal(t, arg.DoctorID, updated.DoctorID)
	require.True(t, updated.CheckedTime.Valid)
}

func TestDeleteOPDLine(t *testing.T) {
	line := CreateAndGetOPDLine(t)

	err := testQueries.DeleteOPDLine(context.Background(), line.ID)
	require.NoError(t, err)

	_, err = testQueries.GetOPDLine(context.Background(), line.ID)
	require.Error(t, err)
}

// Count OPD Line

func TestListOPDWithDetails(t *testing.T) {
	client := CreateAndGetClient(t)
	doctor := CreateAndGetDoctor(t)

	for i := 0; i < 3; i++ {
		_, err := testQueries.CreateOPDLine(context.Background(), CreateOPDLineParams{
			RegTime:     sql.NullTime{Time: time.Now().Add(-time.Duration(i) * time.Hour), Valid: true},
			TokenNumber: int32(100 + i),
			ClientID:    client.ID,
			DoctorID:    doctor.ID,
			Ischecked:   false,
			CheckedTime: sql.NullTime{Valid: false},
		})
		require.NoError(t, err)
	}

	// Run the join query
	results, err := testQueries.ListOPDWithDetails(context.Background(), ListOPDWithDetailsParams{
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.NotEmpty(t, results)

	for _, row := range results {
		require.NotZero(t, row.ID)
		require.NotZero(t, row.RegTime)
		require.NotEmpty(t, row.ClientName)
		require.NotEmpty(t, row.DoctorName)
	}
}

func TestSearchOPDByDate(t *testing.T) {
	// Create test OPD records
	for i := 0; i < 3; i++ {
		CreateAndGetOPDLine(t)
	}

	start := time.Now().Add(-2 * time.Hour)
	end := time.Now().Add(2 * time.Hour)

	// Execute search query
	opdList, err := testQueries.SearchOPDByDate(context.Background(), SearchOPDByDateParams{
		FromRegTime: sql.NullTime{Time: start, Valid: true},
		ToRegTime:   sql.NullTime{Time: end, Valid: true},
		Limit:       5,
		Offset:      0,
	})
	require.NoError(t, err)
	require.NotEmpty(t, opdList)

	for _, o := range opdList {
		require.WithinDuration(t, start, o.RegTime.Time, 2*time.Hour)
		require.NotEmpty(t, o.ClientName)
		require.NotEmpty(t, o.DoctorName)
	}
}
