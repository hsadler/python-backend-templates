// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/api/items": {
            "get": {
                "description": "Returns Items by ids",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Get Items",
                "parameters": [
                    {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        },
                        "collectionFormat": "csv",
                        "description": "Item IDs",
                        "name": "item_ids",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.GetItemsResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates Item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Create Item",
                "parameters": [
                    {
                        "description": "Create Item Request",
                        "name": "createItemRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.CreateItemRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.CreateItemResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/items/{id}": {
            "get": {
                "description": "Returns Item by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "get Item by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.GetItemResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/status": {
            "get": {
                "description": "Returns ` + "`" + `\"ok!\"` + "`" + ` if the server is up",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "status"
                ],
                "summary": "status endpoint",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.statusResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.CreateItemRequest": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/main.ItemIn"
                }
            }
        },
        "main.CreateItemResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/main.Item"
                },
                "meta": {
                    "$ref": "#/definitions/main.CreateItemResponseMeta"
                }
            }
        },
        "main.CreateItemResponseMeta": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "boolean"
                }
            }
        },
        "main.GetItemResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/main.Item"
                },
                "meta": {
                    "type": "object"
                }
            }
        },
        "main.GetItemsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/main.Item"
                    }
                },
                "meta": {
                    "type": "object"
                }
            }
        },
        "main.Item": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "format": "date-time",
                    "example": "2021-01-01T00:00:00.000Z"
                },
                "id": {
                    "type": "integer",
                    "format": "int64",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "format": "string",
                    "example": "foo"
                },
                "price": {
                    "type": "number",
                    "format": "float64",
                    "example": 3.14
                },
                "uuid": {
                    "type": "string",
                    "format": "uuid",
                    "example": "550e8400-e29b-41d4-a716-446655440000"
                }
            }
        },
        "main.ItemIn": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "format": "string",
                    "example": "foo"
                },
                "price": {
                    "type": "number",
                    "format": "float64",
                    "minimum": 0,
                    "example": 3.14
                }
            }
        },
        "main.statusResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "ok!"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Example Server API",
	Description:      "Example Go+Gin+pgx JSON API server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
