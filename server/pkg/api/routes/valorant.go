package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shanesaravia/podium/server/pkg/api/handlers/valorant"
)

func ValorantRoutes(rg *gin.RouterGroup) {
	rg.POST("/summary", valorant.UserSummaryList)
	rg.POST("/profile", valorant.UserProfile)
}
