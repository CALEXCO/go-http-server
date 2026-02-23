package models

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DbProjects struct {
	NumberOfProjects int                 `json:"number_of_projects"`
	Projects         map[string]*Project `json:"projects"`
}

func NewDbProject() *DbProjects {
	return &DbProjects{
		NumberOfProjects: 0,
		Projects:         make(map[string]*Project),
	}
}

func (db *DbProjects) projectExists(p Project) bool {
	for _, project := range db.Projects {
		if project.Name == p.Name {
			return true
		}
	}
	return false
}

func (db *DbProjects) AddNewProject(c *gin.Context) {

	var inputProject Project
	if err := c.ShouldBindJSON(&inputProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	p := Project{
		Name:        inputProject.Name,
		Description: inputProject.Description,
		Leader:      inputProject.Leader,
		Workers:     inputProject.Workers,
	}

	if db.projectExists(p) {
		c.JSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("Project %s already exists", p.Name)})
		return
	}

	db.Projects[p.Name] = &p
	db.NumberOfProjects++
	c.JSON(http.StatusAccepted, gin.H{"data": p})

}

func (db *DbProjects) GetAllProjects(c *gin.Context) {
	c.JSON(http.StatusOK, db.Projects)
}

func (db *DbProjects) DeleteProject(c *gin.Context) {

	var inputProject Project
	if err := c.ShouldBindJSON(&inputProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	p := Project{
		Name: inputProject.Name,
	}
	if !db.projectExists(p) {
		c.JSON(http.StatusNotFound, p.Name)
		return
	}

	delete(db.Projects, p.Name)
	db.NumberOfProjects--
	c.JSON(http.StatusOK, p.Name)

}
