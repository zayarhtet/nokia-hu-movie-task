package repository

import "fmt"

var wholeData *Data
var filename string

func Init(fileName string) {
	filename = fileName
	fmt.Println("Initializing database...")

	var err error
	wholeData, err = loadData(filename)

	if err != nil {
		fmt.Printf("Error loading data: %v\n", err)
        panic("Failed to initialize database")
	}

	fmt.Println("Database initialized successfully.")
}

func Close() {
	fmt.Println("Closing database...")

    saveData(filename, wholeData)

    fmt.Println("Database closed successfully.")
}

func NewMovieRepository() MovieRepository{
	return &MovieRepositoryImpl{data: wholeData}
}

func NewPersonRepository() PersonRepository{
	return &PersonRepositoryImpl{data: wholeData}
}