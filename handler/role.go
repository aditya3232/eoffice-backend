package handler

import (
	"encoding/json"
	"eoffice-backend/constant"
	"eoffice-backend/helper"
	"eoffice-backend/models/role"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RoleHandler struct {
	roleService role.Service
}

func NewRoleHandler(roleService role.Service) *RoleHandler {
	return &RoleHandler{roleService}
}

func (h *RoleHandler) GetRole(c *gin.Context) {
	filter := helper.QueryParamsToMap(c, role.Role{})
	page := helper.NewPagination(helper.StrToInt(c.Query("page")), helper.StrToInt(c.Query("limit")))
	sort := helper.NewSort(c.Query("sort"), c.Query("order"))

	Role, page, err := h.roleService.Get(filter, page, sort)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, helper.Pagination{}, errorMessage)

		c.JSON(response.Meta.Code, response)
		return
	}

	if len(Role) == 0 {
		errorMessage := gin.H{"errors": "Role tidak ditemukan"}
		response := helper.APIResponse(constant.DataNotFound, http.StatusNotFound, helper.Pagination{}, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	response := helper.APIResponse(constant.DataFound, http.StatusOK, page, role.FormatRoles(Role))
	c.JSON(response.Meta.Code, response)
}

func (h *RoleHandler) GetRoleByID(c *gin.Context) {
	roleID, _ := strconv.Atoi(c.Param("id"))

	Role, err := h.roleService.GetByID(roleID)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if Role.ID == 0 {
		errorMessage := gin.H{"errors": "Role tidak ditemukan"}
		response := helper.APIResponse(constant.DataNotFound, http.StatusNotFound, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse(constant.DataFound, http.StatusOK, helper.Pagination{}, role.FormatRole(Role))
	c.JSON(http.StatusOK, response)
}

func (h *RoleHandler) CreateRole(c *gin.Context) {
	var input role.CreateInput

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

	roleStruct := role.Role{
		Nama: input.Nama,
	}

	newRole, err := h.roleService.Create(roleStruct)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.FailedCreateData, http.StatusBadRequest, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessCreateData, http.StatusOK, helper.Pagination{}, role.FormatRole(newRole))
	c.JSON(http.StatusOK, response)
}

func (h *RoleHandler) UpdateRole(c *gin.Context) {
	var input role.UpdateInput

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

	roleID, _ := strconv.Atoi(c.Param("id"))

	roleStruct := role.Role{
		ID:   roleID,
		Nama: input.Nama,
	}

	roleData, err := h.roleService.Update(roleStruct)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.FailedUpdateData, http.StatusBadRequest, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessUpdateData, http.StatusOK, helper.Pagination{}, role.FormatRole(roleData))
	c.JSON(http.StatusOK, response)
}

func (h *RoleHandler) DeleteRole(c *gin.Context) {
	roleID, _ := strconv.Atoi(c.Param("id"))

	err := h.roleService.Delete(roleID)
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
