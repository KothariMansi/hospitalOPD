package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Random int generate a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generate random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomName() string {
	return RandomString(10)
}

func RandomState() string {
	return RandomString(6)
}

func RandomCity() string {
	return RandomString(7)
}

func RandomAge() int32 {
	return int32(RandomInt(3, 100))
}

// RandomGender returns a random gender string
func RandomGender() string {
	genders := []string{"MALE", "FEMALE", "OTHER"}
	return genders[rand.Intn(len(genders))]
}

// RandomUsername returns a realistic username
func RandomUsername() string {
	return fmt.Sprintf("user_%s", RandomString(6))
}

// RandomPassword returns a test password
func RandomPassword() string {
	return fmt.Sprintf("pass_%s", RandomString(8))
}

// RandomSpecialityName returns a random doctor speciality
func RandomSpecialityName() string {
	specs := []string{"Cardiology", "Neurology", "Orthopedics", "ENT", "Pediatrics"}
	return specs[rand.Intn(len(specs))]
}

// RandomPhone returns a dummy Indian-style phone number
func RandomPhone() string {
	return fmt.Sprintf("9%09d", rand.Intn(1e9))
}

// RandomAddress returns a test address string
func RandomAddress() string {
	return fmt.Sprintf("%d %s Street", rand.Intn(100), RandomString(5))
}

// RandomPhotoName returns a dummy image filename
func RandomPhotoName() string {
	return fmt.Sprintf("%s.jpg", RandomString(6))
}
