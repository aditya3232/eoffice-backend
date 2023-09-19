package handler

import (
	"encoding/json"
	"eoffice-backend/constant"
	"eoffice-backend/helper"
	"eoffice-backend/models/position"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PositionHandler struct {
	positionService position.Service
}

func NewPositionHandler(positionService position.Service) *PositionHandler {
	return &PositionHandler{positionService}
}

func (h *PositionHandler) GetPosition(c *gin.Context) {
	filter := helper.QueryParamsToMap(c, position.Position{})
	page := helper.NewPagination(helper.StrToInt(c.Query("page")), helper.StrToInt(c.Query("limit")))
	sort := helper.NewSort(c.Query("sort"), c.Query("order"))

	Position, page, err := h.positionService.Get(filter, page, sort)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, helper.Pagination{}, errorMessage)

		c.JSON(response.Meta.Code, response)
		return
	}

	if len(Position) == 0 {
		errorMessage := gin.H{"errors": "Position tidak ditemukan"}
		response := helper.APIResponse(constant.DataNotFound, http.StatusNotFound, helper.Pagination{}, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	response := helper.APIResponse(constant.DataFound, http.StatusOK, page, position.FormatPositions(Position))
	c.JSON(response.Meta.Code, response)
}

func (h *PositionHandler) GetPositionByID(c *gin.Context) {
	positionID, _ := strconv.Atoi(c.Param("id"))

	Position, err := h.positionService.GetByID(positionID)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if Position.ID == 0 {
		errorMessage := gin.H{"errors": "Position tidak ditemukan"}
		response := helper.APIResponse(constant.DataNotFound, http.StatusNotFound, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse(constant.DataFound, http.StatusOK, helper.Pagination{}, position.FormatPosition(Position))
	c.JSON(http.StatusOK, response)
}

func (h *PositionHandler) CreatePosition(c *gin.Context) {
	var input position.CreateInput

	err := c.ShouldBind(&input)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errors := helper.FormatValidationError(err)
			errorMessage := gin.H{"errors": errors}
			response := helper.APIResponse(constant.FailedCreateData, http.StatusUnprocessableEntity, helper.Pagination{}, errorMessage)
			c.JSON(response.Meta.Code, response)
			return
		case *json.UnmarshalTypeError:
			errors := helper.UnmarshalError(err)
			errorMessage := gin.H{"errors": errors}
			response := helper.APIResponse(constant.FailedCreateData, http.StatusUnprocessableEntity, helper.Pagination{}, errorMessage)
			c.JSON(response.Meta.Code, response)
			return
		}
	}

	positionStruct := position.Position{
		Nama: input.Nama,
	}

	newPosition, err := h.positionService.Create(positionStruct)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.FailedCreateData, http.StatusBadRequest, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessCreateData, http.StatusOK, helper.Pagination{}, position.FormatPosition(newPosition))
	c.JSON(http.StatusOK, response)
}

func (h *PositionHandler) UpdatePosition(c *gin.Context) {
	var input position.UpdateInput

	err := c.ShouldBind(&input)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errors := helper.FormatValidationError(err)
			errorMessage := gin.H{"errors": errors}
			response := helper.APIResponse(constant.FailedUpdateData, http.StatusUnprocessableEntity, helper.Pagination{}, errorMessage)
			c.JSON(response.Meta.Code, response)
			return
		case *json.UnmarshalTypeError:
			errors := helper.UnmarshalError(err)
			errorMessage := gin.H{"errors": errors}
			response := helper.APIResponse(constant.FailedUpdateData, http.StatusUnprocessableEntity, helper.Pagination{}, errorMessage)
			c.JSON(response.Meta.Code, response)
			return
		}
	}

	positionID, _ := strconv.Atoi(c.Param("id"))

	positionStruct := position.Position{
		ID:   positionID,
		Nama: input.Nama,
	}

	positionData, err := h.positionService.Update(positionStruct)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.FailedUpdateData, http.StatusBadRequest, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessUpdateData, http.StatusOK, helper.Pagination{}, position.FormatPosition(positionData))
	c.JSON(http.StatusOK, response)
}

func (h *PositionHandler) DeletePosition(c *gin.Context) {
	positionID, _ := strconv.Atoi(c.Param("id"))

	err := h.positionService.Delete(positionID)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.FailedDeleteData, http.StatusBadRequest, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessDeleteData, http.StatusOK, helper.Pagination{}, nil)
	c.JSON(http.StatusOK, response)
}
