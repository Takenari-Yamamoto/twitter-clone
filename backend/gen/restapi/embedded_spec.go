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
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Twitter Clone API仕様",
    "title": "Twitter Clone API",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/api/v1",
  "paths": {
    "/auth/login": {
      "post": {
        "tags": [
          "auth"
        ],
        "summary": "ログイン",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "email",
                "password"
              ],
              "properties": {
                "email": {
                  "type": "string"
                },
                "password": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ログイン成功",
            "schema": {
              "type": "object",
              "properties": {
                "token": {
                  "type": "string"
                },
                "user": {
                  "$ref": "#/definitions/User"
                }
              }
            }
          },
          "401": {
            "description": "認証エラー",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/auth/signup": {
      "post": {
        "tags": [
          "auth"
        ],
        "summary": "ユーザー登録",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "username",
                "email",
                "password"
              ],
              "properties": {
                "email": {
                  "type": "string"
                },
                "password": {
                  "type": "string"
                },
                "username": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "登録成功",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "不正なリクエスト",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/tweets": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "tags": [
          "tweets"
        ],
        "summary": "ツイート一覧取得",
        "parameters": [
          {
            "type": "integer",
            "default": 1,
            "name": "page",
            "in": "query"
          },
          {
            "type": "integer",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "成功",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Tweet"
              }
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "tags": [
          "tweets"
        ],
        "summary": "ツイート投稿",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "content"
              ],
              "properties": {
                "content": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "投稿成功",
            "schema": {
              "$ref": "#/definitions/Tweet"
            }
          },
          "401": {
            "description": "認証エラー",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/tweets/{tweetId}": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "tags": [
          "tweets"
        ],
        "summary": "特定のツイートを取得",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "tweetId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "成功",
            "schema": {
              "$ref": "#/definitions/Tweet"
            }
          },
          "404": {
            "description": "ツイートが見つかりません",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "tags": [
          "tweets"
        ],
        "summary": "ツイートを削除",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "tweetId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "削除成功"
          },
          "401": {
            "description": "認証エラー",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "403": {
            "description": "権限エラー",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/tweets/{tweetId}/like": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "tags": [
          "tweets"
        ],
        "summary": "ツイートにいいねをする",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "tweetId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "いいね成功"
          },
          "401": {
            "description": "認証エラー",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "tags": [
          "tweets"
        ],
        "summary": "いいねを取り消す",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "tweetId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "いいね取り消し成功"
          },
          "401": {
            "description": "認証エラー",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/users/{userId}/tweets": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "tags": [
          "tweets",
          "users"
        ],
        "summary": "特定のユーザーのツイート一覧を取得",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "userId",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "default": 1,
            "name": "page",
            "in": "query"
          },
          {
            "type": "integer",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "成功",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Tweet"
              }
            }
          },
          "404": {
            "description": "ユーザーが見つかりません",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "LoginResponse": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        },
        "user": {
          "$ref": "#/definitions/User"
        }
      }
    },
    "Tweet": {
      "type": "object",
      "properties": {
        "content": {
          "type": "string"
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "updatedAt": {
          "type": "string",
          "format": "date-time"
        },
        "user": {
          "$ref": "#/definitions/User"
        },
        "userId": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "avatarUrl": {
          "type": "string"
        },
        "bio": {
          "type": "string"
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "displayName": {
          "type": "string"
        },
        "email": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "updatedAt": {
          "type": "string",
          "format": "date-time"
        },
        "username": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Twitter Clone API仕様",
    "title": "Twitter Clone API",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/api/v1",
  "paths": {
    "/auth/login": {
      "post": {
        "tags": [
          "auth"
        ],
        "summary": "ログイン",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "email",
                "password"
              ],
              "properties": {
                "email": {
                  "type": "string"
                },
                "password": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ログイン成功",
            "schema": {
              "type": "object",
              "properties": {
                "token": {
                  "type": "string"
                },
                "user": {
                  "$ref": "#/definitions/User"
                }
              }
            }
          },
          "401": {
            "description": "認証エラー",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/auth/signup": {
      "post": {
        "tags": [
          "auth"
        ],
        "summary": "ユーザー登録",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "username",
                "email",
                "password"
              ],
              "properties": {
                "email": {
                  "type": "string"
                },
                "password": {
                  "type": "string"
                },
                "username": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "登録成功",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "不正なリクエスト",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/tweets": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "tags": [
          "tweets"
        ],
        "summary": "ツイート一覧取得",
        "parameters": [
          {
            "type": "integer",
            "default": 1,
            "name": "page",
            "in": "query"
          },
          {
            "type": "integer",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "成功",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Tweet"
              }
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "tags": [
          "tweets"
        ],
        "summary": "ツイート投稿",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "content"
              ],
              "properties": {
                "content": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "投稿成功",
            "schema": {
              "$ref": "#/definitions/Tweet"
            }
          },
          "401": {
            "description": "認証エラー",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/tweets/{tweetId}": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "tags": [
          "tweets"
        ],
        "summary": "特定のツイートを取得",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "tweetId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "成功",
            "schema": {
              "$ref": "#/definitions/Tweet"
            }
          },
          "404": {
            "description": "ツイートが見つかりません",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "tags": [
          "tweets"
        ],
        "summary": "ツイートを削除",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "tweetId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "削除成功"
          },
          "401": {
            "description": "認証エラー",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "403": {
            "description": "権限エラー",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/tweets/{tweetId}/like": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "tags": [
          "tweets"
        ],
        "summary": "ツイートにいいねをする",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "tweetId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "いいね成功"
          },
          "401": {
            "description": "認証エラー",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "tags": [
          "tweets"
        ],
        "summary": "いいねを取り消す",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "tweetId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "いいね取り消し成功"
          },
          "401": {
            "description": "認証エラー",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/users/{userId}/tweets": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "tags": [
          "tweets",
          "users"
        ],
        "summary": "特定のユーザーのツイート一覧を取得",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "userId",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "default": 1,
            "name": "page",
            "in": "query"
          },
          {
            "type": "integer",
            "default": 20,
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "成功",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Tweet"
              }
            }
          },
          "404": {
            "description": "ユーザーが見つかりません",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "LoginResponse": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        },
        "user": {
          "$ref": "#/definitions/User"
        }
      }
    },
    "Tweet": {
      "type": "object",
      "properties": {
        "content": {
          "type": "string"
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "updatedAt": {
          "type": "string",
          "format": "date-time"
        },
        "user": {
          "$ref": "#/definitions/User"
        },
        "userId": {
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "avatarUrl": {
          "type": "string"
        },
        "bio": {
          "type": "string"
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "displayName": {
          "type": "string"
        },
        "email": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "updatedAt": {
          "type": "string",
          "format": "date-time"
        },
        "username": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
}
