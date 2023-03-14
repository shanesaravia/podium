//go:build go1.18
// +build go1.18

package valorant

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/shanesaravia/podium/server/pkg/api/configs/valorant"
	"gitlab.com/tymonx/go-formatter/formatter"
)

type AccountApiResponse struct {
	Status int         `json:"status"`
	Data   AccountData `json:"data"`
}

type AccountData struct {
	Puuid           string   `json:"puuid"`
	Region          string   `json:"region"`
	Account_level   uint32   `json:"account_level"`
	Name            string   `json:"name"`
	Tag             string   `json:"tag"`
	Last_update     string   `json:"last_update"`
	Last_update_raw uint32   `json:"last_update_raw"`
	Card            CardData `json:"card"`
}

type CardData struct {
	Id    string `json:"id"`
	Small string `json:"small"`
	Large string `json:"large"`
	Wide  string `json:"wide"`
}

type MMRApiResponse struct {
	Status uint16    `json:"status"`
	Data   []MMRData `json:"data"`
}

type MMRData struct {
	Rank       string `json:"currenttierpatched"`
	RankRating uint8  `json:"currenttier"`
	RR         uint8  `json:"ranking_in_tier"`
	Elo        int32  `json:"elo"`
}

type MatchesApiResponse struct {
	Status uint16        `json:"status"`
	Data   []MatchesData `json:"data"`
}

type MatchesData struct {
	Players struct {
		AllPlayers []Player `json:"all_players"`
	}
}
type Player struct {
	Name   string `json:"name"`
	Tag    string `json:"tag"`
	Damage uint32 `json:"damage_made"`
	Stats  struct {
		Score     uint32 `json:"score"`
		Kills     uint8  `json:"kills"`
		Deaths    uint8  `json:"deaths"`
		Assists   uint8  `json:"assists"`
		Headshots uint8  `json:"headshots"`
		Bodyshots uint8  `json:"bodyshots"`
		Legshots  uint8  `json:"legshots"`
	}
}

func GetAccount(username string, tag string) AccountApiResponse {
	url, err := formatter.Format(valorant.Urls["Account"], formatter.Named{
		"username": username,
		"tag":      tag,
	})
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := retryableHTTPRequest(3, 1, req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	userApiResponse := AccountApiResponse{}

	err = json.NewDecoder(resp.Body).Decode(&userApiResponse)
	if err != nil {
		log.Fatalln(err)
	}

	return userApiResponse
}

func GetMMRHistory(username string, tag string) MMRApiResponse {
	url, err := formatter.Format(valorant.Urls["MMR"], formatter.Named{
		"username": username,
		"tag":      tag,
	})
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	resp, err := retryableHTTPRequest(3, 1, req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	mmrApiResponse := MMRApiResponse{}
	err = json.NewDecoder(resp.Body).Decode(&mmrApiResponse)
	if err != nil {
		log.Fatalln(err)
	}
	return mmrApiResponse
}

func GetMatchHistory(username string, tag string) MatchesApiResponse {
	url, err := formatter.Format(valorant.Urls["Matches"], formatter.Named{
		"username": username,
		"tag":      tag,
	})

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := retryableHTTPRequest(3, 1, req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	matchesApiResponse := MatchesApiResponse{}
	err = json.NewDecoder(resp.Body).Decode(&matchesApiResponse)
	if err != nil {
		log.Fatalln(err)
	}

	return matchesApiResponse
}

func retryableHTTPRequest(retries uint8, sleep time.Duration, req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error
	for retries > 0 {
		resp, err = http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode == 429 {
			fmt.Printf("Received 429 status code. Retrying in %d second(s). %d retries left\n", sleep, retries)
			time.Sleep(sleep * time.Second)
			retries--
		} else {
			break
		}
	}
	return resp, err
}
