package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	database "github.com/p-jirayusakul/go-flat-arch-template/database/sqlc"
	"github.com/p-jirayusakul/go-flat-arch-template/handlers/request"
	"github.com/p-jirayusakul/go-flat-arch-template/utils"
)

func (s *ServerHttpHandler) CreateAddresses(c echo.Context) (err error) {
	ctx := context.Background()

	// pare json
	body := new(request.CreateAddressesRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Logic
	err = s.GetTokenID(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	arg := database.CreateAddressesParams{
		StreetAddress: body.Address,
		City:          body.City,
		StateProvince: body.Province,
		PostalCode:    body.PostalCode,
		Country:       body.Country,
		AccountsID:    c.Get("accountsID").(string),
	}

	_, err = s.store.CreateAddresses(ctx, arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Response
	var payload interface{}
	message := "create addresses completed"
	return utils.RespondWithJSON(c, http.StatusCreated, message, payload)
}

func (s *ServerHttpHandler) ListAddresses(c echo.Context) (err error) {
	ctx := context.Background()

	// Logic
	err = s.GetTokenID(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	result, err := s.store.ListAddressesByAccountId(ctx, c.Get("accountsID").(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Response
	message := "get addresses completed"
	return utils.RespondWithJSON(c, http.StatusOK, message, result)
}

func (s *ServerHttpHandler) UpdateAddresses(c echo.Context) (err error) {
	ctx := context.Background()

	// pare json
	body := new(request.UpdateAddressesRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Logic
	err = s.GetTokenID(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	isAlreadyExists, err := s.store.IsAddressesAlreadyExists(ctx, database.IsAddressesAlreadyExistsParams{
		ID:         body.ID,
		AccountsID: c.Get("accountsID").(string),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if !isAlreadyExists {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrDataNotFound.Error())
	}

	arg := database.UpdateAddressByIdParams{
		ID:            body.ID,
		StreetAddress: body.Address,
		City:          body.City,
		StateProvince: body.Province,
		PostalCode:    body.PostalCode,
		Country:       body.Country,
		AccountsID:    c.Get("accountsID").(string),
	}

	err = s.store.UpdateAddressById(ctx, arg)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Response
	var payload interface{}
	message := "update addresses completed"
	return utils.RespondWithJSON(c, http.StatusOK, message, payload)
}

func (s *ServerHttpHandler) DeleteAddresses(c echo.Context) (err error) {
	ctx := context.Background()

	// pare json
	body := new(request.DeleteAddressesRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Logic
	err = s.GetTokenID(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	isAlreadyExists, err := s.store.IsAddressesAlreadyExists(ctx, database.IsAddressesAlreadyExistsParams{
		ID:         body.ID,
		AccountsID: c.Get("accountsID").(string),
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if !isAlreadyExists {
		return echo.NewHTTPError(http.StatusNotFound, utils.ErrDataNotFound.Error())
	}

	err = s.store.DeleteAddressesById(ctx, body.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// Response
	var payload interface{}
	message := "deleted addresses completed"
	return utils.RespondWithJSON(c, http.StatusNoContent, message, payload)
}
