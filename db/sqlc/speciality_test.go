package db

import (
	"context"
	"testing"

	"github.com/KothariMansi/hospitalOPD/db/util"
	"github.com/stretchr/testify/require"
)

func CreateAndGetSpeciality(t *testing.T) Speciality {
	arg := util.RandomSpecialityName()

	result, err := testQueries.CreateSpeciality(context.Background(), arg)
	require.NoError(t, err)

	id, err := result.LastInsertId()
	require.NoError(t, err)

	speciality, err := testQueries.GetSpeciality(context.Background(), id)
	require.NoError(t, err)

	require.Equal(t, arg, speciality.SpecialityName)

	return speciality
}

func TestCreateAndGetSpeciality(t *testing.T) {
	CreateAndGetSpeciality(t)
}

func TestListSpecialities(t *testing.T) {
	for i := 0; i < 3; i++ {
		CreateAndGetSpeciality(t)
	}

	specialities, err := testQueries.ListSpecialities(context.Background(), ListSpecialitiesParams{
		Limit:  2,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Len(t, specialities, 2)
}

func TestUpdateSpeciality(t *testing.T) {
	spec := CreateAndGetSpeciality(t)

	arg := UpdateSpecialityParams{
		ID:             spec.ID,
		SpecialityName: util.RandomSpecialityName(),
	}

	err := testQueries.UpdateSpeciality(context.Background(), arg)
	require.NoError(t, err)

	updated, err := testQueries.GetSpeciality(context.Background(), spec.ID)
	require.NoError(t, err)
	require.Equal(t, arg.SpecialityName, updated.SpecialityName)
}

func TestDeleteSpeciality(t *testing.T) {
	spec := CreateAndGetSpeciality(t)

	err := testQueries.DeleteSpeciality(context.Background(), spec.ID)
	require.NoError(t, err)

	deleted, err := testQueries.GetSpeciality(context.Background(), spec.ID)
	require.Error(t, err)
	require.Empty(t, deleted)
}
