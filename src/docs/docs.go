// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-01-11 15:15:59.323834005 +0800 CST m=+0.250176746

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "contact": {},
        "license": {}
    },
    "paths": {
        "/api/articles/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get Article List",
                "parameters": [
                    {},
                    {}
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":[{\"id\":1,\"title\":\"abc\"}]}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/articles/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get article",
                "parameters": [
                    {}
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{\"id\":1,\"name\":\"\"}]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/user/add": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Add User",
                "parameters": [
                    {},
                    {},
                    {}
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":1}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/user/delete": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Delete User",
                "parameters": [
                    {}
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200, \"msg\": \"????????????\", \"data\":\"\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/user/list": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get User List",
                "parameters": [
                    {},
                    {}
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{\"id\":1,\"name\":\"Mantis\", \"email\":\"tangchunlinit@gmail.com\"}]",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
