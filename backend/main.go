package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Image     string `json:"url"`
	CreatedAt string `json:"createdAt"`
}

func getUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	userJSON, _ := json.Marshal(users[id])
	c.String(http.StatusOK, string(userJSON))
}

func addUser(c *gin.Context) {
	user, _ := ioutil.ReadAll(c.Request.Body)
	userJSON, _ := json.Marshal(user)
	c.String(http.StatusOK, string(userJSON))
}

func removeUser(c *gin.Context) {
	c.String(http.StatusOK, string("fds"))
}

func patchUser(c *gin.Context) {
	c.String(http.StatusOK, string("fds"))
}

func main() {
	router := gin.Default()

	router.GET("/api/user/:id", getUser)
	router.DELETE("/api/user/:id", removeUser)
	router.PATCH("/api/user/:id", patchUser)
	router.POST("/api/user", addUser)

	router.Run(":3000")
}
