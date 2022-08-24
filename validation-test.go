package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Body struct {
	Product string `json:"product" binding:"required,alpha"`
	Price   uint   `json:"price" binding:"required,gte=10,lte=1000"`
}

var bodys = []Body{
	{Product: "sua", Price: 20},
	{Product: "ya", Price: 10},
}

func getAllBodys(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bodys)
}

//type User struct {
//	ID        string `json:"id" binding:"required,startswith=A,len=10"`
//	Name      string `json:"name" binding:"required, len=15"`
//	Age       int    `json:"age" binding:"required, gte=10,lte=120"`
//	BirthDate string `json:"birth" binding:"required,len=8"`
//}

//type User struct {
//	ID        string `json:"id"`
//	Name      string `json:"name"`
//	Age       int    `json:"age"`
//	BirthDate string `json:"birth"`
//}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
		//case "len":
		//	return "Should be greater than" + fe.Param()
	}
	return "Unknown error"
}

func addbody(context *gin.Context) {
	var Newbody Body
	if err := context.BindJSON(&Newbody); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}
	bodys = append(bodys, Newbody)
	context.JSON(http.StatusAccepted, &Newbody)
}

func main() {
	engine := gin.New()
	engine.GET("/test", getAllBodys)
	engine.POST("/test", addbody)
	engine.Run(":8080")
}
