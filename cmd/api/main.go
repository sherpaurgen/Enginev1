package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version="1.0.0"
type config struct {
	port int
	env string
}
type AppStatus struct {
	Status string `json:"status"`
	Enviroment string `json:"enviroment"`
	Version string `json:"version"`
}
type application struct {
	config config
	logger *log.Logger
}

func main(){
var cfg config
logger:=log.New(os.Stdout,"",log.Ldate|log.Ltime)
app:=&application{
	config:cfg,
	logger:logger,
}

flag.IntVar(&cfg.port,"port",4000,"Server port to listen on/default port 4000")
flag.StringVar(&cfg.env,"env","dev","Specify dev or prod enviroment")
flag.Parse()



srv:= &http.Server{
	Addr: fmt.Sprintf(":%d",cfg.port),
	Handler: app.routes(),
	IdleTimeout: time.Minute,
	ReadTimeout: 10*time.Second,
	WriteTimeout: 30*time.Second,
}
logger.Println("Starting server on port ",cfg.port)
err:=srv.ListenAndServe()
if err!=nil{
	log.Fatal(err)
}

}
