package routes

import (
	"encoding/json"
	"net/http"
	"github.com/Michle99/goRestApi/controllers"
	"github.com/Michle99/goRestApi/models"
	"github.com/gorilla/mux"
)

func MovieRoutes() *mux.Router {
	var router = mux.NewRouter()
	router = mux.NewRouter().StrictSlash(true)

	//Home Toute
	router.HandleFunc("/api/",func(rw http.ResponseWriter, r *http.Request) {
		 message := models.Message{
			 Message: "Movie API!!!!",
		 }
		json.NewEncoder(rw).Encode(message)
	})


	//Other Routes
	router.HandleFunc("/api/movies",controllers.AddMovie).Methods(http.MethodPost)
	router.HandleFunc("/api/movies",controllers.GetAllMovies).Methods(http.MethodGet)
	router.HandleFunc("/api/movies/{id}",controllers.GetMovieById).Methods(http.MethodGet)
	router.HandleFunc("/api/movies/{id}",controllers.DeleteMovieById).Methods(http.MethodDelete)
	router.HandleFunc("/api/movies/{id}",controllers.UpdateMovie).Methods(http.MethodPut)
	
	return router
}