package contextx

import (
	"context"

	"github.com/gin-gonic/gin"
)

// AddContextxMiddleware is used to add contextx middleware.
func AddContextxMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a new Contextx instance
		ctxx := WithContext(c.Request.Context())

		// Store the Contextx instance in the request context
		newCtx := context.WithValue(c.Request.Context(), KeyCtx, ctxx)

		// Update the request context
		c.Request = c.Request.WithContext(newCtx)

		// Continue to the next handler
		c.Next()
	}
}
