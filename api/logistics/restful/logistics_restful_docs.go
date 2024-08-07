// Package restful Code generated by swaggo/swag. DO NOT EDIT
package restful

import "github.com/swaggo/swag"

const docTemplatelogistics_restful = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Sean Zheng",
            "url": "https://blog.seancheng.space",
            "email": "blackhorseya@gmail.com"
        },
        "license": {
            "name": "GPL-3.0",
            "url": "https://spdx.org/licenses/GPL-3.0-only.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/healthz": {
            "get": {
                "description": "Check the health of the service.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Check the health of the service.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsex.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responsex.Response"
                        }
                    }
                }
            }
        },
        "/v1/deliveries": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get the list of deliveries",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deliveries"
                ],
                "summary": "Get the list of deliveries",
                "parameters": [
                    {
                        "type": "string",
                        "description": "driver id: adcf23bc-cd32-4176-8d46-68f15ebdfa98",
                        "name": "driver_id",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "default": 1,
                        "description": "Page specifies the page number for pagination.",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "maximum": 100,
                        "minimum": 1,
                        "type": "integer",
                        "default": 10,
                        "description": "Size specifies the number of items per page.",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/responsex.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/github_com_blackhorseya_godine_entity_domain_logistics_model.Delivery"
                                            }
                                        }
                                    }
                                }
                            ]
                        },
                        "headers": {
                            "X-Total-Count": {
                                "type": "number",
                                "description": "total count"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsex.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responsex.Response"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create a new delivery",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deliveries"
                ],
                "summary": "Create a new delivery",
                "parameters": [
                    {
                        "description": "delivery request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_blackhorseya_godine_entity_domain_logistics_model.Delivery"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/responsex.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/github_com_blackhorseya_godine_entity_domain_logistics_model.Delivery"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsex.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responsex.Response"
                        }
                    }
                }
            }
        },
        "/v1/deliveries/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get the delivery by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deliveries"
                ],
                "summary": "Get the delivery by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "delivery id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/responsex.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/github_com_blackhorseya_godine_entity_domain_logistics_model.Delivery"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsex.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responsex.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responsex.Response"
                        }
                    }
                }
            }
        },
        "/v1/deliveries/{id}/status": {
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update the delivery status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "deliveries"
                ],
                "summary": "Update the delivery status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "delivery id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "delivery status request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/deliveries.PatchWithStatusPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/responsex.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/github_com_blackhorseya_godine_entity_domain_logistics_model.Delivery"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responsex.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responsex.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responsex.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "deliveries.PatchWithStatusPayload": {
            "type": "object",
            "required": [
                "status"
            ],
            "properties": {
                "status": {
                    "type": "string",
                    "example": "delivered"
                }
            }
        },
        "github_com_blackhorseya_godine_entity_domain_logistics_model.Delivery": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "CreatedAt is the timestamp when the delivery was created.",
                    "type": "string"
                },
                "delivery_time": {
                    "description": "DeliveryTime is the timestamp when the delivery was completed.",
                    "type": "string"
                },
                "driver_id": {
                    "description": "DriverID is the identifier of the driver assigned to the delivery.",
                    "type": "string"
                },
                "id": {
                    "description": "ID is the unique identifier of the delivery.",
                    "type": "string"
                },
                "order_id": {
                    "description": "OrderID is the identifier of the order associated with the delivery.",
                    "type": "string"
                },
                "pickup_time": {
                    "description": "PickupTime is the timestamp when the delivery was picked up.",
                    "type": "string"
                },
                "status": {
                    "description": "Status is the current status of the delivery (e.g., pending, in transit, delivered)."
                },
                "updated_at": {
                    "description": "UpdatedAt is the timestamp when the delivery was last updated.",
                    "type": "string"
                },
                "user_id": {
                    "description": "UserID is the identifier of the user who placed the order.",
                    "type": "string"
                }
            }
        },
        "responsex.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfologistics_restful holds exported Swagger Info so clients can modify it
var SwaggerInfologistics_restful = &swag.Spec{
	Version:          "0.1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Godine Logistics Restful API",
	Description:      "Godine Logistics Restful API document.",
	InfoInstanceName: "logistics_restful",
	SwaggerTemplate:  docTemplatelogistics_restful,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfologistics_restful.InstanceName(), SwaggerInfologistics_restful)
}
