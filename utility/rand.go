package utility

import (
	"github.com/gogf/gf/v2/util/gconv"
	"math/rand"
	"time"
)

func RandInt(max int) int {
	seed := time.Now().UnixMilli()
	randomNum := rand.New(rand.NewSource(seed))
	return randomNum.Intn(max)
}

// GetOrderNum 隨機訂單號
func GetOrderNum() (number string) {
	seed := time.Now().UnixMilli()
	randomNum := rand.New(rand.NewSource(seed))
	number = gconv.String(time.Now().UnixNano()) + gconv.String(randomNum.Intn(1000))
	return
}
