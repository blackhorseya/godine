package adapterx

import (
	"context"

	"github.com/gin-gonic/gin"
)

// ServiceLegacy is the interface that wraps the basic Serve method.
type ServiceLegacy interface {
	// Start a service asynchronously.
	Start() error

	// AwaitSignal waits for a signal to shut down the service.
	AwaitSignal() error
}

// RestfulLegacy is the interface that wraps the restful api method.
type RestfulLegacy interface {
	ServiceLegacy

	// InitRouting init the routing of restful api.
	InitRouting() error

	// GetRouter returns the router of restful api.
	GetRouter() *gin.Engine
}

// Server is the interface that wraps the Serve method.
type Server interface {
	// Start the server.
	Start(c context.Context) error

	// Shutdown the server.
	Shutdown(c context.Context) error
}

// Restful is the interface that wraps the restful api method.
type Restful interface {
	Server

	// InitRouting init the routing of restful api.
	InitRouting(c context.Context) error

	// GetRouter returns the router of restful api.
	GetRouter() *gin.Engine
}
