package mock

import (
	"fmt"

	"com.nokia.zayar/repository"
)

type MockMovieRepository struct{
	Movies []repository.Movie
}

// AddMovie implements repository.MovieRepository.
func (m *MockMovieRepository) AddMovie(*repository.Movie) {
	panic("unimplemented")
}

// DeletePersonFromMovies implements repository.MovieRepository.
func (m *MockMovieRepository) DeletePersonFromMovies(int) {
	panic("unimplemented")
}

// GetMoviesLength implements repository.MovieRepository.
func (m *MockMovieRepository) GetMoviesLength() int {
	panic("unimplemented")
}

// MovieHasThisDirector implements repository.MovieRepository.
func (m *MockMovieRepository) MovieHasThisDirector(int) string {
	panic("unimplemented")
}

// PrintMovie implements repository.MovieRepository.
func (m *MockMovieRepository) PrintMovie(movie repository.Movie, verbose bool) {
	fmt.Printf("%s by %d in %d, %d\n", movie.Title, movie.DirectorID, movie.ReleaseYear, movie.Length)

	if verbose {
        fmt.Println("    Starrings:")
		for _, actorID := range movie.Actors {
			fmt.Printf("        - %d\n", actorID)
		}
    }
}

func (m *MockMovieRepository) GetAllMovies(asc, desc bool) []repository.Movie {
	return m.Movies
}
