package handler

import (
	"encoding/json"
	"eoffice-backend/constant"
	"eoffice-backend/helper"
	"eoffice-backend/models/employee"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EmployeeHandler struct {
	employeeService employee.Service
}

func NewEmployeeHandler(employeeService employee.Service) *EmployeeHandler {
	return &EmployeeHandler{employeeService}
}

func (h *EmployeeHandler) GetEmployee(c *gin.Context) {
	filter := helper.QueryParamsToMap(c, employee.Employee{})
	page := helper.NewPagination(helper.StrToInt(c.Query("page")), helper.StrToInt(c.Query("limit")))
	sort := helper.NewSort(c.Query("sort"), c.Query("order"))

	Employee, page, err := h.employeeService.Get(filter, page, sort)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, helper.Pagination{}, errorMessage)

		c.JSON(response.Meta.Code, response)
		return
	}

	if len(Employee) == 0 {
		errorMessage := gin.H{"errors": "Employee tidak ditemukan"}
		response := helper.APIResponse(constant.DataNotFound, http.StatusNotFound, helper.Pagination{}, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	response := helper.APIResponse(constant.DataFound, http.StatusOK, page, employee.FormatEmployees(Employee))
	c.JSON(response.Meta.Code, response)
}

func (h *EmployeeHandler) GetEmployeeByID(c *gin.Context) {
	employeeID, _ := strconv.Atoi(c.Param("id"))

	Employee, err := h.employeeService.GetByID(employeeID)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if Employee.ID == 0 {
		errorMessage := gin.H{"errors": "Employee tidak ditemukan"}
		response := helper.APIResponse(constant.DataNotFound, http.StatusNotFound, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse(constant.DataFound, http.StatusOK, helper.Pagination{}, employee.FormatEmployee(Employee))
	c.JSON(http.StatusOK, response)
}

func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var input employee.CreateInput

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

	employeeStruct := employee.Employee{
		Nama:           input.Nama,
		Nip:            input.Nip,
		TempatLahir:    input.TempatLahir,
		TanggalLahir:   helper.StringToDate(input.TanggalLahir),
		Alamat:         input.Alamat,
		NoHp:           input.NoHp,
		EmailPersonal:  input.EmailPersonal,
		EmailCorporate: input.EmailCorporate,
		DivisionID:     input.DivisionID,
		PositionID:     input.PositionID,
		StartDate:      helper.StringToDate(input.StartDate),
		Avatar:         input.Avatar,
	}

	if input.EndDate != "" {
		endDate := helper.StringToDate(input.EndDate)
		employeeStruct.EndDate = &endDate
	} else {
		employeeStruct.EndDate = nil
	}

	newEmployee, err := h.employeeService.Create(employeeStruct)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.FailedCreateData, http.StatusBadRequest, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessCreateData, http.StatusOK, helper.Pagination{}, employee.FormatEmployee(newEmployee))
	c.JSON(http.StatusOK, response)
}

func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	var input employee.UpdateInput

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

	employeeID, _ := strconv.Atoi(c.Param("id"))

	employeeStruct := employee.Employee{
		ID:             employeeID,
		Nama:           input.Nama,
		Nip:            input.Nip,
		TempatLahir:    input.TempatLahir,
		TanggalLahir:   helper.StringToDate(input.TanggalLahir),
		Alamat:         input.Alamat,
		NoHp:           input.NoHp,
		EmailPersonal:  input.EmailPersonal,
		EmailCorporate: input.EmailCorporate,
		DivisionID:     input.DivisionID,
		PositionID:     input.PositionID,
		StartDate:      helper.StringToDate(input.StartDate),
		Avatar:         input.Avatar,
	}

	if input.EndDate != "" {
		endDate := helper.StringToDate(input.EndDate)
		employeeStruct.EndDate = &endDate
	} else {
		employeeStruct.EndDate = nil
	}

	employeeData, err := h.employeeService.Update(employeeStruct)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.FailedUpdateData, http.StatusBadRequest, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessUpdateData, http.StatusOK, helper.Pagination{}, employee.FormatEmployee(employeeData))
	c.JSON(http.StatusOK, response)
}

func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	employeeID, _ := strconv.Atoi(c.Param("id"))

	err := h.employeeService.Delete(employeeID)
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
