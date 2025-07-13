package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/KothariMansi/hospitalOPD/db/util"
	"github.com/stretchr/testify/require"
)

func CreateAndGetHospital(t *testing.T) Hospital {
	arg := CreateHospitalParams{
		Name:    util.RandomName(),
		Photo:   sql.NullString{String: util.RandomPhotoName(), Valid: true},
		State:   util.RandomState(),
		City:    util.RandomCity(),
		Address: util.RandomAddress(),
		Phone:   sql.NullString{String: util.RandomPhone(), Valid: true},
	}

	result, err := testQueries.CreateHospital(context.Background(), arg)
	require.NoError(t, err)

	id, err := result.LastInsertId()
	require.NoError(t, err)

	hospital, err := testQueries.GetHospital(context.Background(), id)
	require.NoError(t, err)

	require.Equal(t, arg.Name, hospital.Name)
	require.Equal(t, arg.Photo.String, hospital.Photo.String)
	require.Equal(t, arg.State, hospital.State)
	require.Equal(t, arg.City, hospital.City)
	require.Equal(t, arg.Address, hospital.Address)
	require.Equal(t, arg.Phone, hospital.Phone)

	return hospital
}

func TestCreateAndGetHospital(t *testing.T) {
	CreateAndGetHospital(t)
}

func TestListHospitals(t *testing.T) {
	for i := 0; i < 3; i++ {
		CreateAndGetHospital(t)
	}

	hospitals, err := testQueries.ListHospitals(context.Background(), ListHospitalsParams{
		Limit:  2,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Len(t, hospitals, 2)
}

func TestUpdateHospital(t *testing.T) {
	hospital := CreateAndGetHospital(t)

	arg := UpdateHospitalParams{
		ID:      hospital.ID,
		Name:    util.RandomName(),
		Photo:   sql.NullString{String: util.RandomPhotoName(), Valid: true},
		State:   util.RandomState(),
		City:    util.RandomCity(),
		Address: util.RandomAddress(),
		Phone:   sql.NullString{String: util.RandomPhone(), Valid: true},
	}

	err := testQueries.UpdateHospital(context.Background(), arg)
	require.NoError(t, err)

	updated, err := testQueries.GetHospital(context.Background(), hospital.ID)
	require.NoError(t, err)

	require.Equal(t, arg.Name, updated.Name)
	require.Equal(t, arg.Photo.String, updated.Photo.String)
	require.Equal(t, arg.State, updated.State)
	require.Equal(t, arg.City, updated.City)
	require.Equal(t, arg.Address, updated.Address)
	require.Equal(t, arg.Phone, updated.Phone)
}

func TestDeleteHospital(t *testing.T) {
	hospital := CreateAndGetHospital(t)

	err := testQueries.DeleteHospital(context.Background(), hospital.ID)
	require.NoError(t, err)

	deleted, err := testQueries.GetHospital(context.Background(), hospital.ID)
	require.Error(t, err)
	require.Empty(t, deleted)
}

func TestCountHospitals(t *testing.T) {
	for i := 0; i < 3; i++ {
		CreateAndGetHospital(t)
	}

	count, err := testQueries.CountHospitals(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, count, int64(3))
}

func TestSearchHospitalsByName(t *testing.T) {
	prefix := "TestHosp_" + util.RandomString(4)

	for i := 0; i < 3; i++ {
		_, err := testQueries.CreateHospital(context.Background(), CreateHospitalParams{
			Name:    fmt.Sprintf("%s_%d", prefix, i),
			Photo:   sql.NullString{String: util.RandomString(8), Valid: true},
			State:   util.RandomState(),
			City:    util.RandomCity(),
			Address: util.RandomString(20),
			Phone:   sql.NullString{String: util.RandomPhone(), Valid: true},
		})
		require.NoError(t, err)
	}

	// Add noise
	for i := 0; i < 2; i++ {
		CreateAndGetHospital(t)
	}

	hospitals, err := testQueries.SearchHospitalsByName(context.Background(), SearchHospitalsByNameParams{
		Name:   prefix + "%",
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(hospitals), 3)
	for _, h := range hospitals {
		require.Contains(t, h.Name, prefix)
	}
}

func TestListHospitalsByLocation(t *testing.T) {
	state := "TestState_" + util.RandomString(3)
	city := "TestCity_" + util.RandomString(3)

	for i := 0; i < 3; i++ {
		_, err := testQueries.CreateHospital(context.Background(), CreateHospitalParams{
			Name:    util.RandomName(),
			Photo:   sql.NullString{String: util.RandomString(10), Valid: true},
			State:   state,
			City:    city,
			Address: util.RandomString(30),
			Phone:   sql.NullString{String: util.RandomPhone(), Valid: true},
		})
		require.NoError(t, err)
	}

	// Add noise
	for i := 0; i < 2; i++ {
		CreateAndGetHospital(t)
	}

	hospitals, err := testQueries.ListHospitalsByLocation(context.Background(), ListHospitalsByLocationParams{
		State:  state,
		City:   city,
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(hospitals), 3)

	for _, h := range hospitals {
		require.Equal(t, state, h.State)
		require.Equal(t, city, h.City)
	}
}
