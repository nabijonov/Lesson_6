package main

import (
	"fmt"

	"github.com/Shemistan/Lesson_6/internal/api"
	"github.com/Shemistan/Lesson_6/internal/models"
	"github.com/Shemistan/Lesson_6/internal/service"
	"github.com/Shemistan/Lesson_6/internal/storage"
)

func main() {

	server1()
	
}

func server1(){
	db := storage.NewIStorage()
	service := service.NewIService(db)
	api := api.NewIApi(service)

	auth := models.SAuth{
		Login: "nabijonov",
		PasswordHash: "123",
	}
	user1 := models.SUser{Login: "nabijonov",PasswordHash: "123",Name: "anvar",Surname: "nabijonov",Status: "active", Role: "active"}
	user2 := models.SUser{Login: "qwerty",PasswordHash: "123",Name: "anvar",Surname: "nabijonov",Status: "active", Role: "active"}
	user3 := models.SUser{Login: "qwerty123",PasswordHash: "123",Name: "anvar",Surname: "nabijonov",Status: "active", Role: "active"}
	
	_, _ = db.Add(&user1)
	_, _ = db.Add(&user2)
	_,_ =service.Add(&user3)
	id4, _ := service.Auth(&auth)
	_, _ = service.GetUser(id4)
		
	// API

	reqAuth := `{"Login":"qwerty","PasswordHash":"123"}`   //json to service_struct
	apigetIDJSON,_ := api.Auth(reqAuth)
	reqAdd := `{"Login":"testAddFromAPI","PasswordHash":"123","Name":"anvar","Surname":"nabijonov","Status":"active","Role":"active"}`
	boo1, _ := api.Add(reqAdd)
	_ = boo1
	// fmt.Println(apigetIDJSON, "POLZOVOTEL")
	reqGetUser := api.GetUser(apigetIDJSON)
	fmt.Println("getting User From API", reqGetUser)
	err := api.DeleteUser(apigetIDJSON)
	if err == nil {
		api.GetUsers()
	}
	fmt.Println()
	api.GetStatistics()
	
}

