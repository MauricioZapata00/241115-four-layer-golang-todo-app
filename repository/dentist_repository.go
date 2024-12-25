package repository

import (
	"database/sql"
	"errors"
	"four-layer-todo-app/model"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jmoiron/sqlx"
	"log"
)

type DentistRepository interface {
	CreateDentist(dentist *model.Dentist) error
	FindAllDentists() ([]*model.Dentist, error)
	FindDentistByDni(dni string) (*model.Dentist, error)
}
type mysqlDentistRepository struct {
	db *sqlx.DB
}

func (o *mysqlDentistRepository) CreateDentist(dentist *model.Dentist) error {
	insertResult, err := o.db.Exec(`INSERT INTO parcialGo.dentists (nombre, apellido, matricula) values (?, ?, ?)`,
		dentist.Name, dentist.LastName, dentist.DNI)
	if err != nil {
		log.Println("Error inserting dentist", err)
	}
	lastId, err := insertResult.LastInsertId()
	if err != nil {
		log.Println("Error inserting dentist", err)
	}
	dentist.ID = uint32(lastId)
	return err
}

func (o *mysqlDentistRepository) FindAllDentists() ([]*model.Dentist, error) {
	var dentists []*model.Dentist
	err := o.db.Select(&dentists, `SELECT * FROM dentists`)
	if err != nil {
		return nil, err
	}
	return dentists, nil
}

func (o *mysqlDentistRepository) FindDentistByDni(dni string) (*model.Dentist, error) {
	var dentist model.Dentist
	err := o.db.Get(&dentist, `SELECT * FROM dentists WHERE matricula = ?`, dni)
	// When no rows are found, return nil without error
	if errors.Is(sql.ErrNoRows, err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &dentist, nil
}

func NewMySqlDentistRepository(db *sqlx.DB) DentistRepository {
	return &mysqlDentistRepository{db: db}
}
