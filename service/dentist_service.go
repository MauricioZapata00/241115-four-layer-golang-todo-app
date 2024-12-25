package service

import (
	"errors"
	"four-layer-todo-app/model"
	"four-layer-todo-app/repository"
	"log"
)

type DentistService interface {
	SaveDentist(doc *model.Dentist) (*model.Dentist, error)
	ListAllDentists() ([]*model.Dentist, error)
}

type dentistService struct {
	repo repository.DentistRepository
}

func (d *dentistService) SaveDentist(doc *model.Dentist) (*model.Dentist, error) {
	if doc == nil {
		return nil, errors.New("doctor cannot be nil")
	}
	if doc.DNI == "" {
		return nil, errors.New("DNI cannot be empty")
	}
	dentistFound, err := d.repo.FindDentistByDni(doc.DNI)
	if err != nil {
		log.Println("Error finding dentist", err)
		return nil, err
	}
	if dentistFound != nil {
		return nil, errors.New("doctor already exists with DNI: " + dentistFound.DNI)
	}
	// Create the new dentist
	err = d.repo.CreateDentist(doc)
	if err != nil {
		log.Println("Error creating dentist:", err)
		return nil, err
	}
	return doc, nil
}

func (d *dentistService) ListAllDentists() ([]*model.Dentist, error) {
	return d.repo.FindAllDentists()
}

func NewDentistService(repo repository.DentistRepository) DentistService {
	return &dentistService{repo: repo}
}
