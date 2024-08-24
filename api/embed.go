package api

import (
	"embed"
)

//go:embed gateway/apidocs.swagger.json
var GatewayOpenAPI embed.FS
