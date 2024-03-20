package request

import "github.com/jackc/pgx/v5/pgtype"

type CreateAddressesRequest struct {
	Address    pgtype.Text `json:"address" validate:"max=255"`
	City       string      `json:"city" validate:"required,max=100"`
	Province   string      `json:"province" validate:"required,max=100"`
	PostalCode string      `json:"postalCode" validate:"required,max=20"`
	Country    string      `json:"country" validate:"required,max=100"`
}

type UpdateAddressesRequest struct {
	ID         string      `param:"id" validate:"uuid4,required"`
	Address    pgtype.Text `json:"address" validate:"max=255"`
	City       string      `json:"city" validate:"required,max=100"`
	Province   string      `json:"province" validate:"required,max=100"`
	PostalCode string      `json:"postalCode" validate:"required,max=20"`
	Country    string      `json:"country" validate:"required,max=100"`
}

type DeleteAddressesRequest struct {
	ID string `param:"id" validate:"uuid4,required"`
}
