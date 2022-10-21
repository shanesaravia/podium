package valorant

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/shanesaravia/podium/server/pkg/api/controllers/valorant"
)

func UserSummary(c *gin.Context) {
	fmt.Println("data: ", c.Request.URL.Query())

	queryData := c.Request.URL.Query()
	username := queryData["username"][0]
	tag := queryData["tag"][0]

	profileData := valorant.GetUserProfile(username, tag)
	// resp, err := http.Get(constants.Hendrick)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// data := userResponse{}
	// jsonErr := json.Unmarshal(body, &data)
	// if jsonErr != nil {
	// 	log.Fatalln(jsonErr)
	// 	return
	// }

	c.JSON(http.StatusOK, profileData)
}
