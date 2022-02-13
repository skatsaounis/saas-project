package routes

import (
	"github.com/gorilla/mux"
	"github.com/skatsaounis/askmeanything/controllers"
)

func HomeRoute(router *mux.Router) {
	router.HandleFunc("/", controllers.Home())
}
