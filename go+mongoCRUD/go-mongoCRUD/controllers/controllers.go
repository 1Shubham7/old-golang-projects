package controllers

import (
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"github.com/1shubham7/go-mongo-crud/models"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// this is how we make a struct method
func (uc UserController) GetUser (w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id:= p.ByName("id")

	// WriteHeader sends an HTTP response header with the provided status code.WriteHeader sends an HTTP response header with the provided status code.
	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	//  u is a struct of type models.user
	u := models.User{}

	// c is collection of the users. also checking for err in the same line.
	err := uc.Session.DB("database-name").C("users").FindId(oid).One(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// marshalling and unmarshalling helps you work with json data
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	// telling tools like Postman, curl that we are sending json data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}
