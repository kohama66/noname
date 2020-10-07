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
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/beautician": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "美容師情報取得",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodel.BeauticianGet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodel.BeauticianGet"
                        }
                    },
                    "500": {
                        "description": "Something went wrong",
                        "schema": {
                            "$ref": "#/definitions/resource.Error"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "美容師登録",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodel.BeauticianCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodel.BeauticianCreate"
                        }
                    },
                    "500": {
                        "description": "Something went wrong",
                        "schema": {
                            "$ref": "#/definitions/resource.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/beautician/all": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "美容師情報取得",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodel.BeauticianGetAll"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodel.BeauticianGetAll"
                        }
                    },
                    "500": {
                        "description": "Something went wrong",
                        "schema": {
                            "$ref": "#/definitions/resource.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/reservation": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "予約登録",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodel.ReservationCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodel.ReservationCreate"
                        }
                    },
                    "500": {
                        "description": "Something went wrong",
                        "schema": {
                            "$ref": "#/definitions/resource.Error"
                        }
                    }
                }
            }
        },
        "/api/v1/reservation/beautician": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "美容師予約情報取得",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requestmodel.ReservationFindByBeautician"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responsemodel.ReservationFindByBeautician"
                        }
                    },
                    "500": {
                        "description": "Something went wrong",
                        "schema": {
                            "$ref": "#/definitions/resource.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "requestmodel.BeauticianCreate": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                }
            }
        },
        "requestmodel.BeauticianGet": {
            "type": "object"
        },
        "requestmodel.BeauticianGetAll": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "menuRandId": {
                    "type": "string"
                },
                "salonRandId": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "requestmodel.ReservationCreate": {
            "type": "object",
            "properties": {
                "beauticiaId": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "menuId": {
                    "type": "integer"
                },
                "spaceId": {
                    "type": "integer"
                }
            }
        },
        "requestmodel.ReservationFindByBeautician": {
            "type": "object",
            "properties": {
                "offset": {
                    "type": "integer"
                }
            }
        },
        "resource.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "responsemodel.Beautician": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "randId": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "responsemodel.BeauticianCreate": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "randId": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "responsemodel.BeauticianGet": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastName": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "randId": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "responsemodel.BeauticianGetAll": {
            "type": "object",
            "properties": {
                "beauticians": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/responsemodel.Beautician"
                    }
                }
            }
        },
        "responsemodel.Reservation": {
            "type": "object",
            "properties": {
                "beauticiaId": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "guestId": {
                    "type": "integer"
                },
                "menuId": {
                    "type": "integer"
                },
                "spaceId": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "responsemodel.ReservationCreate": {
            "type": "object",
            "properties": {
                "beauticiaId": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "guestId": {
                    "type": "integer"
                },
                "menuId": {
                    "type": "integer"
                },
                "spaceId": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "responsemodel.ReservationFindByBeautician": {
            "type": "object",
            "properties": {
                "reservations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/responsemodel.Reservation"
                    }
                }
            }
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
