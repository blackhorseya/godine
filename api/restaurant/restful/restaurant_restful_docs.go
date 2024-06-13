// Package restful Code generated by swaggo/swag. DO NOT EDIT
package restful

import "github.com/swaggo/swag"

const docTemplaterestaurant_restful = `{
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
        "/v1/restaurants": {
            "get": {
                "description": "Get the restaurant list.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "restaurants"
                ],
                "summary": "Get the restaurant list.",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "default": 1,
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "maximum": 100,
                        "minimum": 1,
                        "type": "integer",
                        "default": 10,
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
                                                "$ref": "#/definitions/github_com_blackhorseya_godine_entity_restaurant_model.Restaurant"
                                            }
                                        }
                                    }
                                }
                            ]
                        },
                        "headers": {
                            "X-Total-Count": {
                                "type": "int",
                                "description": "total count"
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
                "description": "Create the restaurant.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "restaurants"
                ],
                "summary": "Create the restaurant.",
                "parameters": [
                    {
                        "description": "restaurant payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/restaurants.PostPayload"
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
                                            "$ref": "#/definitions/github_com_blackhorseya_godine_entity_restaurant_model.Restaurant"
                                        }
                                    }
                                }
                            ]
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
        "/v1/restaurants/{restaurant_id}": {
            "get": {
                "description": "Get the restaurant by id.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "restaurants"
                ],
                "summary": "Get the restaurant by id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "restaurant id",
                        "name": "restaurant_id",
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
                                            "$ref": "#/definitions/github_com_blackhorseya_godine_entity_restaurant_model.Restaurant"
                                        }
                                    }
                                }
                            ]
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
        "/v1/restaurants/{restaurant_id}/menu": {
            "get": {
                "description": "Get the menu list.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "Get the menu list.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "restaurant id",
                        "name": "restaurant_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "default": 1,
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "maximum": 100,
                        "minimum": 1,
                        "type": "integer",
                        "default": 10,
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
                                                "$ref": "#/definitions/github_com_blackhorseya_godine_entity_restaurant_model.MenuItem"
                                            }
                                        }
                                    }
                                }
                            ]
                        },
                        "headers": {
                            "X-Total-Count": {
                                "type": "int",
                                "description": "total count"
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
                "description": "Add a menu item.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menu"
                ],
                "summary": "Add a menu item.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "restaurant id",
                        "name": "restaurant_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "menu item payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/menu.PostPayload"
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
                                            "$ref": "#/definitions/github_com_blackhorseya_godine_entity_restaurant_model.MenuItem"
                                        }
                                    }
                                }
                            ]
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
        "github_com_blackhorseya_godine_entity_restaurant_model.Address": {
            "type": "object",
            "properties": {
                "city": {
                    "description": "City is the city where the restaurant is located.",
                    "type": "string"
                },
                "state": {
                    "description": "State is the state where the restaurant is located.",
                    "type": "string"
                },
                "street": {
                    "description": "Street is the street address of the restaurant.",
                    "type": "string"
                },
                "zip_code": {
                    "description": "ZipCode is the postal code of the restaurant's location.",
                    "type": "string"
                }
            }
        },
        "github_com_blackhorseya_godine_entity_restaurant_model.MenuItem": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "Description provides details about the menu item.",
                    "type": "string"
                },
                "id": {
                    "description": "ID is the unique identifier of the menu item.",
                    "type": "string"
                },
                "is_available": {
                    "description": "IsAvailable indicates whether the menu item is available.",
                    "type": "boolean"
                },
                "name": {
                    "description": "Name is the name of the menu item.",
                    "type": "string"
                },
                "price": {
                    "description": "Price is the cost of the menu item.",
                    "type": "number"
                }
            }
        },
        "github_com_blackhorseya_godine_entity_restaurant_model.Restaurant": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "Address is the address of the restaurant.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/github_com_blackhorseya_godine_entity_restaurant_model.Address"
                        }
                    ]
                },
                "id": {
                    "description": "ID is the unique identifier of the restaurant.",
                    "type": "string"
                },
                "is_open": {
                    "description": "IsOpen indicates whether the restaurant is open for business.",
                    "type": "boolean"
                },
                "menu": {
                    "description": "Menu is the list of menu items available in the restaurant.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_blackhorseya_godine_entity_restaurant_model.MenuItem"
                    }
                },
                "name": {
                    "description": "Name is the name of the restaurant.",
                    "type": "string"
                }
            }
        },
        "menu.PostPayload": {
            "type": "object",
            "required": [
                "name",
                "price"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "example": ""
                },
                "name": {
                    "type": "string",
                    "example": "menu item name"
                },
                "price": {
                    "type": "number",
                    "example": 10
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
        },
        "restaurants.PostPayload": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "example": "restaurant name"
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

// SwaggerInforestaurant_restful holds exported Swagger Info so clients can modify it
var SwaggerInforestaurant_restful = &swag.Spec{
	Version:          "0.1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Godine Restaurant Restful API",
	Description:      "Godine Restaurant Restful API document.",
	InfoInstanceName: "restaurant_restful",
	SwaggerTemplate:  docTemplaterestaurant_restful,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInforestaurant_restful.InstanceName(), SwaggerInforestaurant_restful)
}
