package main

import (
	"example.com/api/server"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	myServer := gin.Default()
	server.RegisterRoutes(myServer)
	myServer.Run(":8080")
}
