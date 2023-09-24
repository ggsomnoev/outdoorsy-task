package repository

import (
	"fmt"
	"simple-rentals-api/entity"	
)

type RentalRepository interface {
	GetRental(id int) ([]entity.Rental, error)
	GetRentals(params entity.QueryParams) ([]entity.Rental, error)
}

type repo struct{}

//NewRentalsRepository creates a new rentals repository
func NewRentalsRepository() RentalRepository {
	return &repo{}
}

//GetRental returns a single rental
func (*repo) GetRental(id int) ([]entity.Rental, error) {
	
	query := fmt.Sprintf("%v AND r.id = %v", BaseQuery, id)

	return ExecuteDatabaseQuery(query)
}

//GetRentals returns multiple rentals
func (*repo) GetRentals(params entity.QueryParams) ([]entity.Rental, error) {
	
	query := GenerateSQLQuery(params)

	return ExecuteDatabaseQuery(query)
}