package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

//--- các phương thức thêm, sửa, xóa, tìm kiếm

func (user User) ToString() string {
	return fmt.Sprintf("id: %s\nName: %s\nPassword: %s\n", user.Id, user.Name, user.Password)
}

var (
	listUser = make([]User, 0)
)

func CreateUser(user User) bool {
	if user.Id != "" && user.Name != "" && user.Password != "" {
		if userF, _ := FindUser(user.Id); userF == nil {
			listUser = append(listUser, user)
			return true
		}
	}
	return false
}

func FindUser(id string) (User, error) {
	for _, user := range listUser {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func UpdateUser(eUser User) bool {
	for index, user := range listUser {
		if user.Id == eUser.Id {
			listUser[index] = eUser
			return true
		}
	}
	return false
}

func DeleteUser(id string) bool {
	for index, user := range listUser {
		if user.Id == id {
			copy(listUser[index:], listUser[index+1:])
			listUser[len(listUser)-1] = User{}
			listUser = listUser[:len(listUser)-1]
			return true
		}
	}
	return false
}

func GetAllUser() []User {
	return listUser
}

// Xử lý requet, response
func FindUserapi(response http.ResponseWriter, request *http.Request) {
	ids, ok := request.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		responseWithError(response, http.StatusBadRequest, "url parm id is missing")
		return
	}
	user, err := FindUser(ids[0])
	if err != nil {
		responseWithError(response, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJSON(response, http.StatusOK, user)
}

func GetAll(response http.ResponseWriter, request *http.Request) {
	users := GetAllUser()
	responseWithJSON(response, http.StatusOK, users)
}

func CreateUserAPI(response http.ResponseWriter, request *http.Request) {
	var user User
	err := json.NewDecoder(request.Body).Decode(user)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		result := CreateUser(user)
		if !result {
			responseWithError(response, http.StatusBadRequest, "Could not create user")
			return
		}
		responseWithJSON(response, http.StatusOK, user)
	}
}

func UpdateUserAPI(response http.ResponseWriter, request *http.Request) {
	var user User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		result := UpdateUser(user)
		if !result {
			responseWithError(response, http.StatusBadRequest, "Could not update user")
			return
		}
		responseWithJSON(response, http.StatusOK, "Update user successfully")
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	ids, ok := request.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		responseWithError(response, http.StatusBadRequest, "Url Param id is missing")
		return
	}
	result := DeleteUser(ids[0])
	if !result {
		responseWithError(response, http.StatusBadRequest, "Could not delete user")
		return
	}
	responseWithJSON(response, http.StatusOK, "Delete user successfully")
}

func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
	responseWithJSON(response, statusCode, map[string]string{
		"error": msg,
	})
}

func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}

func main() {
	router := NewRouter()

}
