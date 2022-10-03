package models

import (
	"errors"
	"fmt"
)

type People struct {
	// gorm.Model
	// gorm.Model `json:"model"`
	ID      int `json:"id" example:"1" format:"int64"`
	Name    string
	Address string
}

// example
var (
	ErrNameInvalidPeople = errors.New("name is empty")
)

// AddPeople example
type AddPeople struct {
	Name    string `json:"name" example:"people name"`
	Address string `json:"address" example:"address detail"`
}

// Validation example
func (data AddPeople) Validation() error {
	switch {
	case len(data.Name) == 0:
		return ErrNameInvalidPeople
	default:
		return nil
	}
}

// PeopleOne example
func PeopleOne(qData int) (People, error) {
	for _, v := range peoples {
		if qData == v.ID {
			return v, nil
		}
	}
	return People{}, nil
}

// PeoplesAll example
func PeoplesAll(q string) ([]People, error) {
	if q == "" {
		return peoples, nil
	}
	as := []People{}
	for k, v := range peoples {
		if q == v.Name {
			as = append(as, peoples[k])
		}
	}
	return as, nil
}

// Insert example
func (data People) InsertPeople() (int, error) {
	peopleMaxID++
	data.ID = peopleMaxID
	data.Name = fmt.Sprintf("%s", data.Name)
	data.Address = fmt.Sprintf("%s", data.Address)
	peoples = append(peoples, data)
	return peopleMaxID, nil
}

// UpdatePeople example
type UpdatePeople struct {
	Name string `json:"name" example:"people name"`
}

// Update example
func (a People) UpdatePeopleAction() error {
	for k, v := range peoples {
		if a.ID == v.ID {
			peoples[k].Name = a.Name
			return nil
		}
	}
	return fmt.Errorf("people id=%d is not found", a.ID)
}

// Delete example
func DeletePeopleAction(id int) error {
	for k, v := range peoples {
		if id == v.ID {
			peoples = append(peoples[:k], peoples[k+1:]...)
			return nil
		}
	}
	return fmt.Errorf("people id=%d is not found", id)
}

var peopleMaxID = 10
var peoples = []People{
	{ID: 1, Name: "people 1", Address: "address 1"},
	{ID: 2, Name: "people 2", Address: "address 2"},
	{ID: 3, Name: "people 3", Address: "address 3"},
}
