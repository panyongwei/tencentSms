package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateRandom() string {
	return fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000))
}
