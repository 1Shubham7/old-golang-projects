package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// we are creating a handler  and we do that by creating a struct that implements the interface http handler

// let's create a struct called Home
type Home struct {
	l *log.Logger
}

func NewHome(l *log.Logger) *Home {
	return &Home{l}
}

func (h*Home) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	h.l.Println("Hello World")
	// instead of log.Println we can use h.l.Println its the same
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Sorry, Please try again!", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "This is the Home page, data passed is %s", d)
}