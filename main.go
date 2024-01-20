package main

import (
	"example.com/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.DELETE("/students/:rollNo", routes.DeleteStudent)

	server.GET("/students", routes.GetStudents)
	server.GET("/students/:rollNo", routes.GetStudent)

	server.POST("/students", routes.CreateStudent)

	server.PUT("/students/:rollNo", routes.UpdateStudent)

	server.Run(":8080")
}
