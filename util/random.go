package util

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
	"time"
	db "github.com/HCMUT-UWC-2-0/backend/db/sqlc"

)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomNorDate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
func RandomNullDate() sql.NullTime {
	return sql.NullTime{
		Time:  RandomNorDate(),
		Valid: true,
	}
}

// generate a random social security number
func generateSSN() string {
	return fmt.Sprintf("%03d-%02d-%04d", rand.Intn(999), rand.Intn(99), rand.Intn(9999))
}

// generate a random name
func generateName(index int) string {
	names := []string{"John", "Jane", "Bob", "Alice", "David", "Mary", "Mike", "Lisa", "James", "Emily"}
	return names[index]
}

// generate a random phone number
func generatePhone() string {
	return fmt.Sprintf("+84 %d-%04d", rand.Intn(999)+100, rand.Intn(9999))
}

// generate a random age
func generateAge() int32 {
	return rand.Int31n(60) + 20
}

// generate a random worker type
func generateWorkerType() db.WorkerType {
	types := []db.WorkerType{db.WorkerTypeJANITOR, db.WorkerTypeCOLLECTOR}
	return types[rand.Intn(len(types))]
}

// generate a random gender
func generateGender() db.GenderType {
	genders := []db.GenderType{db.GenderTypeMALE, db.GenderTypeFEMALE}
	return genders[rand.Intn(len(genders))]
}

// generate a random date of birth
func generateDateOfBirth() time.Time {
	min := time.Date(1960, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	sec := rand.Int63n(max-min) + min
	return time.Unix(sec, 0)
}

// generate a random place of birth
func generatePlaceOfBirth() string {
	places := []string{"Ha Noi", "Ho Chi Minh City", "Da Nang", "Hai Phong", "Can Tho", "Bien Hoa", "Vung Tau", "Nha Trang", "Phan Thiet", "Da Lat"}
	return places[rand.Intn(len(places))]
}


// define a function to generate a random worker
func RandomWorker(index int) db.CreateWorkerParams {
	ssn := generateSSN()
	name := generateName(index)
	phone := generatePhone()
	age := generateAge()
	workerType := generateWorkerType()
	gender := generateGender()
	dateOfBirth := generateDateOfBirth()
	placeOfBirth := generatePlaceOfBirth()

	return db.CreateWorkerParams{
		Ssn:            ssn,
		Name:           name,
		Phone:          phone,
		Age:            age,
		WorkerType:     workerType,
		Gender:         gender,
		DateOfBirth:    dateOfBirth,
		PlaceOfBirth:   placeOfBirth,
	}
}