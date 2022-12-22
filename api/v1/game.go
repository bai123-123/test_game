package v1

import (
	"SnakeLadderGame/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateNewGame(c *gin.Context) {
	var data model.Game
	err := c.ShouldBindJSON(&data)
	game, err := model.PlayGame(data.Player1, data.Player2)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"data":    game,
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"data":    game,
			"message": err,
		})
	}
}
