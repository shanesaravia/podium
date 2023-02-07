package valorant

import "fmt"

const BASE_URL = "https://api.henrikdev.xyz/valorant"

var Urls = map[string]string{
	"Account": fmt.Sprintf("%s/v1/account/{username}/{tag}", BASE_URL),
	"MMR":     fmt.Sprintf("%s/v1/mmr-history/na/{username}/{tag}", BASE_URL),
	"Matches": fmt.Sprintf("%s/v3/matches/na/{username}/{tag}", BASE_URL),
}
