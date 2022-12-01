package routes

import (
	"net/http"

	"github.com/mahendrafathan/go-cake/controllers"
	"github.com/mahendrafathan/go-cake/repositories"

	"github.com/mahendrafathan/go-cake/common/db"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	db, err := db.NewMysqlConnection()
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	router.Use(LoggingMiddleware())

	s := router.PathPrefix("/cakes").Subrouter()
	{
		cakeRepo := repositories.NewCakeRepo(db)
		cakeUsecase := controllers.NewCakeUsecase(&cakeRepo)
		s.HandleFunc("", cakeUsecase.CakeList).Methods(http.MethodGet)
		s.HandleFunc("/{id}", cakeUsecase.CakeDetail).Methods(http.MethodGet)
		s.HandleFunc("", cakeUsecase.AddCake).Methods(http.MethodPost)
		s.HandleFunc("/{id}", cakeUsecase.UpdateCake).Methods(http.MethodPatch)
		s.HandleFunc("/{id}", cakeUsecase.DeleteCake).Methods(http.MethodDelete)
	}

	return router
}
