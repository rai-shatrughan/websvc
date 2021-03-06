// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Define and retrieve Device Info.",
    "title": "API service for Devices",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/api/v1",
  "paths": {
    "/device/version": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Returns current api version.",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/device/{deviceId}": {
      "get": {
        "summary": "Returns a device by ID.",
        "parameters": [
          {
            "minimum": 1,
            "type": "string",
            "format": "uuid",
            "name": "deviceId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A Device object.",
            "schema": {
              "$ref": "#/definitions/device"
            }
          },
          "400": {
            "description": "The specified device ID is invalid (e.g. not a number)."
          },
          "404": {
            "description": "A device with the specified ID was not found."
          },
          "default": {
            "description": "Unexpected error"
          }
        }
      }
    }
  },
  "definitions": {
    "device": {
      "required": [
        "id",
        "name",
        "location",
        "status"
      ],
      "properties": {
        "id": {
          "type": "string",
          "format": "uuid",
          "example": "5ca0f5e4-ac05-4480-9ee0-e896a22b827a"
        },
        "location": {
          "type": "string",
          "example": "Atlantic Bay"
        },
        "name": {
          "type": "string",
          "example": "Assembly device"
        },
        "status": {
          "type": "string",
          "example": "active | offline"
        }
      }
    }
  },
  "securityDefinitions": {
    "apiKey": {
      "type": "apiKey",
      "name": "X-API-Key",
      "in": "header"
    }
  },
  "security": [
    {
      "apiKey": []
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Define and retrieve Device Info.",
    "title": "API service for Devices",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/api/v1",
  "paths": {
    "/device/version": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Returns current api version.",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/device/{deviceId}": {
      "get": {
        "summary": "Returns a device by ID.",
        "parameters": [
          {
            "minimum": 1,
            "type": "string",
            "format": "uuid",
            "name": "deviceId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "A Device object.",
            "schema": {
              "$ref": "#/definitions/device"
            }
          },
          "400": {
            "description": "The specified device ID is invalid (e.g. not a number)."
          },
          "404": {
            "description": "A device with the specified ID was not found."
          },
          "default": {
            "description": "Unexpected error"
          }
        }
      }
    }
  },
  "definitions": {
    "device": {
      "required": [
        "id",
        "name",
        "location",
        "status"
      ],
      "properties": {
        "id": {
          "type": "string",
          "format": "uuid",
          "example": "5ca0f5e4-ac05-4480-9ee0-e896a22b827a"
        },
        "location": {
          "type": "string",
          "example": "Atlantic Bay"
        },
        "name": {
          "type": "string",
          "example": "Assembly device"
        },
        "status": {
          "type": "string",
          "example": "active | offline"
        }
      }
    }
  },
  "securityDefinitions": {
    "apiKey": {
      "type": "apiKey",
      "name": "X-API-Key",
      "in": "header"
    }
  },
  "security": [
    {
      "apiKey": []
    }
  ]
}`))
}
