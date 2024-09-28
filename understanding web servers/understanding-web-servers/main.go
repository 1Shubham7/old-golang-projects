package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/1shubham7/understanding-web-servers/handlers"
	// "io"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	homehandler := handlers.NewHome(l) //you can just say it hh
	hellohandler := handlers.NewHello(l)
	byehandler := handlers.NewBye(l)

	servemux := http.NewServeMux() //you can also call this sm
	servemux.Handle("/", homehandler)
	servemux.Handle("/hello",hellohandler)
	servemux.Handle("/bye", byehandler)

	server := &http.Server{
		Addr: ":6000",
		Handler: servemux,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 *time.Second,
		WriteTimeout: 1 *time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	} ()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Performing Graceful shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
	// http.ListenAndServe(":6000", servemux)
	// second parameter is for http handler. if we say "nil" in second parameter, 
	// the server will take the default http handler, here we specified the http handler



	
}

