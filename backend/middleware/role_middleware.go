package middleware

import (
	"TicketManagement/errorhandler"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(roleID int) gin.HandlerFunc {
	return func(c *gin.Context) {
		rolesVal, exists := c.Get("roles")
		if !exists {
			errorhandler.ErrorHandler(c, &errorhandler.ForbiddenError{Message: "You cannot access this API"})
			c.Abort()
			return
		}

		roles, ok := rolesVal.([]string)
		if !ok {
			errorhandler.ErrorHandler(c, &errorhandler.ForbiddenError{Message: "You cannot access this API"})
			c.Abort()
			return
		}

		for _, r := range roles {
			if r == strconv.Itoa(roleID) {
				c.Next()
				return
			}
		}

		errorhandler.ErrorHandler(c, &errorhandler.ForbiddenError{Message: "You cannot access this API"})
		c.Abort()
	}
}
