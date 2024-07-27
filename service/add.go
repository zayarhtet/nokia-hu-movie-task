package service

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"com.nokia.zayar/command"
	"com.nokia.zayar/repository"
)


func AddItem(pFlag, mFlag bool) {
	if (pFlag) {
		addPerson()
	} else if (mFlag) {
		addMovie()
	}
}

func addPerson() {
	var name string
	var birthYear int
	fmt.Print("Name: ")
	name, _ = command.GetInputString()
	for {
		var works bool
		fmt.Print("Birth Year: ")
		birthYear, works = command.GetInputInteger()
		if !works {
			fmt.Println("Invalid birth year, must be a positive integer.")
		} else {
			break
		}
	}

	newID := personRepository.GetPersonLength() + 1
	newPerson := repository.Person{
		ID:        newID,
		Name:      name,
		BirthYear: birthYear,
	}
	personRepository.AddPerson(&newPerson)
	fmt.Println("Person added successfully!")
}

func addMovie() {
	var title string
	var lengthStr string
	lengthRegexp, _ := regexp.Compile(`^\d{2}:\d{2}$`)
	var directorName string
	var releaseYear int
	var actorName string
	var actorID int
	var directorID int

	fmt.Print("Title: ")
	title, _ = command.GetInputString()

	for {
		fmt.Print("Length (hh:mm): ")
		fmt.Scanln(&lengthStr)
		if matched := lengthRegexp.MatchString(lengthStr); matched {
			break
		} else { fmt.Println("Bad input format (hh:mm), try again!") }
	}
	lengthParts := strings.Split(lengthStr, ":")
	hours, _ := strconv.Atoi(lengthParts[0])
	minutes, _ := strconv.Atoi(lengthParts[1])
	length := hours*60 + minutes

	for {
		fmt.Print("Director: ")

		directorName, _ = command.GetInputString()
	
		directorID = personRepository.DoesPersonExist(directorName)
		if directorID == -1 {
			fmt.Printf("We could not find \"%s\", try again!\n", directorName)
		} else { break }
	}

	for {
		fmt.Print("Released in: ")
		var works bool
		releaseYear, works = command.GetInputInteger()
		if !works {
			fmt.Println("Invalid release year, must be a positive integer.")
		} else { break }
	}

	var actors []int
	fmt.Println("Starring (type 'exit' to finish): ")
	for {
		fmt.Print("> ")
		actorName, _ = command.GetInputString()

		if actorName == "exit" { break }

		actorID = personRepository.DoesPersonExist(actorName)
		if actorID == -1 {
			fmt.Printf("We could not find \"%s\", try again!\n", actorName)
		} else {
			actors = append(actors, actorID)
		}
	}

	newID := movieRepository.GetMoviesLength() + 1
	newMovie := repository.Movie{
		ID:         newID,
		Title:      title,
		DirectorID: directorID,
		ReleaseYear: releaseYear,
		Actors:     actors,
		Length:     length,
	}
	movieRepository.AddMovie(&newMovie)
	fmt.Println("Movie added successfully!")
}