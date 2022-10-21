package routes

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	vr := v1.Group("/valorant")
	ValorantRoutes(vr)
}
