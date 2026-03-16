package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	Type "github.com/Spr1zze/barber-shop-backend/model"
)

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func PostUsers(c *gin.Context) {
	var newUser Type.User

	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

var users = []Type.User{
	{ID: "1", Name: "Camilla Frederiksen", Phone: 23406751},
	{ID: "2", Name: "Johan Henriksen", Phone: 57279182},
	{ID: "3", Name: "Peter Knudsen", Phone: 12749501},
	{ID: "4", Name: "Lone Benedict", Phone: 98347261},
}
