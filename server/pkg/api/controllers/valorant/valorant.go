package valorant

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	constants "github.com/shanesaravia/podium/server/pkg/api/configs"
)

type cardData struct {
	Id    string `json:"id"`
	Small string `json:"small"`
	Large string `json:"large"`
	Wide  string `json:"wide"`
}

type userData struct {
	Puuid           string   `json:"puuid"`
	Region          string   `json:"region"`
	Account_level   uint32   `json:"account_level"`
	Name            string   `json:"name"`
	Tag             string   `json:"tag"`
	Last_update     string   `json:"last_update"`
	Last_update_raw uint32   `json:"last_update_raw"`
	Card            cardData `json:"card"`
}

type userResponse struct {
	Status uint16   `json:"status"`
	Data   userData `json:"data"`
}

func GetUserProfile(username string, tag string) userResponse {
	url := fmt.Sprintf(constants.Hendrick, username, tag)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	data := userResponse{}
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		log.Fatalln(jsonErr)
	}

	return data
}
