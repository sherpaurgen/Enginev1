package main

import (
	"Enginev1/models"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"time"
)

func (app *application) getMovie(w http.ResponseWriter,r *http.Request){
	id:= chi.URLParam(r,"movid")
	movid,err:=strconv.Atoi(id)
	if err != nil {
		app.logger.Println(err)
		app.errorJSON(w,err)
		return
	}
	movie := models.Movie{
		ID:movid,
		Title: "Mission impossible",
		Description: "some description",
		Year:2021,
		ReleaseDate: time.Date(2022,03,22,0,0,0,0,time.Local),
		Runtime: 100,
		Rating:4,
		MPAARating: "PG-13",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	app.writeJSON(w,http.StatusOK,movie,"moviewrap")
	app.logger.Println("the movie id is: ",movid)

}
func (app *application) getAllMovies(w http.ResponseWriter,r *http.Request){

}
