package repository

import (
	"encoding/json"
	"os"
)

type Person struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	BirthYear int    `json:"birth_year"`
}

type Movie struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	DirectorID  int    `json:"director_id"`
	ReleaseYear int    `json:"release_year"`
	Actors      []int  `json:"actors"`
	Length      int    `json:"length"`
}

type Data struct {
	Persons map[int]Person `json:"persons"`
	Movies  []Movie        `json:"movies"`
}

func loadData(filename string) (*Data, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data Data
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func saveData(filename string, data *Data) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	dataBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(dataBytes)
	return err
}