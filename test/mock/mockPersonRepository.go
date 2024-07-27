package mock

import "com.nokia.zayar/repository"

type MockPersonRepository struct{}

// AddPerson implements repository.PersonRepository.
func (m *MockPersonRepository) AddPerson(*repository.Person) {
	panic("unimplemented")
}

// DeletePerson implements repository.PersonRepository.
func (m *MockPersonRepository) DeletePerson(int) {
	panic("unimplemented")
}

// DoesPersonExist implements repository.PersonRepository.
func (m *MockPersonRepository) DoesPersonExist(string) int {
	panic("unimplemented")
}

// GetAllPerson implements repository.PersonRepository.
func (m *MockPersonRepository) GetAllPerson() map[int]repository.Person {
	panic("unimplemented")
}

// GetPersonLength implements repository.PersonRepository.
func (m *MockPersonRepository) GetPersonLength() int {
	panic("unimplemented")
}

// GetPersonName implements repository.PersonRepository.
func (m *MockPersonRepository) GetPersonName(int) string {
	panic("unimplemented")
}
