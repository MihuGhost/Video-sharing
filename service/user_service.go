package service

import (
	"OnlineVideo/models"
	"OnlineVideo/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func CreateUser(c *gin.Context) {
	user := models.User{}
	user.Name = c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	repassword := c.Request.FormValue("repassword")
	fmt.Println(user.Name, ">>>>>>>>>>", password, ">>>>>>>>>>", repassword)
	salt := fmt.Sprintf("%06d", time.Now().Unix())

	data := models.FindUserByName(user.Name)
	if user.Name == "" || password == "" || repassword == "" {
		c.JSON(200, gin.H{
			"code":    9999, // 0000->success 9999->failed
			"message": "username & paaword can not be NULL!",
			"data":    user,
		})
		return
	}

	if data.Name != "" {
		c.JSON(200, gin.H{
			"code":    9999, // 0000->success 9999->failed
			"message": "Username already exists.",
			"data":    user,
		})
		return
	}

	if password != repassword {
		c.JSON(200, gin.H{
			"code":    9999,
			"message": "The passwords do not match.",
			"data":    user,
		})
		return
	}
	user.Password = utils.MakePassword(password, salt)
	user.Salt = salt
	user.Id = utils.GenerateId()
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"code":    0000,
		"message": "success",
		"data":    user,
	})
}
