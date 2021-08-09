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
            "email": "felixanthony1996.fa@gmail.com"
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
        "/v1/accounts": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Find All Accounts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Account"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Create account",
                "parameters": [
                    {
                        "description": "Create account",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Account"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.Account"
                        }
                    }
                }
            }
        },
        "/v1/accounts/{account_id}/balance": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Find Balance Account By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "account_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Account"
                        }
                    }
                }
            }
        },
        "/v1/accounts/{id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Delete Account By ID",
                "parameters": [
                    {
                        "type": "string",
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
                            "$ref": "#/definitions/domain.Account"
                        }
                    }
                }
            }
        },
        "/v1/charity-mrys": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create Charity Mrys",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CharityMrys"
                ],
                "summary": "Create Charity Mrys",
                "parameters": [
                    {
                        "description": "Create charity mrys",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CharityMrys"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.CharityMrys"
                        }
                    }
                }
            }
        },
        "/v1/charity-mrys/create-bulk": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create Bulk Charity Mrys",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CharityMrys"
                ],
                "summary": "Create Bulk Charity Mrys",
                "parameters": [
                    {
                        "description": "Create charity mrys",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecase.CreateBulkCharityMrysInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.CharityMrys"
                            }
                        }
                    }
                }
            }
        },
        "/v1/charity-mrys/list-all": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CharityMrys"
                ],
                "summary": "Find All CharityMrys",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.CharityMrysAll"
                        }
                    }
                }
            }
        },
        "/v1/charity-mrys/list-pagination/{currentPage}/{perPage}/{sort}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CharityMrys"
                ],
                "summary": "Find Pagination CharityMrys",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "CurrentPage",
                        "name": "currentPage",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "PerPage",
                        "name": "perPage",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Sort",
                        "name": "sort",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Search",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.CharityMrysPagination"
                        }
                    }
                }
            }
        },
        "/v1/charity-mrys/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CharityMrys"
                ],
                "summary": "Find One Charity Mrys By ID",
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
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.CharityMrys"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CharityMrys"
                ],
                "summary": "Delete One Charity Mrys By ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Charity Mrys ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update Charity Mrys By ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CharityMrys"
                ],
                "summary": "Update Charity Mrys By ID",
                "parameters": [
                    {
                        "description": "Update charity mrys",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CharityMrys"
                        }
                    },
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.CharityMrys"
                        }
                    }
                }
            }
        },
        "/v1/receipt-lunar": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create Receipt Lunar",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ReceiptLunar"
                ],
                "summary": "Create Receipt Lunar",
                "parameters": [
                    {
                        "description": "Create receipt lunar",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usecase.CreateReceiptLunarInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.ReceiptLunar"
                        }
                    }
                }
            }
        },
        "/v1/receipt-lunar/list-pagination/{currentPage}/{perPage}/{sort}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ReceiptLunar"
                ],
                "summary": "Find Pagination ReceiptLunar",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "CurrentPage",
                        "name": "currentPage",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "PerPage",
                        "name": "perPage",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Sort",
                        "name": "sort",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Search",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.ReceiptLunarPagination"
                        }
                    }
                }
            }
        },
        "/v1/receipt-lunar/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ReceiptLunar"
                ],
                "summary": "Find One ReceiptLunar By ID",
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
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.ReceiptLunar"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Account": {
            "type": "object",
            "properties": {
                "Balance": {
                    "type": "integer",
                    "example": 40000
                },
                "Cpf": {
                    "type": "string",
                    "example": "00.00.111.11"
                },
                "CreatedAt": {
                    "type": "string",
                    "example": "2019-11-09T21:21:46+00:00"
                },
                "id": {
                    "type": "string",
                    "example": "1"
                },
                "name": {
                    "type": "string",
                    "example": "Leo Messi"
                }
            }
        },
        "domain.Branch": {
            "type": "object",
            "properties": {
                "CreatedAt": {
                    "type": "string",
                    "example": "2019-11-09T21:21:46+00:00"
                },
                "Description": {
                    "type": "string",
                    "example": "description"
                },
                "address": {
                    "type": "string",
                    "example": "Pontianak"
                },
                "code": {
                    "type": "string",
                    "example": "PTK"
                },
                "id": {
                    "type": "string",
                    "example": "953c294b-4535-4656-9fd8-be58zzd0402a9a1x"
                },
                "name": {
                    "type": "string",
                    "example": "Pontianak"
                }
            }
        },
        "domain.CharityMrys": {
            "type": "object",
            "properties": {
                "Amount": {
                    "type": "integer",
                    "example": 40000
                },
                "CreatedAt": {
                    "type": "string",
                    "example": "2019-11-09T21:21:46+00:00"
                },
                "Description": {
                    "type": "string",
                    "example": "description"
                },
                "Month": {
                    "type": "integer",
                    "example": 1
                },
                "Userame": {
                    "type": "string",
                    "example": "usename"
                },
                "Year": {
                    "type": "integer",
                    "example": 2021
                },
                "branch": {
                    "$ref": "#/definitions/domain.Branch"
                },
                "id": {
                    "type": "string",
                    "example": "1"
                },
                "name": {
                    "type": "string",
                    "example": "Leo Messi"
                },
                "user_id": {
                    "type": "string",
                    "example": "1"
                }
            }
        },
        "domain.CharityMrysAll": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.CharityMrys"
                    }
                }
            }
        },
        "domain.CharityMrysPagination": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.CharityMrys"
                    }
                },
                "meta": {
                    "$ref": "#/definitions/domain.MetaPagination"
                }
            }
        },
        "domain.MetaPagination": {
            "type": "object",
            "properties": {
                "currentPage": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "perPage": {
                    "type": "integer"
                },
                "sort": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                },
                "totalPage": {
                    "type": "integer"
                }
            }
        },
        "domain.ReceiptLunar": {
            "type": "object",
            "properties": {
                "branch": {
                    "$ref": "#/definitions/domain.Branch"
                },
                "created_at": {
                    "type": "string",
                    "example": "2019-11-09T21:21:46+00:00"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "example": "1"
                },
                "internation_date": {
                    "type": "string",
                    "example": "2021-08-04"
                },
                "lunar_date": {
                    "type": "string",
                    "example": "2021-08-04"
                },
                "receipt_lunar_detail": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.ReceiptLunarDetail"
                    }
                },
                "total": {
                    "type": "integer",
                    "example": 40000
                },
                "user_id": {
                    "type": "string",
                    "example": "1"
                },
                "username": {
                    "type": "string",
                    "example": "username"
                }
            }
        },
        "domain.ReceiptLunarDetail": {
            "type": "object",
            "properties": {
                "CreatedAt": {
                    "type": "string",
                    "example": "2019-11-09T21:21:46+00:00"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "example": "1"
                },
                "name": {
                    "type": "string"
                },
                "receipt_lunar": {
                    "$ref": "#/definitions/domain.ReceiptLunar"
                },
                "total": {
                    "type": "integer",
                    "example": 40000
                }
            }
        },
        "domain.ReceiptLunarPagination": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.ReceiptLunar"
                    }
                },
                "meta": {
                    "$ref": "#/definitions/domain.MetaPagination"
                }
            }
        },
        "usecase.CreateBulkCharityMrysInput": {
            "type": "object",
            "required": [
                "amount",
                "month_from",
                "month_to",
                "name",
                "year"
            ],
            "properties": {
                "amount": {
                    "type": "integer",
                    "example": 40000
                },
                "branch": {
                    "$ref": "#/definitions/domain.Branch"
                },
                "description": {
                    "type": "string",
                    "example": "description"
                },
                "id": {
                    "type": "string",
                    "example": "1"
                },
                "month_from": {
                    "type": "integer",
                    "example": 2
                },
                "month_to": {
                    "type": "integer",
                    "example": 10
                },
                "name": {
                    "type": "string",
                    "example": "Leo Messi"
                },
                "year": {
                    "type": "integer",
                    "example": 2021
                }
            }
        },
        "usecase.CreateReceiptLunarInput": {
            "type": "object",
            "required": [
                "international_date",
                "lunar_date"
            ],
            "properties": {
                "branch": {
                    "$ref": "#/definitions/domain.Branch"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "international_date": {
                    "type": "string"
                },
                "lunar_date": {
                    "type": "string"
                },
                "receipt_lunar_detail": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.ReceiptLunarDetail"
                    }
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
	Version:     "1.0",
	Host:        "localhost:8000/vhry/data/",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Data API",
	Description: "Data API",
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
