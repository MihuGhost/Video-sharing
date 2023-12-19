package service

import (
	"OnlineVideo/models"
	"OnlineVideo/utils"
	"fmt"
	"github.com/asaskevich/govalidator"
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

func FindUserByNameAndPwd(c *gin.Context) {
	data := models.User{}
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    9999,
			"message": "User not exists",
			"data":    data,
		})
		return
	}

	flag := utils.ValidPassword(password, user.Salt, user.Password)
	if !flag {
		c.JSON(200, gin.H{
			"code":    9999,
			"message": "Password is not correct",
			"data":    data,
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    0000,
		"message": "Success",
		// todo : set for data
		"data": "Secret",
	})
}

func DeleteUser(c *gin.Context) {
	user := models.User{}
	id := c.Query("id")
	user.Id = id
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"code":    0000,
		"message": "Delete successfully",
		// todo : set for data
		"data": "Delete successfully",
	})
}

func UpdateUser(c *gin.Context) {
	user := models.User{}
	id := c.PostForm("id")
	user.Id = id
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")

	_, err := govalidator.ValidateStruct(user)

	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"code":    9999,
			"message": "do not match",
			"data":    user,
		})
	} else {
		models.UpdateUser(user)
		c.JSON(200, gin.H{
			"code":    0000,
			"message": "Update successfully",
			// todo : set for data
			"data": "Update user successfully",
		})
	}
}
