package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateAndGetDoctorSpeciality(t *testing.T) Doctorspeciality {
	spec := CreateAndGetSpeciality(t)
	doc := CreateAndGetDoctor(t)

	arg := CreateDoctorSpecialityParams{
		SpecialityID: spec.ID,
		DocterID:     doc.ID,
	}

	result, err := testQueries.CreateDoctorSpeciality(context.Background(), arg)
	require.NoError(t, err)

	id, err := result.LastInsertId()
	require.NoError(t, err)

	row, err := testQueries.GetDoctorSpeciality(context.Background(), id)
	require.NoError(t, err)

	require.Equal(t, arg.SpecialityID, row.SpecialityID)
	require.Equal(t, arg.DocterID, row.DocterID)
	return row
}

func TestCreateAndGetDoctorSpeciality(t *testing.T) {
	CreateAndGetDoctorSpeciality(t)
}

func TestListDoctorSpecialities(t *testing.T) {
	for i := 0; i < 3; i++ {
		CreateAndGetDoctorSpeciality(t)
	}

	list, err := testQueries.ListDoctorSpecialities(context.Background(), ListDoctorSpecialitiesParams{
		Limit:  2,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Len(t, list, 2)
}

func TestUpdateDoctorSpeciality(t *testing.T) {
	row := CreateAndGetDoctorSpeciality(t)
	newSpec := CreateAndGetSpeciality(t)
	newDoc := CreateAndGetDoctor(t)

	arg := UpdateDoctorSpecialityParams{
		ID:           row.ID,
		SpecialityID: newSpec.ID,
		DocterID:     newDoc.ID,
	}

	err := testQueries.UpdateDoctorSpeciality(context.Background(), arg)
	require.NoError(t, err)

	updated, err := testQueries.GetDoctorSpeciality(context.Background(), row.ID)
	require.NoError(t, err)
	require.Equal(t, arg.SpecialityID, updated.SpecialityID)
	require.Equal(t, arg.DocterID, updated.DocterID)
}

func TestDeleteDoctorSpeciality(t *testing.T) {
	row := CreateAndGetDoctorSpeciality(t)

	err := testQueries.DeleteDoctorSpeciality(context.Background(), row.ID)
	require.NoError(t, err)

	_, err = testQueries.GetDoctorSpeciality(context.Background(), row.ID)
	require.Error(t, err)
}

func TestListSpecialitiesByDoctorID(t *testing.T) {
	doctor := CreateAndGetDoctor(t)

	// Create multiple specialities
	for i := 0; i < 3; i++ {
		spec := CreateAndGetSpeciality(t)

		_, err := testQueries.CreateDoctorSpeciality(context.Background(), CreateDoctorSpecialityParams{
			SpecialityID: spec.ID,
			DocterID:     doctor.ID,
		})
		require.NoError(t, err)
	}

	specialities, err := testQueries.ListSpecialitiesByDoctorID(context.Background(), doctor.ID)
	require.NoError(t, err)
	require.Len(t, specialities, 3)

	for _, s := range specialities {
		require.NotEmpty(t, s.SpecialityName)
	}
}

func TestListDoctorsBySpecialityID(t *testing.T) {
	spec := CreateAndGetSpeciality(t)

	// Create and assign multiple doctors
	for i := 0; i < 3; i++ {
		doc := CreateAndGetDoctor(t)
		_, err := testQueries.CreateDoctorSpeciality(context.Background(), CreateDoctorSpecialityParams{
			SpecialityID: spec.ID,
			DocterID:     doc.ID,
		})
		require.NoError(t, err)
	}

	doctors, err := testQueries.ListDoctorsBySpecialityID(context.Background(), ListDoctorsBySpecialityIDParams{
		SpecialityID: spec.ID,
		Limit:        5,
		Offset:       0,
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(doctors), 3)
	for _, d := range doctors {
		require.NotEmpty(t, d.Name)
	}
}
