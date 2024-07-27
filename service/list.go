package service

import (
	"log"
	"regexp"

	"com.nokia.zayar/repository"
)

func ListItems(verbose bool, titleRegex, directorRegex, actorRegex string, asc, desc bool) {
	regexpTitle, err := regexp.Compile(titleRegex)
	if err!= nil {
        log.Printf("Invalid title regex: %v\n", err)
        return
    }
	regexpDirector, err := regexp.Compile(directorRegex)
	if err!= nil {
        log.Printf("Invalid director regex: %v\n", err)
        return
    }
	regexpActor, err := regexp.Compile(actorRegex)
	if err!= nil {
        log.Printf("Invalid actor regex: %v\n", err)
        return
    }

	var movies []repository.Movie = movieRepository.GetAllMovies(asc, desc)

	var none bool = false
	for _, movie := range movies {

		if len(titleRegex) != 0 && 
						!regexpTitle.MatchString(movie.Title) { continue }

		if len(directorRegex) != 0 && 
						!regexpDirector.MatchString(personRepository.GetPersonName(movie.DirectorID)) { continue }

		var include bool = false
		if len(actorRegex) == 0 {
			include = true
		} else {
			for _, actorID := range movie.Actors {
                if regexpActor.MatchString(personRepository.GetPersonName(actorID)) {
                    include = true
                    break
                }
            }
		}

		if include { movieRepository.PrintMovie(movie, verbose); none = true  }

	}

	if !none { log.Println("There is no matching movies in the database.")}
}

