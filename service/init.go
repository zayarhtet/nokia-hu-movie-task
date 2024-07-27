package service

import (

    "com.nokia.zayar/repository"
)

var movieRepository repository.MovieRepository 
var personRepository repository.PersonRepository

func SetMovieRepository(repo repository.MovieRepository) {
	movieRepository = repo
}

func SetPersonRepository(repo repository.PersonRepository) {
	personRepository = repo
}

func Init() {
	movieRepository = repository.NewMovieRepository()
	personRepository = repository.NewPersonRepository()
}