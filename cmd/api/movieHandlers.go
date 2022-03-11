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
	movid,_:=strconv.Atoi(id)
	movie := models.Movie{
		ID:movid,
		Title: "mission impossible",
		Description: "some description",
		Year:2021,
		ReleaseDate: time.Date(2022,03,22,0,0,0,0,time.Local),
		Runtime: 100,
		Rating:4,
		MPAARating: "PG-13",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := app.writeJSON(w,http.StatusOK,movie,"moviewrap")
	app.logger.Println("the movie id is: ",movid)
	if err != nil {
		app.logger.Println(err)
	}
}
func (app *application) getAllMovies(w http.ResponseWriter,r *http.Request){

}
