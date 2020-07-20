package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)
type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

// It will return null; due to initialization
//var Users1 []User
var Users []User


func main() {
	tmp := make([]User, 5)
	for i := range tmp {
		tmp[i].ID = uuid.New().String()
		tmp[i].Name = "name"
		tmp[i].Age = i
		tmp[i].Email = "name@gmail.com"
		Users = append(Users, tmp[i])
	}
	
	
	r := gin.Default()
	
	users := r.Group("/users" )
	{
		users.GET("", GetUsers)
		users.POST("", CreateUser)
		users.PUT("/:id", EditUser)
		users.DELETE("/:id", DeleteUser)
	}
	
	if err := r.Run(":3000"); err != nil {
		log.Fatal(err.Error())
	}
}

func getCountries() []string {
	countries := []string{"United states", "United kingdom",   "Austrilia", "India", "China", "Russia", "France", "Germany", "Spain"} // can be much more
	return countries
}

func GetUsers(ctx *gin.Context)  {
	ctx.JSON(http.StatusAccepted, Users)
}
func CreateUser(ctx *gin.Context)  {
	var reqBody User
	if err  := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Bad request",
		})
		return
	}
	reqBody.ID = uuid.New().String()
	Users = append(Users, reqBody)
	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "User's added successfully",
	})
}
func EditUser(ctx *gin.Context)   {
	id := ctx.Param("id")
	var reqBody User
	if err  := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Bad request",
		})
		return
	}
	for i, u := range Users {
		if u.ID == id {
			Users[i].Name = reqBody.Name
			Users[i].Age = reqBody.Age
			Users[i].Email = reqBody.Email
			
			ctx.JSON(http.StatusOK, gin.H{
				"error":   false,
				"message": "User's edit successfully",
			})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{
		"error":   true,
		"message": "User's Not Found",
	})
}
func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var reqBody User
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Bad request",
		})
		return
	}
	for i, u := range Users {
		if u.ID == id {
			Users = append(Users[:i], Users[i+1:]...)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"error":   false,
			"message": "User's deleted successfully",
		})
		return
	}
	ctx.JSON(http.StatusNotFound, gin.H{
		"error":   true,
		"message": "User's Not Found",
	})
}
