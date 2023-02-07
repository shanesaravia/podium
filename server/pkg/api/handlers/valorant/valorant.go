package valorant

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	client "github.com/shanesaravia/podium/server/pkg/api/clients/valorant"
	"github.com/shanesaravia/podium/server/pkg/api/controllers/valorant"
)

type UserSummaryResponse map[string]UserSummary

type UserSummary struct {
	Summary Summary `json:"summary"`
	Stats   Stats   `json:"stats"`
}

type Summary struct {
	Puuid         string          `json:"puuid"`
	Region        string          `json:"region"`
	AccountLevel  uint32          `json:"accountLevel"`
	Username      string          `json:"username"`
	Tag           string          `json:"tag"`
	LastUpdate    string          `json:"lastUpdate"`
	LastUpdateRaw uint32          `json:"lastUpdateRaw"`
	Card          client.CardData `json:"card"`
	CurrentRank   string          `json:"currentRank"`
	CurrentRR     uint8           `json:"currentRR"`
}

type Stats struct {
	EloChange             int32   `json:"eloChange"`
	RankRating            uint8   `json:"rankRating"`
	AverageScore          uint32  `json:"averageScore"`
	AverageDamagePerMatch uint32  `json:"averageDamagePerMatch"`
	HeadshotPercentage    float32 `json:"headshotPercentage"`
	KillDeathRatio        float32 `json:"killDeathRatio"`
	KillDeathAssistRatio  float32 `json:"killDeathAssistRatio"`
}

type User struct {
	Username string `json:"username"`
	Tag      string `json:"tag"`
}

func (u *User) getSummary() UserSummaryResponse {
	profileData := valorant.GetUserProfile(u.Username, u.Tag)
	mmrData := valorant.GetUserMMR(u.Username, u.Tag)
	matchData := valorant.GetUserMatches(u.Username, u.Tag)

	userSummary := UserSummary{
		Summary{
			Puuid:         profileData.Puuid,
			Region:        profileData.Region,
			AccountLevel:  profileData.Account_level,
			Username:      profileData.Name,
			Tag:           profileData.Tag,
			LastUpdate:    profileData.Last_update,
			LastUpdateRaw: profileData.Last_update_raw,
			Card:          profileData.Card,
			CurrentRank:   mmrData.Rank,
			CurrentRR:     mmrData.RR,
		},
		Stats{
			EloChange:             mmrData.EloChange,
			RankRating:            mmrData.RankRating,
			AverageScore:          matchData.AverageScore,
			AverageDamagePerMatch: matchData.AverageDamagePerMatch,
			HeadshotPercentage:    matchData.HeadshotPercentage,
			KillDeathRatio:        matchData.KillDeathRatio,
			KillDeathAssistRatio:  matchData.KillDeathAssistRatio,
		},
	}

	return UserSummaryResponse{strings.ToLower(profileData.Name): userSummary}
}

func UserSummaryList(c *gin.Context) {
	var requestBody []User
	var response []UserSummaryResponse
	c.BindJSON(&requestBody)
	for _, user := range requestBody {
		appendSummary(user, &response)
	}

	c.JSON(http.StatusOK, response)
}

func appendSummary(user User, response *[]UserSummaryResponse) {
	defer recovery()
	summary := user.getSummary()
	*response = append(*response, summary)
}

func recovery() {
	if err := recover(); err != nil {
		log.Println(err)
	}
}
