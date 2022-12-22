package v1

import (
	"SnakeLadderGame/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func GetReplay(c *gin.Context) {
	gameID := c.Query("game_id")
	id, err := strconv.ParseUint(gameID, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"data":    nil,
			"message": err,
		})
	}
	res, err := model.ReplayGame(uint(id))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"data":    res,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"data":    res,
			"message": "success",
		})
	}
}
