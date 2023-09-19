package middleware

import (
	"eoffice-backend/helper"
	"eoffice-backend/library/JWT"
	"eoffice-backend/models/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header from request
		authHeader := c.GetHeader("Authorization")

		// Get user ID from JWT token
		userID, err := JWT.GetUserIDFromToken(authHeader)
		if err != nil {
			// Return 401 Unauthorized if JWT token is invalid or missing
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, helper.Pagination{}, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		user, err := userService.GetByID(userID)
		if err != nil {
			// Return 401 Unauthorized if user doesn't exist in database
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, helper.Pagination{}, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Check if user.Token is different from token headers
		if user.Token != authHeader {
			// Return 401 Unauthorized if user.Token is different from token headers
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, helper.Pagination{}, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

		// Call next middleware/handler function
		c.Next()
	}
}
