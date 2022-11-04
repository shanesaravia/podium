package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shanesaravia/podium/server/pkg/api/routes"
)

func main() {
	router := gin.Default()
	routes.InitRouter(router)
	router.Run(":8000")
}
