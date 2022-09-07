package main

import (
	_ "connnnntrollers/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// message represents request response with a message
type message struct {
	Message string `json:"message"`
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	BirthDate string `json:"birth"`
}

var userList = []User{
	{ID: "1", Name: " Vu Van An ", Age: 23, BirthDate: "10/9"},
	{ID: "2", Name: " Nguyen Van Tung ", Age: 23, BirthDate: "25/5"},
	{ID: "3", Name: "Dang Quoc Dai", Age: 23, BirthDate: "17/10"},
	{ID: "4", Name: " Ngo Ngoc Binh ", Age: 23, BirthDate: "9/4"},
	{ID: "5", Name: "Trinh Thi Thanh Hien", Age: 23, BirthDate: "29/11"},
}

// @Summary get all items in the todo list
// @ID get-all-Users
// @Produce json
// @Success 200 {object} User
// @Router /User [get]
func getAllTodos(c *gin.Context) {
	c.JSON(http.StatusOK, userList)
}

// @Summary get a User item by ID
// @ID get-User-by-id
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} User
// @Failure 404 {object} message
// @Router /User/{id} [get]
func getTodoByID(c *gin.Context) {
	ID := c.Param("id")

	// loop through todoList and return item with matching ID
	for _, todo := range userList {
		if todo.ID == ID {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	// return error message if todo is not found
	r := message{"todo not found"}
	c.JSON(http.StatusNotFound, r)
}

// @Summary add a new item to the User list
// @ID create-User
// @Produce json
// @Param data body User true "User data"
// @Success 200 {object} User
// @Failure 400 {object} message
// @Router /User [post]
func AddNewUser(c *gin.Context) {
	var NewUser User

	// bind the received JSON data to newTodo
	if err := c.BindJSON(&NewUser); err != nil {
		r := message{"an error occurred while creating todo"}
		c.JSON(http.StatusBadRequest, r)
		return
	}

	// add the new todo item to todoList
	userList = append(userList, NewUser)
	c.JSON(http.StatusCreated, NewUser)
}

// @Summary delete a User item by ID
// @ID delete-User-by-id
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} User
// @Failure 404 {object} message
// @Router /User/{id} [delete]
func deleteTodo(c *gin.Context) {
	ID := c.Param("id")

	// loop through todoList and delete item with matching ID
	for index, todo := range userList {
		if todo.ID == ID {
			userList = append(userList[:index], userList[index+1:]...)
			r := message{"successfully deleted todo"}
			c.JSON(http.StatusOK, r)
			return
		}
	}

	// return error message if todo is not found
	r := message{"todo not found"}
	c.JSON(http.StatusNotFound, r)
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
	// configure the Gin server
	router := gin.Default()
	router.GET("/User", getAllTodos)
	router.GET("/User/:id", getTodoByID)
	router.POST("/User", AddNewUser)
	router.DELETE("/User/:id", deleteTodo)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// run the Gin server
	router.Run("localhost:8080")
}
