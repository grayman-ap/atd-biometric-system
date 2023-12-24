package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijlmnopqrstuvwxyz"

func Random(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomEmail() string {
	uq := "@gmailcom"
	email := Random(6)
	return email + uq
}
func RandomDepartment() string {
	dept := []string{"Computer Department", "Information Technology Department", "Information Media Department"}

	n := len(dept)

	return dept[rand.Intn(n)]
}

func RandomSchool() string {
	school := []string{"SEET", "SET", "Agric", "SICT", "URP", "SLS", "SPS"}

	n := len(school)

	return school[rand.Intn(n)]
}

func RandomCourseCode() string {

	course := []string{fmt.Sprintf("IFT%d", RandomInt(3, 200)), fmt.Sprintf("SVG%d", RandomInt(3, 200)), fmt.Sprintf("STA%d", RandomInt(3, 200)), fmt.Sprintf("IFT%d", RandomInt(3, 200)), fmt.Sprintf("GEO%d", RandomInt(3, 200)), fmt.Sprintf("IFT%d", RandomInt(3, 200)), fmt.Sprintf("IFT%d", RandomInt(3, 200)), fmt.Sprintf("IFT%d", RandomInt(1, 200)), fmt.Sprintf("IFT%d", RandomInt(1, 200)), fmt.Sprintf("IFT%d", RandomInt(1, 200)), fmt.Sprintf("IFT%d", RandomInt(1, 200)), fmt.Sprintf("IFT%d", RandomInt(1, 200))}

	n := len(course)

	return course[rand.Intn(n)]
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
