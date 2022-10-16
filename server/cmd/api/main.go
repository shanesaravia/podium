package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"/Users/shanesaravia/Development/podium/server/pkg/api/configs/constants.go"

	"github.com/gin-gonic/gin"
	"github.com/shanesaravia/podium/server/pkg/api/configs/constants"
	// "../../pkg/api/configs/constants"
)

func testEndpoint(c *gin.Context) {
	data, _ := json.Marshal(map[string]string{
		"client_id":     "play-valorant-web-prod",
		"nonce":         "1",
		"redirect_uri":  "https://playvalorant.com/opt_in",
		"response_type": "token id_token",
	})
	requestBody := bytes.NewBuffer(data)
	fmt.Println("here1")
	resp, err := http.Post(constants.Auth, "application/json", requestBody)
	fmt.Println("here2")

	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	fmt.Println("here3")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("failed here")
		log.Fatalln(err)
	}
	fmt.Println("here4")
	fmt.Println("body", body)
	fmt.Println("str body", string(body))
	sb := string(body)
	log.Printf(sb)
	c.JSON(http.StatusOK, gin.H{"data": sb})
}
