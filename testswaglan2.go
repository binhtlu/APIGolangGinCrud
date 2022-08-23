package main

import (
	_ "connnnntrollers/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//export PATH=$(go env GOPATH)/bin:$PATH

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	BirthDate string `json:"birth"`
}

// message represents request response with a message
type message struct {
	Message string `json:"message"`
}

// todo slice to seed todo list data
var users = []User{
	{ID: "1", Name: " Vu Van An ", Age: 23, BirthDate: "10/9"},
	{ID: "2", Name: " Nguyen Van Tung ", Age: 23, BirthDate: "25/5"},
	{ID: "3", Name: "Dang Quoc Dai", Age: 23, BirthDate: "17/10"},
	{ID: "4", Name: " Ngo Ngoc Binh ", Age: 23, BirthDate: "9/4"},
	{ID: "5", Name: "Trinh Thi Thanh Hien", Age: 23, BirthDate: "29/11"},
}

// @Summary get all items in the User list
// @ID get-all-todos
// @Produce json
// @Success 200 {object} User
// @Router /api/users [get]
func getAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

// @Summary find a User item by ID
// @ID find-User-by-id
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} User
// @Failure 404 {object} message
// @Router /api/users/{id} [get]
func FindUsers(c *gin.Context) {
	ID := c.Param("id")

	// loop through todoList and return item with matching ID
	for _, todo := range users {
		if todo.ID == ID {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	// return error message if todo is not found
	r := message{"todo not found"}
	c.IndentedJSON(http.StatusNotFound, r)
}

// @Summary add a new item to the User list
// @ID create-User
// @Produce json
// @Param data body User true "User data"
// @Success 200 {object} User
// @Failure 400 {object} message
// @Router /api/users [post]
func createUser(c *gin.Context) {
	var newUser User

	// bind the received JSON data to newTodo
	if err := c.BindJSON(&newUser); err != nil {
		r := message{"an error occurred while creating todo"}
		c.IndentedJSON(http.StatusBadRequest, r)
		return
	}

	// add the new todo item to todoList
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

// @Summary delete a User item by ID
// @ID delete-User-by-id
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} User
// @Failure 404 {object} message
// @Router /api/users/{id} [delete]
func deleteUser(c *gin.Context) {
	id := c.Param("id")

	// loop through todoList and delete item with matching ID
	for index, todo := range users {
		if todo.ID == id {
			users = append(users[:index], users[index+1:]...)

		}
	}

	// return error message if todo is not found
	r := message{"todo not found"}
	c.IndentedJSON(http.StatusNotFound, r)
}

// @Summary update a User item by ID
// @ID update-User-by-id
// @Produce json
// @Param id path string true "User ID"
// @description update an users
// @
// @Success 200 {object} User
// @Failure 404 {object} message
// @Router /api/users/{id} [PUT]
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
			return
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User updated"})
}

func Handles() {
	router := gin.Default()

	router.GET("/api/users", getAllUsers)

	router.POST("/api/users", createUser)

	router.GET("/api/users/:id", FindUsers)

	router.PUT("/api/users/:id", updateUser)

	router.DELETE("/api/users/:id", deleteUser)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("localhost:8080")
}

// @title Go + Gin Todo API
// @version 1.0
// @description This is a sample server todo server. You can visit the GitHub repository at https://github.com/LordGhostX/swag-gin-demo

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {
	Handles()
}
