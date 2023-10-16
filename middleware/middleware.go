package middleware

import (
	"easy-rbac/models"
	"regexp"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckPermission(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		matchedAny := false
		userId := c.MustGet("userId").(uint) // get the userId which is this after token authentication
		var user models.User
		db.Preload("Roles").Preload("Roles.Permissions").Where("ID = ?", userId).First(&user)

		for _, role := range user.Roles {
			for _, permission := range role.Permissions {
				if matches(c.Request.Method, permission.Action) && matches(c.Request.URL.Path, permission.Route) {
					matchedAny = true
					if !permission.Allowed {
						c.AbortWithStatusJSON(403, gin.H{"error": "Permission denied"})
						return
					}
				}
			}
		}

		if !matchedAny {
			c.AbortWithStatusJSON(403, gin.H{"error": "Permission denied"})
			return
		}

		c.Next()
	}
}

// case insensitive regex matching
func matches(requestValue, patternValue string) (matched bool) {
	// case insensitive: https://stackoverflow.com/a/9655186
	matched, _ = regexp.MatchString("(?i)"+patternValue, requestValue)
	return
}

func parseToken(token string) (userId uint, err error) {
	// parse the token
	return
}

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the token from the header
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		userID, err := parseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		// set the user id in the context
		c.Set("userId", userID)
		c.Next()
	}
}
