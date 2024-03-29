package handlers

import (
	"context"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
	database "github.com/p-jirayusakul/go-flat-arch-template/database/sqlc"
	"github.com/p-jirayusakul/go-flat-arch-template/handlers/request"
	"github.com/p-jirayusakul/go-flat-arch-template/pkg/common"
	"github.com/p-jirayusakul/go-flat-arch-template/pkg/utils"
)

// Address
// @Summary      Create Address
// @Description  register
// @Tags         profile
// @Accept       json
// @Produce      json
// @Param request body request.CreateAddressesRequest true "body request"
// @Success      201  {object}  utils.SuccessResponse
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /api/v1/profile/addresses [post]
// @Security Bearer
func (s *ServerHttpHandler) CreateAddresses(c echo.Context) (err error) {
	ctx := context.Background()

	// pare json
	body := new(request.CreateAddressesRequest)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	// Logic
	err = s.GetTokenID(c)
	if err != nil {
		return utils.RespondWithError(http.StatusUnauthorized, err.Error())
	}

	arg := database.CreateAddressesParams{
		StreetAddress: body.Address,
		City:          body.City,
		StateProvince: body.Province,
		PostalCode:    body.PostalCode,
		Country:       body.Country,
		AccountsID:    c.Get("accountsID").(string),
	}

	_, err = s.store.CreateAddresses(ctx, &arg)
	if err != nil {
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	// Response
	var payload interface{}
	message := "create addresses completed"
	return utils.RespondWithJSON(c, http.StatusCreated, message, payload)
}

// Address
// @Summary      Get List Address
// @Description  list address
// @Tags         profile
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.SuccessResponse
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /api/v1/profile/addresses/me [get]
// @Security Bearer
func (s *ServerHttpHandler) ListAddresses(c echo.Context) (err error) {
	ctx := context.Background()

	// Logic
	err = s.GetTokenID(c)
	if err != nil {
		return utils.RespondWithError(http.StatusUnauthorized, err.Error())
	}

	result, err := s.store.ListAddressesByAccountId(ctx, c.Get("accountsID").(string))
	if err != nil {
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	// Response
	message := "get addresses completed"
	return utils.RespondWithJSON(c, http.StatusOK, message, result)
}

// Address
// @Summary      Get List Address
// @Description  list address
// @Tags         profile
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.SuccessResponse
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /api/v1/profile/addresses [get]
// @Security Bearer
func (s *ServerHttpHandler) SearchAddresses(c echo.Context) (err error) {
	ctx := context.Background()

	// pare json
	body := new(request.SearchAddressesRequest)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	// Logic
	err = s.GetTokenID(c)
	if err != nil {
		return utils.RespondWithError(http.StatusUnauthorized, err.Error())
	}

	pageNumber := body.PageNumber
	pageSize := body.PageSize

	if body.PageSize == 0 {
		pageSize = common.PAGE_SIZE
	}

	if pageSize > common.MAX_PAGE_SIZE {
		pageSize = common.MAX_PAGE_SIZE
	}

	if pageNumber > 0 {
		pageNumber = (pageNumber - 1) * pageSize
	}

	arg := database.SearchAddressesParams{
		PageNumber:    pageNumber,
		PageSize:      pageSize,
		City:          body.City,
		StateProvince: body.Province,
		PostalCode:    body.PostalCode,
		Country:       body.Country,
		AccountsID:    body.AccountsID,
		OrderBy:       body.OrderBy,
		OrderType:     body.OrderType,
	}

	result, err := s.store.SearchAddresses(ctx, arg)
	if err != nil {
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	var totalPages int
	if len(result.Data) > 0 {
		result.TotalItems = int(result.TotalItems)
		totalPages = int(math.Ceil(float64(result.TotalItems) / float64(pageSize)))
	} else {
		result.TotalItems = 0
		totalPages = 0
	}

	if body.PageNumber == 0 {
		pageNumber = 1
	} else {
		pageNumber = body.PageNumber
	}

	result.PageNumber = pageNumber
	result.PageSize = pageSize
	result.TotalPages = totalPages

	// Response
	message := "get addresses completed"
	return utils.RespondWithJSON(c, http.StatusOK, message, result)
}

// Address
// @Summary      Update Address
// @Description  update address
// @Tags         profile
// @Accept       json
// @Produce      json
// @Param        address_id   path      string  true  "Address ID"
// @Param request body request.UpdateAddressesRequest true "body request"
// @Success      200  {object}  utils.SuccessResponse
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /api/v1/profile/addresses/{address_id} [put]
// @Security Bearer
func (s *ServerHttpHandler) UpdateAddresses(c echo.Context) (err error) {
	ctx := context.Background()

	// pare json
	body := new(request.UpdateAddressesRequest)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	// Logic
	err = s.GetTokenID(c)
	if err != nil {
		return utils.RespondWithError(http.StatusUnauthorized, err.Error())
	}

	isAlreadyExists, err := s.store.IsAddressesAlreadyExists(ctx, &database.IsAddressesAlreadyExistsParams{
		ID:         body.ID,
		AccountsID: c.Get("accountsID").(string),
	})
	if err != nil {
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	if !isAlreadyExists {
		return utils.RespondWithError(http.StatusNotFound, common.ErrDataNotFound.Error())
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

	err = s.store.UpdateAddressById(ctx, &arg)
	if err != nil {
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	// Response
	var payload interface{}
	message := "update addresses completed"
	return utils.RespondWithJSON(c, http.StatusOK, message, payload)
}

// Delete Address
// @Summary      Delete Address By Address Id
// @Description  Delete Address
// @Tags         profile
// @Accept       json
// @Produce      json
// @Param        address_id   path      string  true  "Address ID"
// @Success      204
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /api/v1/profile/addresses/{address_id} [delete]
// @Security Bearer
func (s *ServerHttpHandler) DeleteAddresses(c echo.Context) (err error) {
	ctx := context.Background()

	// pare json
	body := new(request.DeleteAddressesRequest)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	// Logic
	err = s.GetTokenID(c)
	if err != nil {
		return utils.RespondWithError(http.StatusUnauthorized, err.Error())
	}

	isAlreadyExists, err := s.store.IsAddressesAlreadyExists(ctx, &database.IsAddressesAlreadyExistsParams{
		ID:         body.ID,
		AccountsID: c.Get("accountsID").(string),
	})

	if err != nil {
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	if !isAlreadyExists {
		return utils.RespondWithError(http.StatusNotFound, common.ErrDataNotFound.Error())
	}

	err = s.store.DeleteAddressesById(ctx, body.ID)
	if err != nil {
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	// Response
	var payload interface{}
	message := "deleted addresses completed"
	return utils.RespondWithJSON(c, http.StatusNoContent, message, payload)
}
