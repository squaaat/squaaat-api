// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
	"schemes": {{ marshal .Schemes }},
	"swagger": "2.0",
	"info": {
		"description": "{{.Description}}",
		"title": "{{.Title}}",
		"termsOfService": "http://swagger.io/terms/",
		"contact": {
			"name": "API Support",
			"email": "fiber@swagger.io"
		},
		"license": {
			"name": "Apache 2.0",
			"url": "http://www.apache.org/licenses/LICENSE-2.0.html"
		},
		"version": "{{.Version}}"
	},
	"host": "{{.Host}}",
	"basePath": "{{.BasePath}}",
	"paths": {
		"/api/v1/exercise/search": {
			"get": {
				""
			}
		},
		"/api/v1/auth/login": {
			"post": {
				"description": "login with device id",
				"consumes": [
					"application/json"
				],
				"produces": [
					"application/json"
				],
				"parameters": [
					{
						"in": "body",
						"name": "device_token",
						"description": "Device's token",
						"schema": {
							"type": "object",
							"properties": {
								"device_token": {
									"type": "string"
								}
							},
							"example": {
								"device_token": "helloworld"
							}
						}
					},
				],
				"responses": {
					"200": {
						"description": "OK",
						"headers": {
							"X-Auth-Token": {
								"type": "string",
								"description": "session token for squaaat-api requests"
							}
						}
					}
				}
			}
		},

		"/accounts/{id}": {
			"get": {
				"description": "get string by ID",
				"consumes": [
					"application/json"
				],
				"produces": [
					"application/json"
				],
				"summary": "Show a account",
				"operationId": "get-string-by-int",
				"parameters": [
					{
						"type": "integer",
						"description": "Account ID",
						"name": "id",
						"in": "path",
						"required": true
					}
				],
				"responses": {
					"200": {
						"description": "OK",
						"schema": {
							"$ref": "#/definitions/main.Account"
						}
					},
					"400": {
						"description": "Bad Request",
						"schema": {
							"$ref": "#/definitions/main.HTTPError"
						}
					},
					"404": {
						"description": "Not Found",
						"schema": {
							"$ref": "#/definitions/main.HTTPError"
						}
					},
					"500": {
						"description": "Internal Server Error",
						"schema": {
							"$ref": "#/definitions/main.HTTPError"
						}
					}
				}
			}
		}
	},
	"definitions": {
		"main.Account": {
			"type": "object",
			"properties": {
				"id": {
					"type": "string"
				}
			}
		},
		"main.HTTPError": {
			"type": "object"
		}
	}
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.0.1v",
	Host:        "localhost:4000",
	BasePath:    "/",
	Schemes:     []string{"http", "https"},
	Title:       "SQUAAAT API",
	Description: "SQUAAAT API http server",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}