package contextx

import (
	"context"

	"github.com/gin-gonic/gin"
)

// AddContextxMiddleware is used to add contextx middleware.
func AddContextxMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), KeyContextx, WithContextLegacy(c.Request.Context()))

		c.Request = c.Request.WithContext(ctx)

		// Continue to the next handler
		c.Next()
	}
}
