package main

import (
	"backend/models"
	"backend/storage"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {

	router := gin.Default()

	router.GET("/user/all", GetUsers)

	router.POST("/user/create", CreateUser)

	router.GET("/user/:id", GetIdByPath)

	router.PUT("/user/update/:id", UpdateUser)

	router.DELETE("/user/delete/:id", DeleteUser)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Println(err)
	}

}

func CreateUser(c *gin.Context) {
	var reqUser models.User
	err := c.BindJSON(&reqUser)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	id := uuid.NewString()
	reqUser.ID = id

	respUser, err := storage.CreateUser(reqUser)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusCreated, respUser)
}

func GetUsers(c *gin.Context) {
	page := c.Query("page")
	intPage, err := strconv.Atoi(page)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	limit := c.Query("limit")
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	respUsers, err := storage.GetAllUsers(intPage, intLimit)

	c.IndentedJSON(http.StatusFound, respUsers)
}

func UpdateUser(c *gin.Context) {
	userId := c.Param("id")

	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	updatedUser, err := storage.UpdateUserById(userId, user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusAccepted, updatedUser)
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("id")

	respUser, err := storage.DeleteUserById(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusFound, respUser)
}


func GetIdByPath(c *gin.Context) {
	userId := c.Param("id")

	respUser, err := storage.GetUserById(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusFound, respUser)
}
