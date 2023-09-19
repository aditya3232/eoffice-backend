package handler

import (
	"encoding/json"
	"eoffice-backend/constant"
	"eoffice-backend/helper"
	"eoffice-backend/models/auth"
	"eoffice-backend/models/employee"
	"eoffice-backend/models/user"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authService     auth.Service
	employeeService employee.Service
}

func NewAuthHandler(authService auth.Service, employeeService employee.Service) *AuthHandler {
	return &AuthHandler{authService, employeeService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input auth.LoginInput

	err := c.ShouldBind(&input)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errors := helper.FormatValidationError(err)
			errorMessage := gin.H{"errors": errors}
			response := helper.APIResponse("Login gagal", http.StatusUnprocessableEntity, helper.Pagination{}, errorMessage)
			c.JSON(response.Meta.Code, response)
			return
		case *json.UnmarshalTypeError:
			errors := helper.UnmarshalError(err)
			errorMessage := gin.H{"errors": errors}
			response := helper.APIResponse("Login gagal", http.StatusUnprocessableEntity, helper.Pagination{}, errorMessage)
			c.JSON(response.Meta.Code, response)
			return
		}
	}

	employee, _, err := h.employeeService.Get(
		map[string]string{
			"nip": strconv.Itoa(input.Nip),
		},
		helper.NewPagination(1, 1),
		helper.NewSort("id", "asc"),
	)

	if err != nil {
		errorMessage := helper.FormatError(err)
		response := helper.APIResponse("Login gagal", http.StatusBadRequest, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if len(employee) == 0 {
		response := helper.APIResponse(constant.DataNotFound, http.StatusNotFound, helper.Pagination{}, gin.H{"error": helper.FormatError(errors.New("User tidak ditemukan"))})
		c.JSON(response.Meta.Code, response)
		return
	}

	userLogin, err := h.authService.Login(
		user.User{
			EmployeeID: int(employee[0].ID),
			Password:   input.Password,
		},
	)

	if err != nil {
		response := helper.APIResponse("Login gagal", http.StatusBadRequest, helper.Pagination{}, gin.H{"error": helper.FormatError(err)})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Login berhasil", http.StatusOK, helper.Pagination{},
		gin.H{
			"auth": gin.H{
				"token":   userLogin.Token,
				"expires": helper.DateTimeToString(time.Now().AddDate(0, 0, 30)),
			},
			"user": user.FormatUser(userLogin),
		})

	c.JSON(response.Meta.Code, response)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		response := helper.APIResponse("Logout gagal", http.StatusBadRequest, helper.Pagination{},
			gin.H{
				"error": helper.FormatError(errors.New("Token tidak ditemukan")),
			})
		c.JSON(response.Meta.Code, response)
		return
	}

	err := h.authService.Logout(token)
	if err != nil {
		response := helper.APIResponse("Logout gagal", http.StatusBadRequest, helper.Pagination{},
			gin.H{
				"error": helper.FormatError(err),
			})
		c.JSON(response.Meta.Code, response)
		return
	}

	response := helper.APIResponse("Logout berhasil", http.StatusOK, helper.Pagination{}, nil)
	c.JSON(response.Meta.Code, response)
}
