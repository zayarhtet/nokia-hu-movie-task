package repository

type PersonRepository interface {
	GetAllPerson() map[int]Person
	GetPersonName(int) string
	DoesPersonExist(string) int
	GetPersonLength() int
	AddPerson(*Person)
	DeletePerson(int)
}

type PersonRepositoryImpl struct {
	data *Data
}


func (m *PersonRepositoryImpl) GetAllPerson() map[int]Person {
	return m.data.Persons
}

func (m *PersonRepositoryImpl) GetPersonName(id int) string {
	person := m.data.Persons[id]

    return person.Name
}

func (m *PersonRepositoryImpl) DoesPersonExist(name string) int {
	for id, person := range m.data.Persons {
        if person.Name == name {
            return id
        }
    }
	return -1
}

func (m *PersonRepositoryImpl) GetPersonLength() int {
	return len(m.data.Persons)
}

func (m *PersonRepositoryImpl) AddPerson(person *Person) {
	m.data.Persons[person.ID] = *person
}

func (m *PersonRepositoryImpl) DeletePerson(personID int) {
	delete(m.data.Persons, personID)
}