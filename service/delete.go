package service

import (
	"fmt"
)

func DeleteItem(personName string) {
	if len(personName) == 0 { return }

	var personID int = personRepository.DoesPersonExist(personName)
	
	if personID == -1 {
		fmt.Printf("Person \"%s\" not found in the database.\n", personName)
		return
	}

	if title := movieRepository.MovieHasThisDirector(personID); len(title) != 0 {
		fmt.Printf("Cannot delete \"%s\" because they are a director of the movie \"%s\".\n", personName, title)
		return
	}

	personRepository.DeletePerson(personID)
	movieRepository.DeletePersonFromMovies(personID)

	fmt.Printf("Person \"%s\" deleted successfully!\n", personName)

}