package controllers

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

//ExamenController ...
type ExamenController struct{}

var names = []string{
	"Triple H", "La Roca", "Stone Cold", "Edge", "Jhon Cena", "Undertaker", "Andre", "Riquichi", "Kane", "Big Show",
}

type nameRequest struct {
	Name string `json:"name" binding:"required"`
}

//GetRandomName ...
func (ctrl ExamenController) GetRandomName(c *gin.Context) {

	ran := rand.Intn(10)

	c.JSON(http.StatusOK, gin.H{"name": names[ran]})
}

//GetStatus ...
func (ctrl ExamenController) GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Listos para el examen"})
}

//PostSameName ...
func (ctrl ExamenController) PostSameName(c *gin.Context) {

	var name nameRequest

	if c.ShouldBindJSON(&name) != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Name was not provided"})
		return
	}

	if !doesNameExist(name.Name) {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Name is not valid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"name": name.Name})
}

//PutName ...
func (ctrl ExamenController) PutName(c *gin.Context) {

	var name nameRequest

	if c.ShouldBindJSON(&name) != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Name was not provided"})
		return
	}

	if !doesNameExist(name.Name) {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Name is not valid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"name": fmt.Sprintf("Hola %s", name.Name)})
}

func doesNameExist(name string) (exist bool) {
	for _, n := range names {
		if n == name {
			exist = true
			continue
		}
	}
	return
}
