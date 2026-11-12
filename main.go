package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	RespCode    string `json:"respCode"`
	RespMessage string `json:"respMessage"`
	RespData    any    `json:"respData"`
}

var users = []User{
	{ID: 1, Name: "John"},
	{ID: 2, Name: "Doe"},
}

// List users
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		RespCode:    "00",
		RespMessage: "success",
		RespData:    users,
	})
}

// Get user by ID
func GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			RespCode:    "01",
			RespMessage: "invalid user id",
			RespData:    nil,
		})
		return
	}

	for _, user := range users {
		if user.ID == id {
			c.JSON(http.StatusOK, Response{
				RespCode:    "00",
				RespMessage: "success",
				RespData:    user,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, Response{
		RespCode:    "02",
		RespMessage: "user not found",
		RespData:    nil,
	})
}

// Create user
func CreateUser(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			RespCode:    "01",
			RespMessage: err.Error(),
			RespData:    nil,
		})
		return
	}

	users = append(users, user)

	c.JSON(http.StatusCreated, Response{
		RespCode:    "00",
		RespMessage: "user created successfully",
		RespData:    user,
	})
}

// Update user
func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			RespCode:    "01",
			RespMessage: "invalid user id",
			RespData:    nil,
		})
		return
	}

	var updateUser User
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			RespCode:    "01",
			RespMessage: err.Error(),
			RespData:    nil,
		})
		return
	}

	for i := range users {
		if users[i].ID == id {
			users[i].Name = updateUser.Name

			c.JSON(http.StatusOK, Response{
				RespCode:    "00",
				RespMessage: "user updated successfully",
				RespData:    users[i],
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, Response{
		RespCode:    "02",
		RespMessage: "user not found",
		RespData:    nil,
	})
}

// Delete user
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			RespCode:    "01",
			RespMessage: "invalid user id",
			RespData:    nil,
		})
		return
	}

	for i := range users {
		if users[i].ID == id {
			deletedUser := users[i]
			users = append(users[:i], users[i+1:]...)

			c.JSON(http.StatusOK, Response{
				RespCode:    "00",
				RespMessage: "user deleted successfully",
				RespData:    deletedUser,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, Response{
		RespCode:    "02",
		RespMessage: "user not found",
		RespData:    nil,
	})
}

func main() {
	r := gin.Default()

	r.GET("/api/v1/users", GetUsers)
	r.GET("/api/v1/users/:id", GetUserByID)
	r.POST("/api/v1/users", CreateUser)
	r.PUT("/api/v1/users/:id", UpdateUser)
	r.DELETE("/api/v1/users/:id", DeleteUser)
	fmt.Println("RUNNING")
	r.Run(":3000")
}