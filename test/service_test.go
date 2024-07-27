package test

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"

	"com.nokia.zayar/repository"
	"com.nokia.zayar/service"
	"com.nokia.zayar/test/mock"
	"github.com/stretchr/testify/assert"
)

func captureOutput(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		log.Fatalf("Failed to create pipe: %v", err)
	}
	defer r.Close()
	defer w.Close()

	old := os.Stdout
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	f()

	os.Stdout = old
	w.Close()
	return <-outC
}

func TestListMovies(t *testing.T) {
	service.Init()
	service.SetMovieRepository(&mock.MockMovieRepository{
		Movies: []repository.Movie{
			{Title: "Inception", DirectorID: 1, ReleaseYear: 2010, Length: 148, Actors: []int{1, 2, 3}},
			{Title: "Interstellar", DirectorID: 1, ReleaseYear: 2014, Length: 169, Actors: []int{4}},
			{Title: "Nokia Movie", DirectorID: 2, ReleaseYear: 2025, Length: 175, Actors: []int{2, 5}},
			{Title: "Nokia Movie I.", DirectorID: 2, ReleaseYear: 2025, Length: 115, Actors: []int{2, 5}},
		},
	})
	service.SetPersonRepository(&mock.MockPersonRepository{})

	testCases := []struct {
		name         string
		verbose      bool
		titleRegex   string
		directorRegex string
		actorRegex   string
		expected     string
	}{
		{
			name:         "No filters, verbose off",
			verbose:      false,
			titleRegex:   "",
			directorRegex: "",
			actorRegex:   "",
			expected: `Inception by 1 in 2010, 148
Interstellar by 1 in 2014, 169
Nokia Movie by 2 in 2025, 175
Nokia Movie I. by 2 in 2025, 115
`,
		},
		{
			name:         "Title filter, verbose on",
			verbose:      true,
			titleRegex:   "Nokia",
			directorRegex: "",
			actorRegex:   "",
			expected: `Nokia Movie by 2 in 2025, 175
    Starrings:
        - 2
        - 5
Nokia Movie I. by 2 in 2025, 115
    Starrings:
        - 2
        - 5
`,
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := captureOutput(func() {
				service.ListItems(tc.verbose, tc.titleRegex, tc.directorRegex, tc.actorRegex, false, false)
			})

			assert.Equal(t, tc.expected, output)
		})
	}

}