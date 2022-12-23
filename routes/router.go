package routes

import (
	v1 "SnakeLadderGame/api/v1"
	"SnakeLadderGame/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.POST("register", v1.Register)
		router.POST("login", v1.Login)
		router.POST("newGame", v1.CreateNewGame)
		router.POST("move", v1.MakeAMove)
		router.GET("replay", v1.GetReplay)
		router.POST("makeMap", v1.MakeMap)
	}

	r.Run(utils.HttpPort)
}
