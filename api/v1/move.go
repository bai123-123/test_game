package v1

import (
	"SnakeLadderGame/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MakeAMove(c *gin.Context) {
	var data model.Move
	_ = c.ShouldBindJSON(&data)
	game, done, err := model.PlayerMove(data.PlayerID, data.GameID)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"data":    game,
			"message": err.Error(),
			"win":     done,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"data":    game,
			"message": err,
			"win":     done,
		})
	}
}
