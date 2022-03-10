package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
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
	Addr: fmt.Sprintf("%d",cfg.port),
	Handler: router,
}
}
