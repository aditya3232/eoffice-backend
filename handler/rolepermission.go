package handler

import (
	"encoding/json"
	"eoffice-backend/constant"
	"eoffice-backend/helper"
	"eoffice-backend/models/rolepermission"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RolePermissionHandler struct {
	rolePermissionService rolepermission.Service
}

func NewRolePermissionHandler(rolePermissionService rolepermission.Service) *RolePermissionHandler {
	return &RolePermissionHandler{rolePermissionService}
}

func (h *RolePermissionHandler) GetRolePermission(c *gin.Context) {
	filter := helper.QueryParamsToMap(c, rolepermission.RolePermission{})
	page := helper.NewPagination(helper.StrToInt(c.Query("page")), helper.StrToInt(c.Query("limit")))
	sort := helper.NewSort(c.Query("sort"), c.Query("order"))

	RolePermission, page, err := h.rolePermissionService.Get(filter, page, sort)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, helper.Pagination{}, errorMessage)

		c.JSON(response.Meta.Code, response)
		return
	}

	if len(RolePermission) == 0 {
		errorMessage := gin.H{"errors": "Role Permission tidak ditemukan"}
		response := helper.APIResponse(constant.DataNotFound, http.StatusNotFound, helper.Pagination{}, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	response := helper.APIResponse(constant.DataFound, http.StatusOK, page, rolepermission.FormatRolePermissions(RolePermission))
	c.JSON(response.Meta.Code, response)
}

func (h *RolePermissionHandler) GetRolePermissionByID(c *gin.Context) {
	rolePermissionID, _ := strconv.Atoi(c.Param("id"))

	RolePermission, err := h.rolePermissionService.GetByID(rolePermissionID)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if RolePermission.ID == 0 {
		errorMessage := gin.H{"errors": "Role Permission tidak ditemukan"}
		response := helper.APIResponse(constant.DataNotFound, http.StatusNotFound, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse(constant.DataFound, http.StatusOK, helper.Pagination{}, rolepermission.FormatRolePermission(RolePermission))
	c.JSON(http.StatusOK, response)
}

func (h *RolePermissionHandler) CreateRolePermission(c *gin.Context) {
	var input rolepermission.CreateInput

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

	rolePermissionStruct := rolepermission.RolePermission{
		RoleID:       input.RoleID,
		PermissionID: input.PermissionID,
	}

	newRolePermission, err := h.rolePermissionService.Create(rolePermissionStruct)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.FailedCreateData, http.StatusBadRequest, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessCreateData, http.StatusOK, helper.Pagination{}, rolepermission.FormatRolePermission(newRolePermission))
	c.JSON(http.StatusOK, response)
}

func (h *RolePermissionHandler) UpdateRolePermission(c *gin.Context) {
	var input rolepermission.UpdateInput

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

	rolePermissionID, _ := strconv.Atoi(c.Param("id"))

	rolePermissionStruct := rolepermission.RolePermission{
		ID:           rolePermissionID,
		RoleID:       input.RoleID,
		PermissionID: input.PermissionID,
	}

	rolePermissionData, err := h.rolePermissionService.Update(rolePermissionStruct)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.FailedUpdateData, http.StatusBadRequest, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessUpdateData, http.StatusOK, helper.Pagination{}, rolepermission.FormatRolePermission(rolePermissionData))
	c.JSON(http.StatusOK, response)
}

func (h *RolePermissionHandler) DeleteRolePermission(c *gin.Context) {
	rolePermissionID, _ := strconv.Atoi(c.Param("id"))

	err := h.rolePermissionService.Delete(rolePermissionID)
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
