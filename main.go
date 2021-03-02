package main

import (
	"Final/config"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/gorilla/mux"
)

func main() {

	router := gin.Default()

	config.ConnectDatabase()

	//route

	router.Run(":3000")

}
