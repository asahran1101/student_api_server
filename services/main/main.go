package main

import (
	"github.com/asahran1101/student_api_server/services/server"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	myServer := gin.Default()
	server.RegisterRoutes(myServer)
	myServer.Run(":8080")
}
