package handler

import (
	"eoffice-backend/constant"
	"eoffice-backend/helper"
	"eoffice-backend/models/permission"
	"eoffice-backend/models/profile"
	"eoffice-backend/models/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileService profile.Service
}

func NewProfileHandler(profileService profile.Service) *ProfileHandler {
	return &ProfileHandler{profileService}
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	profileID := c.MustGet("currentUser").(user.User).ID

	profileData, err := h.profileService.GetProfile(profileID)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.DataFound, http.StatusOK, helper.Pagination{}, profile.FormatProfile(profileData))
	c.JSON(response.Meta.Code, response)
}

func (h *ProfileHandler) GetPermission(c *gin.Context) {
	profileID := c.MustGet("currentUser").(user.User).ID

	permissionData, err := h.profileService.GetPermission(profileID)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, helper.Pagination{}, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if len(permissionData) == 0 {
		response := helper.APIResponse(constant.DataNotFound, http.StatusNotFound, helper.Pagination{}, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	response := helper.APIResponse(constant.DataFound, http.StatusOK, helper.Pagination{}, permission.FormatPermissions(permissionData))
	c.JSON(response.Meta.Code, response)
}
