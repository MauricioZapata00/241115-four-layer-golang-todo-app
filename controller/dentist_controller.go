package controller

import (
	"four-layer-todo-app/model"
	"four-layer-todo-app/service"
	"github.com/gofiber/fiber/v2"
	"log"
)

type DentistController struct {
	ds service.DentistService
}

func NewDentistController(ds service.DentistService) *DentistController {
	return &DentistController{ds: ds}
}

func (dc *DentistController) CreateDentist(ctx *fiber.Ctx) error {
	dentist := new(model.Dentist)
	if err := ctx.BodyParser(dentist); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	dentistSaved, err := dc.ds.SaveDentist(dentist)
	if err != nil {
		log.Println("error saving dentist", err)
		return ctx.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(201).JSON(dentistSaved)
}

func (dc *DentistController) GetAllDentists(ctx *fiber.Ctx) error {
	currentDentists, err := dc.ds.ListAllDentists()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(200).JSON(currentDentists)
}
