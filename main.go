package main

import (
	"OnlineVideo/router"
	"OnlineVideo/utils"
)

func main() {
	utils.InitMySQL()
	r := router.Router()
	r.Run(":8080")
}
