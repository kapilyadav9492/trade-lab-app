package main

import (
	"github.com/trade-lab-app/backend/golang/router"
	"net/http"
	"time"
	"github.com/trade-lab-app/backend/golang/common/conf"
	"github.com/trade-lab-app/backend/golang/controllers"
	"log"
)

func init() {
	go controllers.ClientWebSocket()
}

func main() {
	log.Println("------------ Server Started ------------",)

	router := router.NewRouter()

	router.Methods("GET", "POST", "DELETE", "UPDATE")

	serv := &http.Server{
		Handler: router,
		Addr:    conf.Cfg.Server.ServerAddress,

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println(serv.ListenAndServe())

	defer log.Println(" ----- Server Closed !!! Please restart server... ------- ")
}
