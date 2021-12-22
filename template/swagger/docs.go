// Package swagger GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package swagger

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "konstantin@shamiev.ru"
        },
        "license": {
            "name": "Sample License"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/sun/general/ping": {
            "get": {
                "tags": [
                    "General"
                ],
                "summary": "ping",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/sun/general/test/{id}": {
            "get": {
                "tags": [
                    "General"
                ],
                "summary": "test",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "user",
                        "schema": {
                            "$ref": "#/definitions/mdsun.User"
                        }
                    }
                }
            }
        },
        "/api/sun/general/version": {
            "get": {
                "tags": [
                    "General"
                ],
                "summary": "получение версии приложения",
                "responses": {
                    "200": {
                        "description": "version",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/sun/websocket/gorilla/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "General"
                ],
                "summary": "пример работы с веб-сокетом (http://localhost:8080/template/gorilla/index.html)",
                "responses": {
                    "101": {
                        "description": "Switching Protocols to websocket",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "mdsun.User": {
            "type": "object",
            "properties": {
                "cnt": {
                    "type": "integer"
                },
                "cnt2": {
                    "type": "integer"
                },
                "cnt4": {
                    "type": "integer"
                },
                "cnt8": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05Z"
                },
                "deleted_at": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05Z"
                },
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "number",
                    "example": 0
                },
                "id": {
                    "type": "string",
                    "example": "8ca3c9c3-cf1a-47fe-8723-3f957538ce42"
                },
                "is_online": {
                    "type": "boolean"
                },
                "login": {
                    "type": "string"
                },
                "metrika": {
                    "type": "string",
                    "example": "JSON"
                },
                "price": {
                    "type": "number",
                    "example": 0.01
                },
                "summa_one": {
                    "type": "number",
                    "example": 0.01
                },
                "summa_two": {
                    "type": "number",
                    "example": 0.01
                },
                "updated_at": {
                    "type": "string",
                    "example": "2006-01-02T15:04:05Z"
                }
            }
        }
    },
    "tags": [
        {
            "description": "Общие запросы",
            "name": "General"
        }
    ]
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
	Version:     "1.0.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{"http", "https"},
	Title:       "Sungora API",
	Description: "Sungora",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
