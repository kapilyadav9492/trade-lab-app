package controllers

import (
	"net/http"
	"log"
)

type Controller struct {
}

func (c *Controller) HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("App is up and running!")
	response := []byte("{\"Message\":\"success\"}")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	return
}