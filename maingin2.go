package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// tạo một cấu trúc để định nghĩa đối tượng
type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	BirthDate string `json:"birth"`
}

// tạo một map dữ liệu cho cấu trúc đối tượng tương ứng
var users = []User{
	{ID: "1", Name: " Vu Van An ", Age: 23, BirthDate: "10/9"},
	{ID: "2", Name: " Nguyen Van Tung ", Age: 23, BirthDate: "25/5"},
	{ID: "3", Name: "Dang Quoc Dai", Age: 23, BirthDate: "17/10"},
	{ID: "3", Name: " Ngo Ngoc Binh ", Age: 23, BirthDate: "9/4"},
	{ID: "4", Name: "Trinh Thi Thanh Hien", Age: 23, BirthDate: "29/11"},
}

// chuyển các dữ liệu thành dạng json
/*Gin.Context là phần quan trọng nhất của Gin. Nó mang chi tiết yêu cầu,
xác nhận và tuần tự hóa JSON, và nhiều hơn nữa.*/
//Gọi Context.Indentedjson để chuyển struct thành JSON và thêm nó vào (response) phản hồi.
// http.StatusOK gửi tráng thái cho client 200Ok
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

// PostUsers thêm một user từ JSON nhận được trong request(yêu cầu) body
func postUsers(c *gin.Context) {
	var NewUsers User
	//Gọi Bindjson để ràng buộc JSON đã nhận được tới NewUsers
	if err := c.BindJSON(&NewUsers); err != nil {
		return
	}
	// thêm một user mới tới Slice
	users = append(users, NewUsers)
	c.IndentedJSON(http.StatusCreated, NewUsers)
	//Sử dụng Context.BindJSON để rằng buộc request body tới NewUsers
	//Nối(append) cấu trúc user được khởitạo(initialized) từ JSON vào Users slice.
	//Thêm mã trạng thái 201 vào phản hồi(response), cùngvới(along) JSON đạidiện(repesenting) cho album bạn đã thêm.
}

// trả một phần tử nhất định nào đó

func getUsersID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

// update new User

func updateUser(c *gin.Context) {
	id := c.Param("id")
	var NewUser User

	// get body trong NewUser
	if err := c.BindJSON(&NewUser); err != nil {
		return
	}
	// tạo vòng lặp trong the slice và update
	for index, b := range users {
		if b.ID == id {
			//{ID: "1", Name: " Vu Van An ", Age: 23, BirthDate: "10/9"},
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

// xoa user
func deleteUser(c *gin.Context) {
	id := c.Param("id")

	// iterate and delete the element
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

// chứa các phương thức response
func Handles() {
	// khởi tại(Intialize) Gin router sử dụng Default
	router := gin.Default()

	//Sử dụng  GET function để liên kết GET HTTP method và users với chức năng xử lý.
	/*Lưu ý rằng bạn đang chuyển tên của hàm getalbums. Điều này khác với việc chuyển kết quả của hàm,
	mà bạn sẽ làm bằng cách vượt qua getalbums () (lưu ý dấu ngoặc đơn).*/
	router.GET("/api/users", getUsers)
	// PostUsers
	router.POST("/api/users", postUsers)
	// timf một phần tử xác định
	router.GET("/api/users/id", getUsersID)
	// sua
	router.PUT("/api/users/id", updateUser)
	// xoa
	router.DELETE("/api/users/id", deleteUser)
	//Sử dụng Run function để gắn Router vào http.server và khởi động server trên cổng 8080.
	router.Run("localhost:8080")
}

func main() {
	Handles()
}
