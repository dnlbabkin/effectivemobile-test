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
        "/addCar": {
            "post": {
                "description": "Добавление новой машины в каталог",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Добавление новой машины",
                "parameters": [
                    {
                        "description": "Данные о новой машине",
                        "name": "newCar",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/effectivemobile-test_internal_dao.NewCar"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Машина успешно добавлена"
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        },
        "/cars": {
            "get": {
                "description": "Получение списка автомобилей с возможностью пагинации",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Получение списка автомобилей с возможностью пагинации",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Номер страницы (по умолчанию 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Максимальное количество записей на странице (по умолчанию 10)",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный запрос"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        },
        "/delete/{id}": {
            "delete": {
                "description": "Удаление записи о машине по идентификатору",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Удаление записи о машине по идентификатору",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор машины",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Запись успешно удалена"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        },
        "/update/{id}": {
            "put": {
                "description": "Обновление данных о машине по идентификатору",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Обновление данных о машине по идентификатору",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор машины",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Обновленные данные о машине",
                        "name": "updatedCar",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/effectivemobile-test_internal_models.Car"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Данные о машине успешно обновлены"
                    },
                    "400": {
                        "description": "Некорректный запрос"
                    },
                    "500": {
                        "description": "Внутренняя ошибка сервера"
                    }
                }
            }
        }
    },
    "definitions": {
        "effectivemobile-test_internal_dao.NewCar": {
            "type": "object",
            "properties": {
                "mark": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "owner": {
                    "type": "object",
                    "properties": {
                        "name": {
                            "type": "string"
                        },
                        "patronymic": {
                            "type": "string"
                        },
                        "surname": {
                            "type": "string"
                        }
                    }
                },
                "regNums": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "effectivemobile-test_internal_models.Car": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "mark": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "owner": {
                    "$ref": "#/definitions/effectivemobile-test_internal_models.People"
                },
                "owner_id": {
                    "type": "integer"
                },
                "regNums": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "effectivemobile-test_internal_models.People": {
            "type": "object",
            "properties": {
                "cars": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/effectivemobile-test_internal_models.Car"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Каталог новых автомобилей",
	Description:      "API для каталога новых автомобилей",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
