package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateAccountNo() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("2025%08d", rand.Intn(100000000))
}
