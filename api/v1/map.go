package v1

import (
	"SnakeLadderGame/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MakeMap(c *gin.Context) {
	var data model.Rule
	err := c.ShouldBindJSON(&data)
	jsons, err := json.Marshal(data)
	mmap := model.CreateMap(string(jsons))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"data":    mmap,
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"data":    mmap,
			"message": err,
		})
	}
}
