// Package restful Code generated by swaggo/swag. DO NOT EDIT
package restful

import "github.com/swaggo/swag"

const docTemplateorder_restful = `{
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
        "/v1/orders": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get order list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get order list",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "default": 1,
                        "description": "Page is the page number.",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "RestaurantID is the ID of the restaurant that received the order.",
                        "name": "restaurant_id",
                        "in": "query"
                    },
                    {
                        "maximum": 100,
                        "minimum": 1,
                        "type": "integer",
                        "default": 10,
                        "description": "Size is the number of items per page.",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "pending",
                            "confirmed",
                            "preparing",
                            "delivering",
                            "delivered",
                            "canceled"
                        ],
                        "type": "string",
                        "description": "Status is the status of the order.",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "UserID is the ID of the user who placed the order.",
                        "name": "user_id",
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
                                                "$ref": "#/definitions/github_com_blackhorseya_godine_entity_order_model.Order"
                                            }
                                        }
                                    }
                                }
                            ]
                        },
                        "headers": {
                            "X-Total-Count": {
                                "type": "int",
                                "description": "Total number of items"
                            }
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
                "description": "Create a new order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Create a new order",
                "parameters": [
                    {
                        "description": "order payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/orders.PostPayload"
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
                                            "$ref": "#/definitions/github_com_blackhorseya_godine_entity_order_model.Order"
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
        "/v1/orders/{order_id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get order by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Get order by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "order id",
                        "name": "order_id",
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
                                            "$ref": "#/definitions/github_com_blackhorseya_godine_entity_order_model.Order"
                                        }
                                    }
                                }
                            ]
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
        "github_com_blackhorseya_godine_entity_order_model.Order": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "CreatedAt is the timestamp when the order was created.",
                    "type": "string"
                },
                "id": {
                    "description": "ID is the unique identifier of the order.",
                    "type": "string"
                },
                "items": {
                    "description": "Items are the list of items in the order.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_blackhorseya_godine_entity_order_model.OrderItem"
                    }
                },
                "restaurant_id": {
                    "description": "RestaurantID is the identifier of the restaurant where the order was placed.",
                    "type": "string"
                },
                "status": {
                    "description": "Status is the current status of the order (e.g., pending, confirmed, delivered)."
                },
                "total_amount": {
                    "description": "TotalAmount is the total amount of the order.",
                    "type": "number"
                },
                "updated_at": {
                    "description": "UpdatedAt is the timestamp when the order was last updated.",
                    "type": "string"
                },
                "user_id": {
                    "description": "UserID is the identifier of the user who placed the order.",
                    "type": "string"
                }
            }
        },
        "github_com_blackhorseya_godine_entity_order_model.OrderItem": {
            "type": "object",
            "properties": {
                "menu_item_id": {
                    "description": "MenuItemID is the identifier of the menu item.",
                    "type": "string",
                    "example": "174e9519-4c47-42f2-bb1c-b0eaa8f76d05"
                },
                "price": {
                    "description": "Price is the price of a single unit of the menu item.",
                    "type": "number",
                    "example": 10
                },
                "quantity": {
                    "description": "Quantity is the quantity of the menu item ordered.",
                    "type": "integer",
                    "example": 2
                }
            }
        },
        "orders.PostPayload": {
            "type": "object",
            "required": [
                "items",
                "restaurant_id",
                "user_id"
            ],
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_blackhorseya_godine_entity_order_model.OrderItem"
                    }
                },
                "restaurant_id": {
                    "type": "string",
                    "example": "a1dbb32b-05f0-4354-8253-60f4c6deae12"
                },
                "user_id": {
                    "type": "string",
                    "example": "adcf23bc-cd32-4176-8d46-68f15ebdfa98"
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

// SwaggerInfoorder_restful holds exported Swagger Info so clients can modify it
var SwaggerInfoorder_restful = &swag.Spec{
	Version:          "0.1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Godine Order Restful API",
	Description:      "Godine Order Restful API document.",
	InfoInstanceName: "order_restful",
	SwaggerTemplate:  docTemplateorder_restful,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfoorder_restful.InstanceName(), SwaggerInfoorder_restful)
}
