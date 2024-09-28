package handlers

import (
	"log"
	"net/http"
	"io/ioutil"
	"fmt"
)

type Hello struct{
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h*Hello) ServeHTTP (rw http.ResponseWriter, r *http.Request){
	h.l.Println("Hello Guys")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Sorry, Please try again", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(rw, "This is the hello page. Hello %s", d)
}