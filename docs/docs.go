// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/sign-in": {
            "post": {
                "description": "Signs in",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "SignIn",
                "parameters": [
                    {
                        "description": "SignIn info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/inventory": {
            "get": {
                "description": "Show inventory list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Inventory"
                ],
                "summary": "GetAllInventory",
                "responses": {}
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Adds a new item in inventory",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Inventory"
                ],
                "summary": "AddNewInventory",
                "parameters": [
                    {
                        "description": "item info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Inventory"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete item in inventory",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Inventory"
                ],
                "summary": "DeleteInventory",
                "parameters": [
                    {
                        "description": "item info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Inventory"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/sign-up": {
            "post": {
                "description": "Signs up",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "SignUp",
                "parameters": [
                    {
                        "description": "User info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "entity.Inventory": {
            "type": "object",
            "properties": {
                "dep_name": {
                    "type": "string"
                },
                "fk_dep_id": {
                    "type": "integer"
                },
                "full_name": {
                    "type": "string"
                },
                "invent_laptop": {
                    "type": "string"
                },
                "invent_mfu": {
                    "type": "string"
                },
                "invent_monitors": {
                    "type": "string"
                },
                "invent_other": {
                    "type": "string"
                },
                "invent_printer": {
                    "type": "string"
                },
                "inventory_id": {
                    "type": "integer"
                },
                "inventory_num": {
                    "type": "string"
                },
                "pc_id": {
                    "type": "string"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7540",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Inventory Service API",
	Description:      "API server for inventory service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
