package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/KothariMansi/hospitalOPD/db/util"
	"github.com/stretchr/testify/require"
)

func CreateAndGetDoctor(t *testing.T) Doctor {
	// Create dependencies
	hosp := CreateAndGetHospital(t)
	timeSlot := CreateAndGetCheckUpTime(t)

	arg := CreateDoctorParams{
		Name:            "Dr. " + util.RandomName(),
		Username:        util.RandomString(8),
		Password:        util.RandomPassword(),
		HospitalID:      hosp.ID,
		ResidentAddress: sql.NullString{String: util.RandomCity() + " Clinic", Valid: true},
		CheckupTimeID:   timeSlot.ID,
		IsOnLeave:       sql.NullBool{Bool: false, Valid: true},
	}

	result, err := testQueries.CreateDoctor(context.Background(), arg)
	require.NoError(t, err)

	id, err := result.LastInsertId()
	require.NoError(t, err)

	doctor, err := testQueries.GetDoctor(context.Background(), id)
	require.NoError(t, err)

	require.Equal(t, arg.Name, doctor.Name)
	require.Equal(t, arg.Username, doctor.Username)
	require.Equal(t, arg.Password, doctor.Password)
	require.Equal(t, arg.HospitalID, doctor.HospitalID)
	require.Equal(t, arg.CheckupTimeID, doctor.CheckupTimeID)
	require.Equal(t, arg.IsOnLeave, doctor.IsOnLeave)

	return doctor
}

func TestCreateAndGetDoctor(t *testing.T) {
	CreateAndGetDoctor(t)
}

func TestListDoctors(t *testing.T) {
	for i := 0; i < 3; i++ {
		CreateAndGetDoctor(t)
	}
	arg := ListDoctorsParams{
		Limit:  2,
		Offset: 0,
	}
	list, err := testQueries.ListDoctors(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, list, 2)
}

func TestUpdateDoctor(t *testing.T) {
	doc := CreateAndGetDoctor(t)
	newHosp := CreateAndGetHospital(t)
	newTime := CreateAndGetCheckUpTime(t)

	arg := UpdateDoctorParams{
		ID:              doc.ID,
		Name:            "Updated " + util.RandomName(),
		Username:        util.RandomString(8),
		Password:        util.RandomPassword(),
		HospitalID:      newHosp.ID,
		ResidentAddress: sql.NullString{String: util.RandomCity() + " Hall", Valid: true},
		CheckupTimeID:   newTime.ID,
		IsOnLeave:       sql.NullBool{Bool: true, Valid: true},
	}

	err := testQueries.UpdateDoctor(context.Background(), arg)
	require.NoError(t, err)

	updated, err := testQueries.GetDoctor(context.Background(), doc.ID)
	require.NoError(t, err)
	require.Equal(t, arg.Name, updated.Name)
	require.Equal(t, arg.IsOnLeave, updated.IsOnLeave)
	require.Equal(t, arg.HospitalID, updated.HospitalID)
}

func TestDeleteDoctor(t *testing.T) {
	doc := CreateAndGetDoctor(t)

	err := testQueries.DeleteDoctor(context.Background(), doc.ID)
	require.NoError(t, err)

	_, err = testQueries.GetDoctor(context.Background(), doc.ID)
	require.Error(t, err)
}

func TestCountDoctors(t *testing.T) {
	for i := 0; i < 3; i++ {
		CreateAndGetDoctor(t)
	}

	count, err := testQueries.CountDoctors(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, count, int64(3))
}

func TestSearchDoctorsByName(t *testing.T) {
	prefix := "Dr_" + util.RandomString(3)
	hospital := CreateAndGetHospital(t)
	time := CreateAndGetCheckUpTime(t)

	for i := 0; i < 3; i++ {
		_, err := testQueries.CreateDoctor(context.Background(), CreateDoctorParams{
			Name:            fmt.Sprintf("%s_%d", prefix, i),
			Username:        util.RandomUsername(),
			Password:        util.RandomPassword(),
			HospitalID:      hospital.ID,
			ResidentAddress: sql.NullString{String: util.RandomString(30), Valid: true},
			CheckupTimeID:   time.ID,
			IsOnLeave:       sql.NullBool{Bool: false, Valid: true},
		})
		require.NoError(t, err)
	}

	doctors, err := testQueries.SearchDoctorsByName(context.Background(), SearchDoctorsByNameParams{
		Name:   prefix + "%",
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(doctors), 3)
	for _, d := range doctors {
		require.Contains(t, d.Name, prefix)
	}
}

func TestListDoctorsWithHospital(t *testing.T) {
	// Make sure there is at least 1 hospital and doctor
	CreateAndGetDoctor(t)

	doctors, err := testQueries.ListDoctorsWithHospital(context.Background(), ListDoctorsWithHospitalParams{
		Limit:  5,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Greater(t, len(doctors), 0)

	for _, d := range doctors {
		require.NotEmpty(t, d.HospitalName)
		require.NotEmpty(t, d.Name)
	}
}

func TestListDoctorsByHospital(t *testing.T) {
	hospital := CreateAndGetHospital(t)
	time := CreateAndGetCheckUpTime(t)

	for i := 0; i < 2; i++ {
		_, err := testQueries.CreateDoctor(context.Background(), CreateDoctorParams{
			Name:            util.RandomName(),
			Username:        util.RandomUsername(),
			Password:        util.RandomPassword(),
			HospitalID:      hospital.ID,
			ResidentAddress: sql.NullString{String: util.RandomString(30), Valid: true},
			CheckupTimeID:   time.ID,
			IsOnLeave:       sql.NullBool{Bool: false, Valid: true},
		})
		require.NoError(t, err)
	}

	doctors, err := testQueries.ListDoctorsByHospital(context.Background(), ListDoctorsByHospitalParams{
		HospitalID: hospital.ID,
		Limit:      10,
		Offset:     0,
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(doctors), 2)
	for _, d := range doctors {
		require.Equal(t, hospital.ID, d.HospitalID)
	}
}

func TestListDoctorsOnLeave(t *testing.T) {
	hospital := CreateAndGetHospital(t)
	time := CreateAndGetCheckUpTime(t)

	_, err := testQueries.CreateDoctor(context.Background(), CreateDoctorParams{
		Name:            util.RandomName(),
		Username:        util.RandomUsername(),
		Password:        util.RandomPassword(),
		HospitalID:      hospital.ID,
		ResidentAddress: sql.NullString{String: util.RandomString(30), Valid: true},
		CheckupTimeID:   time.ID,
		IsOnLeave:       sql.NullBool{Bool: true, Valid: true},
	})
	require.NoError(t, err)

	doctors, err := testQueries.ListDoctorsOnLeave(context.Background(), ListDoctorsOnLeaveParams{
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.NotEmpty(t, doctors)
	for _, d := range doctors {
		require.True(t, d.IsOnLeave.Bool)
	}
}

func TestGetDoctorByUsername(t *testing.T) {
	doctor := CreateAndGetDoctor(t)

	d, err := testQueries.GetDoctorByUsername(context.Background(), doctor.Username)
	require.NoError(t, err)
	require.Equal(t, doctor.Username, d.Username)
}
