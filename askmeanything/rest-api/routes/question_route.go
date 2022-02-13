package routes

import (
	"github.com/gorilla/mux"
	"github.com/skatsaounis/askmeanything/controllers"
)

func QuestionRoute(router *mux.Router) {
	router.HandleFunc("/question", controllers.CreateQuestion()).Methods("POST")
	router.HandleFunc("/question/{questionId}", controllers.GetAQuestion()).Methods("GET")
	router.HandleFunc("/question/{questionId}", controllers.EditAQuestion()).Methods("PUT")
	router.HandleFunc("/question/{questionId}", controllers.DeleteAQuestion()).Methods("DELETE")
	router.HandleFunc("/questions", controllers.GetAllQuestions()).Methods("GET")
}
