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
        "contact": {
            "name": "Caio Sousa",
            "url": "http://www.swagger.io/support",
            "email": "caiosousafernandesferreira@hotmail.com"
        },
        "license": {
            "name": "Mozilla Public License 2.0",
            "url": "https://www.mozilla.org/en-US/MPL/2.0/"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/projetos": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Get list all project",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Get All Projects",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/projetos.ReqProjetos"
                            }
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "POST a new project. For the request to be met, the \"nome_projeto\", \"equipe_id\", \"descricao_projeto\" are required. The status already goes with a predefined value \"A Fazer\". the \"prazo_entrega\" is the number of days that the delivery time will be",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "POST a new Project",
                "parameters": [
                    {
                        "description": "NewProject",
                        "name": "NewProject",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/projetos.ReqProjeto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/errorstratment.ResError"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/errorstratment.ResError"
                            }
                        }
                    }
                }
            }
        },
        "/projetos/status/{status}": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "GET all registered projects that have the status passed as a parameter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Get Status of Projects with a specific status with Param Status",
                "parameters": [
                    {
                        "enum": [
                            "A Fazer",
                            "Em Andamento",
                            "Em Teste",
                            "Concluido"
                        ],
                        "type": "string",
                        "description": "Status",
                        "name": "status",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/projetos.ReqStatusProjeto"
                            }
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/projetos/{id}": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "GET a project with a specific ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Get Project with specific ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Projeto ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/projetos.ReqProjetos"
                            }
                        }
                    },
                    "400": {
                        "description": "Project does not exist",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "not authorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "PUT a specific project. For the request to be met, the \"nome_projeto\" and \"equipe_id\" and \"descricao_projeto\" are required",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "PUT Project with ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Projeto ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Project",
                        "name": "Project",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/projetos.ReqAtualizarProjetoData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/errorstratment.ResError"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/errorstratment.ResError"
                            }
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "DELETE a specific project. For the request to be met, the \"id_projeto\" are required",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Delete a specific Project",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Projeto ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/errorstratment.ResOk"
                        }
                    },
                    "401": {
                        "description": "Not Authorized",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/errorstratment.ResError"
                            }
                        }
                    }
                }
            }
        },
        "/projetos/{id}/status": {
            "put": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "PUT Status of a specific project. For the request to be met, the \"status\" are required",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "PUT Status of a Project",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Status-Project",
                        "name": "Status-Project",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/projetos.ReqUpdateStatusProjeto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/errorstratment.ResError"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/errorstratment.ResError"
                            }
                        }
                    }
                }
            }
        },
        "/projetos/{id}/tasks": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "GET all tasks of a project with ID_Projeto specific",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Get Tasks of Project with Param ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Projeto ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/projetos.ReqTasksProjeto"
                            }
                        }
                    },
                    "401": {
                        "description": "not authorized",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Project does not exist",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errorstratment.ResError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "errorstratment.ResOk": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "message": {
                    "type": "string",
                    "example": "Project deleted successfully"
                }
            }
        },
        "projetos.ReqAtualizarProjetoData": {
            "type": "object",
            "properties": {
                "descricao_projeto": {
                    "type": "string",
                    "example": "Criacao de sistema e-commerce"
                },
                "equipe_id": {
                    "type": "integer",
                    "example": 1
                },
                "nome_projeto": {
                    "type": "string",
                    "example": "Casas Bahias"
                }
            }
        },
        "projetos.ReqProjeto": {
            "type": "object",
            "properties": {
                "descricao_projeto": {
                    "type": "string",
                    "example": "Descricao"
                },
                "equipe_id": {
                    "type": "integer",
                    "example": 2
                },
                "nome_projeto": {
                    "type": "string",
                    "example": "Nome"
                },
                "prazo_entrega": {
                    "type": "integer",
                    "example": 2
                }
            }
        },
        "projetos.ReqProjetos": {
            "type": "object",
            "properties": {
                "data_conclusao": {
                    "type": "string",
                    "example": ""
                },
                "data_criacao": {
                    "type": "string",
                    "example": "2022-07-25"
                },
                "descricao_projeto": {
                    "type": "string",
                    "example": "Descricao"
                },
                "equipe_id": {
                    "type": "integer",
                    "example": 2
                },
                "id_projeto": {
                    "type": "integer",
                    "example": 58
                },
                "nome_equipe": {
                    "type": "string",
                    "example": "Cariri Inovação"
                },
                "nome_projeto": {
                    "type": "string",
                    "example": "Nome"
                },
                "prazo_entrega": {
                    "type": "string",
                    "example": "2022-07-25"
                },
                "status": {
                    "type": "string",
                    "example": "Concluido"
                }
            }
        },
        "projetos.ReqStatusProjeto": {
            "type": "object",
            "properties": {
                "data_conclusao": {
                    "type": "string",
                    "example": ""
                },
                "data_criacao": {
                    "type": "string",
                    "example": "2022-07-25"
                },
                "descricao_projeto": {
                    "type": "string"
                },
                "equipe_id": {
                    "type": "integer"
                },
                "id_projeto": {
                    "type": "integer"
                },
                "nome_projeto": {
                    "type": "string"
                },
                "prazo_entrega": {
                    "type": "string",
                    "example": "2022-07-25"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "projetos.ReqTasksProjeto": {
            "type": "object",
            "properties": {
                "data_conclusao": {
                    "type": "string"
                },
                "data_criacao": {
                    "type": "string"
                },
                "descricao_task": {
                    "type": "string"
                },
                "id_projeto": {
                    "type": "integer"
                },
                "id_task": {
                    "type": "integer"
                },
                "nome_equipe": {
                    "type": "string"
                },
                "nome_pessoa": {
                    "type": "string"
                },
                "nome_projeto": {
                    "type": "string"
                },
                "pessoa_id": {
                    "type": "integer"
                },
                "prazo_entrega": {
                    "type": "string"
                },
                "prioridade": {
                    "type": "integer"
                },
                "projeto_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "projetos.ReqUpdateStatusProjeto": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "Em Andamento"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
