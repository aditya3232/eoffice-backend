package handler

import (
	"encoding/json"
	"eoffice-backend/constant"
	"eoffice-backend/helper"
	"eoffice-backend/models/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	filter := helper.QueryParamsToMap(c, user.User{})
	page := helper.NewPagination(helper.StrToInt(c.Query("page")), helper.StrToInt(c.Query("limit")))
	sort := helper.NewSort(c.Query("sort"), c.Query("order"))

	Users, page, err := h.userService.Get(filter, page, sort)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, helper.Pagination{}, errorMessage)

		c.JSON(response.Meta.Code, response)
		return
	}

	if len(Users) == 0 {
		errorMessage := gin.H{"errors": "User tidak ditemukan"}
		response := helper.APIResponse(constant.DataNotFound, http.StatusNotFound, helper.Pagination{}, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	response := helper.APIResponse(constant.DataFound, http.StatusOK, page, user.FormatUsers(Users))
	c.JSON(response.Meta.Code, response)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))

	User, err := h.userService.GetByID(userID)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if User.ID == 0 {
		errorMessage := gin.H{"errors": "User tidak ditemukan"}
		response := helper.APIResponse(constant.DataNotFound, http.StatusNotFound, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.APIResponse(constant.DataFound, http.StatusOK, helper.Pagination{}, user.FormatUser(User))
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input user.CreateInput

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

	newUser, err := h.userService.Create(
		user.User{
			EmployeeID: input.EmployeeID,
			Password:   input.Password,
			RoleID:     input.RoleID,
		},
	)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.FailedCreateData, http.StatusBadRequest, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessCreateData, http.StatusOK, helper.Pagination{}, user.FormatUser(newUser))
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var input user.UpdateInput

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

	userID, _ := strconv.Atoi(c.Param("id"))

	userData, err := h.userService.Update(
		user.User{
			ID:       userID,
			Remarks:  input.Remarks,
			Password: input.Password,
			RoleID:   input.RoleID,
		},
	)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.FailedUpdateData, http.StatusBadRequest, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessUpdateData, http.StatusOK, helper.Pagination{}, user.FormatUser(userData))
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))

	err := h.userService.Delete(userID)
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
