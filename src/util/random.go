package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt generate a random number in range (min, max)
func RandomInt(min, max int64) int64 {
	return min + random.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[random.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))
}

func RandomName() string {
	return RandomString(6)
}

func RandomPhoneNumber() string {
	number := RandomInt(1000000000, 9999999999)
	return strconv.FormatInt(number, 10)
}

func RandomFloat(min, max float64) string {
	return strconv.FormatFloat(min+random.Float64()*(max-min), 'f', 2, 32)
}

func RandomCity() string {
	cities := []string{SJ, SF, MV, GR}
	n := len(cities)
	return cities[random.Intn(n)]
}

func RandomState() string {
	states := []string{CA, TX, NV}
	n := len(states)
	return states[random.Intn(n)]
}
