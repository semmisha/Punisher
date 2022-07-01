package utils

import (
	"math/rand"
	"time"
)

func TimeRandom() time.Duration {

	rand.Seed(time.Now().UnixMicro())
	rndtime := rand.Intn(15)
	return time.Duration(rndtime)
}
