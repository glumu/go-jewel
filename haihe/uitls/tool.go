package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandnum() int {
	i := time.Time.Nanosecond(time.Now())
	return i
}

func IsMockApi(mockApi []string, api string) bool {
	if mockApi != nil {
		for _, val := range mockApi {
			if val == api {
				return true
			}
		}
	}
	return false
}

func GenMockId() string {
	return fmt.Sprintf("cccccccc-cccc-cccc-cccc-%s", GetRandomString(12))
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
