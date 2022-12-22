package v1

import (
	"SnakeLadderGame/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var data model.Player
	_ = c.ShouldBindJSON(&data)
	err := model.CreatePlayer(&data)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
		"msg":  err,
	})
}

func Login(c *gin.Context) {
	var formData model.Player
	_ = c.ShouldBindJSON(&formData)
	var err error
	formData, err = model.CheckLogin(formData.Email, formData.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  500,
			"data":    formData.Email,
			"id":      formData.ID,
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"data":    formData.Email,
			"id":      formData.ID,
			"message": err,
		})
	}

}
