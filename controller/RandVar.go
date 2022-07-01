package controller

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func PhoneRandom(spbPhoneSlice []string) (phone string) {

	var numBody string
	for i := 0; i <= 6; i++ {
		rand.Seed(time.Now().UnixMicro())
		rnd := rand.Intn(9)
		if i == 0 && rnd == 0 {
			rnd++
		}
		numBody = numBody + strconv.Itoa(rnd)

	}

	phone = fmt.Sprintf("+7 (%v) %v", spbPhoneSlice[rand.Intn(len(spbPhoneSlice))], numBody)
	return
}

func NameRandom(nameSlice []string) (name string) {

	rand.Seed(time.Now().UnixMicro())
	name = nameSlice[rand.Intn(len(nameSlice))]

	return
}


