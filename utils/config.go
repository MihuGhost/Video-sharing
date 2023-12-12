package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateId() string {
	formatNow := time.Unix(time.Now().Unix(), 0).Format("20060102150405")
	randNumber := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	return fmt.Sprintf("%s%s", formatNow, randNumber)
}
