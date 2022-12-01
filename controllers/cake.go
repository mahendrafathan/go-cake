package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mahendrafathan/go-cake/helpers"
	"github.com/mahendrafathan/go-cake/models"

	"github.com/gorilla/mux"
)

type CakeUsecase struct {
	cakeRepo CakeRepository
}

func NewCakeUsecase(cakeRepo CakeRepository) *CakeUsecase {
	return &CakeUsecase{
		cakeRepo: cakeRepo,
	}
}

type CakeResponseOutput struct {
	Data    interface{}
	Message string
}

func (c *CakeUsecase) CakeList(w http.ResponseWriter, r *http.Request) {
	cakes, err := c.cakeRepo.FindAll()
	if err != nil {
		helpers.RespondWithError(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, cakes)
}

func (c *CakeUsecase) CakeDetail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		helpers.RespondWithError(r.Context(), w, http.StatusBadRequest, err.Error())
		return
	}

	cake, err := c.cakeRepo.FindOne(id)
	if err != nil {
		helpers.RespondWithError(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, cake)
}

func (c *CakeUsecase) AddCake(w http.ResponseWriter, r *http.Request) {
	cake := models.Cake{}
	if err := json.NewDecoder(r.Body).Decode(&cake); err != nil {
		helpers.RespondWithError(r.Context(), w, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.cakeRepo.Create(cake); err != nil {
		helpers.RespondWithError(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, cake)
}

func (c *CakeUsecase) UpdateCake(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		helpers.RespondWithError(r.Context(), w, http.StatusBadRequest, err.Error())
		return
	}

	cake := models.Cake{}
	if err := json.NewDecoder(r.Body).Decode(&cake); err != nil {
		helpers.RespondWithError(r.Context(), w, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.cakeRepo.Update(id, cake); err != nil {
		helpers.RespondWithError(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, nil)
}

func (c *CakeUsecase) DeleteCake(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		helpers.RespondWithError(r.Context(), w, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.cakeRepo.Delete(id); err != nil {
		helpers.RespondWithError(r.Context(), w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, nil)
}
