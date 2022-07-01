package main

import (
	"Punisher/controller"
	"Punisher/logging"
	"Punisher/repository"
	"Punisher/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	punURL     = "https://vagner-skoda.ru/form/2"
	namesPath  = "/app/files/names.txt"
	phonesPath = "/app/files/phones.txt"
	token      = "klPGmOEqsWYpBodl3NkMgHJnxqxArxingzAh8pcE"
)

type Skoda struct {
	Token     string `json:"_token,omitempty"`
	Name      string `json:"name,omitempty"`
	Telephone string `json:"telephone,omitempty"`
}

func main() {
	logger := logging.Logger()
	data := NewSkoda()
	names := repository.FileToVar(namesPath)
	phones := repository.FileToVar(phonesPath)
	//
	for {
		msk, err := time.LoadLocation("Europe/Moscow")
		if err != nil {
			logger.Warnf("\nUnable to set location MSK +3, error:%v\n", err)

		}
		if time.Now().UTC().Hour() >= 12 && time.Now().UTC().Hour() <= 23 {

			data.Fill(token, names, phones)

			mdata, err := json.Marshal(data)

			r := strings.NewReader(fmt.Sprint(mdata))

			if err != nil {
				logger.Panicf("\nUnable to marshall, error%v\n", err)

			}

			resp, err := http.Post(punURL, "application/x-www-form-urlencoded", r)
			if err != nil {
				logger.Panicf("\nUnable to receive respnonse, error%v\n", err)

			}

			logger.Printf("Time: %v Name: %v Phone:%v ApiResponse status: %v\n ", time.Now().UTC().Format(time.Kitchen), data.Name, data.Telephone, resp.StatusCode)
			err = resp.Body.Close()
			if err != nil {
				logger.Panicf("\nUnable to close body, error%v\n", err)

			}
			randDuration := utils.TimeRandom()
			time.Sleep(randDuration * time.Minute)
		}
		fmt.Println("Z-z-z-z-z-z")
		time.Sleep(30 * time.Minute)
	}

}

func NewSkoda() *Skoda {
	return &Skoda{
		Token:     "",
		Name:      "",
		Telephone: "",
	}
}

func (s *Skoda) Fill(token string, names, phones []string) {

	*s = Skoda{
		Token:     token,
		Name:      controller.NameRandom(names),
		Telephone: controller.PhoneRandom(phones),
	}

}
