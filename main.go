package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/passwords", getPasswords)
	router.POST("/passwords", postPassword)

	router.Run("localhost:5000")
}

type password struct {
	Id       string `json:"_id"`
	Service  string `json:"service"`
	Password string `json:"password"`
}

var passwords = []password{
	{Id: "2c10963a-337d-4b2c-ab3e-19526c579c87", Service: "gitlab", Password: "12345678"},
	{Id: "c0c4fb59-1d09-445f-b438-8346aea51955", Service: "github", Password: "abcdefgh"},
	{Id: "f27ac592-b825-4252-83ce-3bd0fcf7b8d3", Service: "steam", Password: "!@#$%^&*"},
}

func getPasswords(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, passwords)
}

func postPassword(c *gin.Context) {
	var newPassword password

	if error := c.BindJSON(&newPassword); error != nil {
		return
	}

	passwords = append(passwords, newPassword)
	c.IndentedJSON(http.StatusCreated, newPassword)
}
