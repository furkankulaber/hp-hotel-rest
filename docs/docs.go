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
        "/auth/login": {
            "post": {
                "description": "Authenticate user with provided credentials and generate JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User logged in successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.APIResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "401": {
                        "description": "Invalid credentials",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    }
                }
            }
        },
        "/auth/protected": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "A protected route that requires authentication",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "responses": {
                    "200": {
                        "description": "Protected Route",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.APIResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register a new user with the provided credentials",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User registered successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.APIResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    }
                }
            }
        },
        "/hotel/reviews/{reviewID}": {
            "put": {
                "description": "Update an existing review",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reviews"
                ],
                "summary": "Update a review",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Review ID",
                        "name": "reviewID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Review request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateReviewRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Review updated successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.APIResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.ReviewResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    }
                }
            }
        },
        "/hotel/{hotelID}/reviews": {
            "get": {
                "description": "Get reviews for a specific hotel",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reviews"
                ],
                "summary": "Get reviews by hotel ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Hotel ID",
                        "name": "hotelID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Reviews fetched successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/utils.APIResponse"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "data": {
                                                "type": "array",
                                                "items": {
                                                    "$ref": "#/definitions/model.ReviewResponse"
                                                }
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid hotel ID",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new review for a specific hotel",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reviews"
                ],
                "summary": "Add a new review",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Hotel ID",
                        "name": "hotelID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Review request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateReviewRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Review added successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.APIResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/model.ReviewResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    }
                }
            }
        },
        "/hotel/{id}": {
            "get": {
                "description": "get detail of hotel by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hotels"
                ],
                "summary": "Get a hotel by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Hotel ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Hotel fetched successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.APIResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.HotelDetailResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid Hotel ID",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    },
                    "404": {
                        "description": "Hotel Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    }
                }
            }
        },
        "/hotels": {
            "get": {
                "description": "get list of all hotels",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "hotels"
                ],
                "summary": "Get all hotels",
                "responses": {
                    "200": {
                        "description": "Hotels fetched successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.APIResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.HotelListResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Failed",
                        "schema": {
                            "$ref": "#/definitions/utils.APIResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CreateReviewRequest": {
            "type": "object",
            "properties": {
                "rating": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                },
                "user_email": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "model.HotelDetailResponse": {
            "type": "object",
            "properties": {
                "amenities": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "location": {
                    "$ref": "#/definitions/model.LocationResponse"
                },
                "name": {
                    "type": "string"
                },
                "photos": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "rating": {
                    "type": "number"
                },
                "reviews": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ReviewResponse"
                    }
                },
                "rooms": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.RoomResponse"
                    }
                },
                "stars": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.HotelListResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "location": {
                    "$ref": "#/definitions/model.LocationResponse"
                },
                "name": {
                    "type": "string"
                },
                "photos": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "rating": {
                    "type": "number"
                },
                "stars": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "model.LocationResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "district": {
                    "type": "string"
                }
            }
        },
        "model.LoginRequest": {
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
        "model.ReviewResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "model.RoomResponse": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "model.UpdateReviewRequest": {
            "type": "object",
            "properties": {
                "rating": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                },
                "user_email": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                }
            }
        },
        "utils.APIResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "\"Please using \u003cb\u003eBearer: JWT\u003c/b\u003e\"",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "HotelPro Hotel REST API",
	Description:      "This is a sample API for managing hotels and reviews.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
