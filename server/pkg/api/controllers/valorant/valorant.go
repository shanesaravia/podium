package valorant

import (
	"errors"
	"math"
	"net/http"
	"strings"

	"github.com/shanesaravia/podium/server/pkg/api/clients/valorant"
)

type MMRData struct {
	Rank       string
	RankRating uint8
	RR         uint8
	EloChange  int32
}

type PlayerData struct {
	AverageScore          uint32
	AverageDamagePerMatch uint32
	HeadshotPercentage    float32
	KillDeathRatio        float32
	KillDeathAssistRatio  float32
}

func GetUserProfile(username string, tag string) (valorant.AccountApiResponse, error) {
	accountData := valorant.GetAccount(username, tag)
	if accountData.Status != http.StatusOK {
		return accountData, errors.New("Unable to fetch user profile")
	}
	return accountData, nil
}

func GetUserMMR(username string, tag string) MMRData {
	mmrHistory := valorant.GetMMRHistory(username, tag)
	currentData := mmrHistory.Data[0]
	oldestData := mmrHistory.Data[len(mmrHistory.Data)-1]

	mmrData := MMRData{
		Rank:       currentData.Rank,
		RankRating: currentData.RankRating,
		RR:         currentData.RR,
		EloChange:  currentData.Elo - oldestData.Elo,
	}
	return mmrData
}

func GetUserMatches(username string, tag string) PlayerData {
	matchHistory := valorant.GetMatchHistory(username, tag)

	var score uint32
	var damage uint32
	var headshotPercentage float32
	var killsDeaths float32
	var killsDeathsAssists float32

	for _, match := range matchHistory.Data {
		allPlayers := match.Players.AllPlayers
		for _, player := range allPlayers {
			if strings.EqualFold(player.Name, username) && strings.EqualFold(player.Tag, tag) {
				score += player.Stats.Score
				damage += player.Damage
				totalShots := player.Stats.Headshots + player.Stats.Bodyshots + player.Stats.Legshots
				hsPercent := float32(player.Stats.Headshots) / float32(totalShots) * 100
				if math.IsNaN(float64(hsPercent)) {
					hsPercent = float32(0)
				}
				headshotPercentage += hsPercent
				kd := float32(player.Stats.Kills) / float32(player.Stats.Deaths)
				killsDeaths += kd
				kda := float32(player.Stats.Kills+player.Stats.Assists) / float32(player.Stats.Deaths)
				killsDeathsAssists += kda
				break
			}
		}
	}
	matchHistoryLength := len(matchHistory.Data)
	player := PlayerData{}

	if matchHistoryLength > 0 {
		player = PlayerData{
			AverageScore:          score / uint32(matchHistoryLength),
			AverageDamagePerMatch: damage / uint32(matchHistoryLength),
			HeadshotPercentage:    round(headshotPercentage / float32(matchHistoryLength)),
			KillDeathRatio:        round(killsDeaths / float32(matchHistoryLength)),
			KillDeathAssistRatio:  round(killsDeathsAssists / float32(matchHistoryLength)),
		}
	}

	return player
}

func round(value float32) float32 {
	return float32(math.Round(float64(value*100)) / 100)
}
