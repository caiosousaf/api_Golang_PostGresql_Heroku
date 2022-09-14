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
