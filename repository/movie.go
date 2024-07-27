package repository

import (
	"fmt"
	"sort"
)

type MovieRepository interface {
	GetAllMovies(bool, bool) []Movie
    GetMoviesLength() int
    AddMovie(*Movie)
	MovieHasThisDirector(int) string
	DeletePersonFromMovies(int)
	PrintMovie(Movie, bool)
}

type MovieRepositoryImpl struct {
	data *Data
}

func (m *MovieRepositoryImpl) GetMoviesLength() int {
	return len(m.data.Movies)
}

func (m *MovieRepositoryImpl) GetAllMovies(asc, desc bool) []Movie {
	if asc {
		sort.Slice(m.data.Movies, func(i, j int) bool {
			if m.data.Movies[i].Length == m.data.Movies[j].Length {
				return m.data.Movies[i].Title < m.data.Movies[j].Title
			}
			return m.data.Movies[i].Length < m.data.Movies[j].Length
		})
	} else if (desc) {
		sort.Slice(m.data.Movies, func(i, j int) bool {
			if m.data.Movies[i].Length == m.data.Movies[j].Length {
				return m.data.Movies[i].Title < m.data.Movies[j].Title
			}
			return m.data.Movies[i].Length > m.data.Movies[j].Length
		})
	} else {
		sort.Slice(m.data.Movies, func(i, j int) bool {
			return m.data.Movies[i].Title < m.data.Movies[j].Title
		})
	}
	return m.data.Movies
}

func (m *MovieRepositoryImpl) AddMovie(movie *Movie) {
	m.data.Movies = append(m.data.Movies, *movie)
}

func (m *MovieRepositoryImpl) MovieHasThisDirector(directorID int) string {
	for _, movie := range m.data.Movies {
        if movie.DirectorID == directorID {
            return movie.Title
        }
    }
    return ""
}

func (m *MovieRepositoryImpl) DeletePersonFromMovies(personID int) {
	for i, movie := range m.data.Movies {
		newActors := []int{}
		for _, actorID := range movie.Actors {
			if actorID != personID {
				newActors = append(newActors, actorID)
			}
		}
		m.data.Movies[i].Actors = newActors
	}
}

func (m *MovieRepositoryImpl) PrintMovie(movie Movie, verbose bool) {
	fmt.Printf("%s by %s in %d, %s\n", movie.Title, m.data.Persons[movie.DirectorID].Name, movie.ReleaseYear, formatLength(movie.Length))

	if verbose {
        fmt.Println("    Starrings:")
		for _, actorID := range movie.Actors {
			fmt.Printf("        - %s at age %d\n", m.data.Persons[actorID].Name, calculateActorAge(actorID, movie.ReleaseYear))
		}
    }
}