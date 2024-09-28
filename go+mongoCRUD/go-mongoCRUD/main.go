package main

import (
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"
	"github.com/1shubham7/go-mongo-crud/controllers"
)

func main() {

	router := httprouter.New()
	uc := controllers.NewUserController(getSession())
	// we will create a NewUserController with parameter that is the output (session) of getSession()

	router.GET("/users", uc.GetAllUsers)
	router.GET("/user/:id", uc.GetUserById)

	router.DELETE("delete/user/:id", uc.DeleteUserById)

	router.POST("add/user/:id", uc.AddUserById)

	router.PUT("update/user/:id", uc.UpdateUserBuId)
	
	http.ListenAndServe(":8080", router)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("moongodb://localhost:27107")
	// this is basically a mongo db session

	if err!= nil{
		panic(err)
	}
	return s
}
