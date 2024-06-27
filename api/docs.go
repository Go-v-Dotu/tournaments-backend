// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

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
        "/tournaments": {
            "post": {
                "description": "host tournament by authorized player",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tournaments"
                ],
                "summary": "Host Tournament",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization info",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Tournament info",
                        "name": "tournament_info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.TournamentInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/views.HostTournamentResponse"
                        }
                    }
                }
            }
        },
        "/tournaments/{id}": {
            "get": {
                "description": "get tournament",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tournaments"
                ],
                "summary": "Get Tournament",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization info",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID of the tournament",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/views.GetTournamentResponse"
                        }
                    }
                }
            }
        },
        "/tournaments/{id}/players": {
            "get": {
                "description": "enroll a player that isn't a registered user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tournaments",
                    "players"
                ],
                "summary": "Enroll Guest Player",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization info",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID of the tournament",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Guest info",
                        "name": "user_info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.GuestUserInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/views.EnrollGuestPlayerResponse"
                        }
                    }
                }
            }
        },
        "/tournaments/{id}/players/{user_id}": {
            "put": {
                "description": "enroll a player that is a registered user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tournaments",
                    "players"
                ],
                "summary": "Enroll Player",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization info",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID of the tournament",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID of the user to be added",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/views.EnrollPlayerResponse"
                        }
                    }
                }
            }
        },
        "/user/tournaments": {
            "get": {
                "description": "get all tournaments hosted by authorized user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tournaments"
                ],
                "summary": "Hosted Tournaments",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization info",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/views.HostedTournamentsResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.GuestUserInfo": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.TournamentInfo": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "views.EnrollGuestPlayerResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "views.EnrollPlayerResponse": {
            "type": "object"
        },
        "views.GetPlayersResponse": {
            "type": "object",
            "properties": {
                "players": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/views.Player"
                    }
                }
            }
        },
        "views.GetTournamentResponse": {
            "type": "object",
            "properties": {
                "tournament": {
                    "$ref": "#/definitions/views.Tournament"
                }
            }
        },
        "views.HostTournamentResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "views.HostedTournamentsResponse": {
            "type": "object",
            "properties": {
                "tournaments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/views.TournamentPreview"
                    }
                }
            }
        },
        "views.Player": {
            "type": "object",
            "properties": {
                "dropped": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "views.Tournament": {
            "type": "object",
            "properties": {
                "date": {
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
        "views.TournamentPreview": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "total_players": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "127.0.0.1:30001",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Tournament Management Service",
	Description:      "Service for managing lifecycle of the tournaments",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
