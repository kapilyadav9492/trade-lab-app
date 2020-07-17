package router

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/trade-lab-app/backend/golang/controllers"
)

var controller = &controllers.Controller{}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Role        []string
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{"Health Check", "GET", "/healthcheck", controller.HealthCheck, []string{"ADMIN"}},

	Route{"Currency", "GET", "/currency", controller.Currency, []string{"ADMIN"}},

	Route{"Get All Currency", "GET", "/get/all/currency", controller.GetAllCurrency, []string{"ADMIN"}},
}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(CommonMiddleWare)

	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
func CommonMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-type", "application/json")
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(res, req)
	})
}
