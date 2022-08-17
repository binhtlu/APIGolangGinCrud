package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	BirthDate string `json:"birth"`
}

var users = []User{
	{ID: "1", Name: " Vu Van An ", Age: 23, BirthDate: "10/9"},
	{ID: "2", Name: " Nguyen Van Tung ", Age: 23, BirthDate: "25/5"},
	{ID: "3", Name: "Dang Quoc Dai", Age: 23, BirthDate: "17/10"},
	{ID: "4", Name: " Ngo Ngoc Binh ", Age: 23, BirthDate: "9/4"},
	{ID: "5", Name: "Trinh Thi Thanh Hien", Age: 23, BirthDate: "29/11"},
}

// getUser User
// @Summary  print users
// @Description print a list of users
// @Tags users
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /api/users [GET]

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

// CreateUser User
// @Summary Create a user
// @Description Create a new user item
// @Tags users
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /api/users [POST]
func CreateUsers(c *gin.Context) {
	var NewUsers User
	if err := c.BindJSON(&NewUsers); err != nil {
		return
	}
	users = append(users, NewUsers)
	c.IndentedJSON(http.StatusCreated, NewUsers)
}

// FindUser User
// @Summary Find a user
// @Description fiind a new user item
// @Tags users
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /api/users/:id [GET]
func FindUsers(c *gin.Context) {
	id := c.Param("id")

	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

// updateUser User
// @Summary update a user
// @Description update a new user item
// @Tags users
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /api/users/:id [PUT]
func updateUser(c *gin.Context) {
	id := c.Param("id")
	var NewUser User

	if err := c.BindJSON(&NewUser); err != nil {
		return
	}
	for index, b := range users {
		if b.ID == id {
			if NewUser.Name != b.Name && NewUser.Name != "" {
				users[index].Name = NewUser.Name
			}
			if NewUser.Age != b.Age && NewUser.Age != 0 {
				users[index].Age = NewUser.Age
			}
			if NewUser.BirthDate != b.BirthDate && NewUser.BirthDate != "" {
				users[index].BirthDate = NewUser.BirthDate
			}
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User updated"})
}

// deleteUser User
// @Summary delete a user
// @Description update a new user item
// @Tags users
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Router /api/users/:id [DELETE]
func deleteUser(c *gin.Context) {
	id := c.Param("id")

	for index, user := range users {
		if user.ID == id {
			users = append(users[:index], users[index+1:]...)
		}
	}

	for _, user := range users {
		fmt.Println(user.ID, user.Name, user.Age, user.BirthDate)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "success"})
}

func Handles() {
	router := gin.Default()

	router.GET("/api/users", getUsers)

	router.POST("/api/users", CreateUsers)

	router.GET("/api/users/:id", FindUsers)

	router.PUT("/api/users/:id", updateUser)

	router.DELETE("/api/users/:id", deleteUser)

	router.Run("localhost:8080")
}

// @title GO rest API
// @version 1.0
// @description go rest api + gin + swagger
// @host localhost:8080
// BasePath /
func main() {
	Handles()
}
