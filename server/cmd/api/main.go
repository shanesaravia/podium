package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	constants "github.com/shanesaravia/podium/server/pkg/api/configs"
	"github.com/shanesaravia/podium/server/pkg/api/routes"
)

func main() {
	fmt.Println("testing")
	fmt.Println(constants.Hendrick)

	router := gin.Default()
	routes.InitRouter(router)
	router.Run(":8000")
}

func testEndpoint(c *gin.Context) {
	resp, err := http.Get(constants.Hendrick)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("resp: ", resp)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("body", body)
	fmt.Println("body string", string(body))
	// sb := string(body)
	// log.Printf(sb)
	c.JSON(http.StatusOK, gin.H{"data": string(body)})

	// data, _ := json.Marshal(map[string]string{
	// 	"client_id":     "play-valorant-web-prod",
	// 	"nonce":         "1",
	// 	"redirect_uri":  "https://playvalorant.com/opt_in",
	// 	"response_type": "token id_token",
	// })
	// requestBody := bytes.NewBuffer(data)
	// fmt.Println("here1")
	// resp, err := http.Post(constants.Hendrick, "application/json", requestBody)
	// fmt.Println("here2")

	// //Handle Error
	// if err != nil {
	// 	log.Fatalf("An Error Occured %v", err)
	// }
	// defer resp.Body.Close()
	// //Read the response body
	// fmt.Println("here3")
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("failed here")
	// 	log.Fatalln(err)
	// }
	// fmt.Println("here4")
	// fmt.Println("body", body)
	// fmt.Println("str body", string(body))
	// sb := string(body)
	// log.Printf(sb)
	// c.JSON(http.StatusOK, gin.H{"data": sb})
}
