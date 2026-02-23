package main

import (
	"fmt"
	"log"

	"github.com/CALEXCO/go-http-server/models"
	"github.com/gin-gonic/gin"
)

var DbProjects = models.NewDbProject()

var DbUsers = make([]models.User, 0)

var logger = log.Default()

func main() {

	// Setting logger
	// TODO: Set flags of logger Ldate, microseconds
	for i := 1; i <= 10; i++ {
		DbUsers = append(DbUsers, models.NewUser(fmt.Sprintf("User%d", i), fmt.Sprintf("User%d@gmail.com", i), fmt.Sprintf("Company%d", i), fmt.Sprintf("123145%d", i)))
	}

	// Crear 1 proyect de ejemplo
	p1 := models.NewProject("P1", "This is an example project", models.User{Name: "User123", Mail: "User123@gmail.com", Company: "Company1", DNI: "098766589"})
	for _, user := range DbUsers {
		p1.AddNewWorker(user)
	}

	//Start with 1 Project

	DbProjects.Projects["P1"] = &p1

	r := gin.Default()

	r.GET("/projects", DbProjects.GetAllProjects)
	r.POST("/projects", DbProjects.AddNewProject)
	r.DELETE("/prjects/:id", DbProjects.DeleteProject)
	logger.Println("Server listening on port :8080")
	r.Run(":8080")
}
