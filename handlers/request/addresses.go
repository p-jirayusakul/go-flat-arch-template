package request

import "github.com/jackc/pgx/v5/pgtype"

type CreateAddressesRequest struct {
	Address    pgtype.Text `json:"address" validate:"max=255" example:"22 Rue du Grenier"`
	City       string      `json:"city" validate:"required,max=100" example:"Paris"`
	Province   string      `json:"province" validate:"required,max=100" example:"Saint-Lazare"`
	PostalCode string      `json:"postalCode" validate:"required,max=20" example:"75003"`
	Country    string      `json:"country" validate:"required,max=100" example:"France"`
}

type UpdateAddressesRequest struct {
	ID         string      `param:"id" validate:"uuid4,required" swaggerignore:"true"`
	Address    pgtype.Text `json:"address" validate:"max=255" example:"22 Rue du Grenier"`
	City       string      `json:"city" validate:"required,max=100" example:"Paris"`
	Province   string      `json:"province" validate:"required,max=100" example:"Saint-Lazare"`
	PostalCode string      `json:"postalCode" validate:"required,max=20" example:"75003"`
	Country    string      `json:"country" validate:"required,max=100" example:"France"`
}

type DeleteAddressesRequest struct {
	ID string `param:"id" validate:"uuid4,required"`
}
