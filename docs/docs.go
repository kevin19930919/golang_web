// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/api/v1/account": {
            "get": {
                "summary": "get all account record",
                "responses": {
                    "200": {
                        "description": "{\"data\":[{title:title1,complete:1},{title:title2,complete:0}]}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "summary": "add account record",
                "parameters": [
                    {
                        "description": "account",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateAccountModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/account/{email}": {
            "get": {
                "summary": "get account record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"data\":[{email:example@.xxx.com,complete:1}]}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/book": {
            "get": {
                "summary": "get all book record",
                "responses": {
                    "200": {
                        "description": "{\"data\":[{title:title1,id:0},{title:title2,id:1}]}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "summary": "add book record",
                "parameters": [
                    {
                        "description": "book",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateBookModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/book/{id}": {
            "get": {
                "summary": "get book record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"data\":[{title:title,complete:1}]}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/booklist/{list_id}/book/{book_id}": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "summary": "create booklist record",
                "parameters": [
                    {
                        "description": "booklist",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.InsertBooklistInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "summary": "delete booklist record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "book_id",
                        "name": "book_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "summary": "login",
                "parameters": [
                    {
                        "description": "login",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/order": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "summary": "add order record",
                "parameters": [
                    {
                        "description": "book-list",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.CreateOrderByListInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/order/{id}": {
            "patch": {
                "summary": "return order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "order_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/test/not-use-tag": {
            "get": {
                "summary": "check what happen if not use json tag",
                "responses": {}
            }
        },
        "/api/v1/todo": {
            "get": {
                "summary": "get all todo record",
                "responses": {
                    "200": {
                        "description": "{\"data\":[{title:title1,complete:1},{title:title2,complete:0}]}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "summary": "add todo record",
                "parameters": [
                    {
                        "description": "todo",
                        "name": "title",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateTodoModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/todo/{title}": {
            "get": {
                "summary": "get todo record",
                "parameters": [
                    {
                        "type": "string",
                        "description": "title",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"data\":[{title:title,complete:1}]}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CreateAccountModel": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.CreateBookModel": {
            "type": "object",
            "properties": {
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.CreateTodoModel": {
            "type": "object",
            "properties": {
                "completed": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.InsertBooklistInfo": {
            "type": "object",
            "required": [
                "book_id",
                "list_id"
            ],
            "properties": {
                "book_id": {
                    "type": "string"
                },
                "list_id": {
                    "type": "string"
                }
            }
        },
        "model.LoginInfo": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "service.CreateOrderByListInfo": {
            "type": "object",
            "properties": {
                "list_id": {
                    "type": "string"
                },
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/service.OrderDetail"
                    }
                }
            }
        },
        "service.OrderDetail": {
            "type": "object",
            "required": [
                "book_id",
                "end_date"
            ],
            "properties": {
                "book_id": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
