package db

import (
	"context"
	"database/sql"
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
