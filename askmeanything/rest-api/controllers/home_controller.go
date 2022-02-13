package controllers

import "net/http"

func Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ask Me Anything Rest API!"))
	}
}
