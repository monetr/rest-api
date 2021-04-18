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
        "contact": {
            "name": "Support",
            "url": "http://github.com/harderthanitneedstobe/rest-api",
            "email": "no-reply@harderthanitneedstobe.com"
        },
        "license": {
            "name": "Business Source License 1.1",
            "url": "https://github.com/HarderThanItNeedsToBe/rest-api/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/authentication/login": {
            "post": {
                "description": "Authenticate a user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Login",
                "operationId": "login",
                "responses": {
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.ApiError"
                        }
                    }
                }
            }
        },
        "/bank_accounts": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Lists all of the bank accounts for the currently authenticated user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bank Accounts"
                ],
                "summary": "List All Bank Accounts",
                "operationId": "list-all-bank-accounts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.BankAccount"
                            }
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
                "description": "Create a bank account for the provided link. Note: Bank accounts can only be created this way for manual links. Attempting to create a bank account for a link that is associated with Plaid will result in an error.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bank Accounts"
                ],
                "summary": "Create Bank Account",
                "operationId": "create-bank-account",
                "parameters": [
                    {
                        "description": "New Bank Account",
                        "name": "newBankAccount",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/swag.BankAccountCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.BankAccount"
                        }
                    }
                }
            }
        },
        "/bank_accounts/{bankAccountId}/balances": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get the balances for the specified bank account (including calculated balances).",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Bank Accounts"
                ],
                "summary": "Get Bank Account Balances",
                "operationId": "get-bank-account-balances",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bank Account ID",
                        "name": "bankAccountId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/repository.Balances"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.InvalidBankAccountIdError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.ApiError"
                        }
                    }
                }
            }
        },
        "/bank_accounts/{bankAccountId}/funding_schedules": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "List all of the funding schedule's for the current bank account.",
                "tags": [
                    "Funding Schedules"
                ],
                "summary": "List Funding Schedules",
                "operationId": "list-funding-schedules",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bank Account ID",
                        "name": "bankAccountId",
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
                                "$ref": "#/definitions/models.FundingSchedule"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.InvalidBankAccountIdError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.ApiError"
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
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Funding Schedules"
                ],
                "summary": "Create a funding schedule for the specified bank account.",
                "operationId": "create-funding-schedule",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bank Account ID",
                        "name": "bankAccountId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "New Funding Schedule",
                        "name": "fundingSchedule",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.FundingSchedule"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.FundingSchedule"
                        }
                    },
                    "400": {
                        "description": "Malformed JSON or invalid RRule.",
                        "schema": {
                            "$ref": "#/definitions/controller.ApiError"
                        }
                    },
                    "500": {
                        "description": "Failed to persist data.",
                        "schema": {
                            "$ref": "#/definitions/controller.ApiError"
                        }
                    }
                }
            }
        },
        "/bank_accounts/{bankAccountId}/spending": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "List all of the spending for the specified bank account.",
                "tags": [
                    "Spending"
                ],
                "operationId": "list-spending",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bank Account ID",
                        "name": "bankAccountId",
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
                                "$ref": "#/definitions/models.Spending"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.InvalidBankAccountIdError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.ApiError"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Spending"
                ],
                "summary": "Update an existing expense or goal spending object.",
                "operationId": "update-spending",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bank Account ID",
                        "name": "bankAccountId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated spending",
                        "name": "Spending",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Spending"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Spending"
                        }
                    },
                    "400": {
                        "description": "Malformed JSON or invalid RRule.",
                        "schema": {
                            "$ref": "#/definitions/controller.ApiError"
                        }
                    },
                    "500": {
                        "description": "Failed to persist data.",
                        "schema": {
                            "$ref": "#/definitions/controller.ApiError"
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
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Spending"
                ],
                "summary": "Create an spending for the specified bank account.",
                "operationId": "create-spending",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bank Account ID",
                        "name": "bankAccountId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "New spending",
                        "name": "Spending",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Spending"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Spending"
                        }
                    },
                    "400": {
                        "description": "Malformed JSON or invalid RRule.",
                        "schema": {
                            "$ref": "#/definitions/controller.ApiError"
                        }
                    },
                    "500": {
                        "description": "Failed to persist data.",
                        "schema": {
                            "$ref": "#/definitions/controller.ApiError"
                        }
                    }
                }
            }
        },
        "/bank_accounts/{bankAccountId}/spending/transfer": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Spending"
                ],
                "summary": "Transfer allocated funds to or from a spending object.",
                "operationId": "transfer-spending",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bank Account ID",
                        "name": "bankAccountId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Transfer",
                        "name": "Spending",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.SpendingTransfer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Spending"
                            }
                        }
                    },
                    "400": {
                        "description": "Malformed JSON or invalid RRule.",
                        "schema": {
                            "$ref": "#/definitions/controller.ApiError"
                        }
                    },
                    "500": {
                        "description": "Failed to persist data.",
                        "schema": {
                            "$ref": "#/definitions/controller.ApiError"
                        }
                    }
                }
            }
        },
        "/config": {
            "get": {
                "description": "Provides the configuration that should be used by the frontend application or UI.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Config"
                ],
                "summary": "Get Config",
                "operationId": "app-config",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/links": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Lists all of the links for the currently authenticated user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Links"
                ],
                "summary": "List All Links",
                "operationId": "list-all-links",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Link"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.ApiError"
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
                "description": "Create a manual link.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Links"
                ],
                "summary": "Create A Link",
                "operationId": "create-link",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Link"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.ApiError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ApiError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "controller.InvalidBankAccountIdError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "invalid bank account Id provided"
                }
            }
        },
        "controller.SpendingTransfer": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "fromSpendingId": {
                    "type": "integer"
                },
                "toSpendingId": {
                    "type": "integer"
                }
            }
        },
        "models.Account": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "integer"
                },
                "timezone": {
                    "type": "string"
                }
            }
        },
        "models.BankAccount": {
            "type": "object",
            "properties": {
                "accountSubType": {
                    "type": "string",
                    "example": "checking"
                },
                "accountType": {
                    "type": "string",
                    "example": "depository"
                },
                "availableBalance": {
                    "type": "integer",
                    "example": 102356
                },
                "bankAccountId": {
                    "type": "integer",
                    "example": 1234
                },
                "currentBalance": {
                    "description": "Current Balance is a 64-bit representation of a bank account's total balance (excluding pending transactions) in\nthe form of an integer. To derive a decimal value divide this value by 100.",
                    "type": "integer",
                    "example": 102400
                },
                "linkId": {
                    "type": "integer",
                    "example": 2345
                },
                "mask": {
                    "type": "string",
                    "example": "0000"
                },
                "name": {
                    "type": "string",
                    "example": "Checking Account"
                },
                "officialName": {
                    "type": "string",
                    "example": "US Bank - Checking Account"
                },
                "originalName": {
                    "type": "string",
                    "example": "Checking Account #1"
                }
            }
        },
        "models.EmailVerification": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "emailAddress": {
                    "type": "string"
                },
                "expiresAt": {
                    "type": "string"
                },
                "isVerified": {
                    "type": "boolean"
                },
                "verifiedAt": {
                    "type": "string"
                }
            }
        },
        "models.FundingSchedule": {
            "type": "object",
            "properties": {
                "bankAccountId": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "fundingScheduleId": {
                    "type": "integer"
                },
                "lastOccurrence": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nextOccurrence": {
                    "type": "string"
                },
                "rule": {
                    "type": "string",
                    "example": "FREQ=MONTHLY;BYMONTHDAY=15,-1"
                }
            }
        },
        "models.Link": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "createdByUserId": {
                    "type": "integer"
                },
                "customInstitutionName": {
                    "type": "string"
                },
                "institutionName": {
                    "type": "string"
                },
                "linkId": {
                    "type": "integer"
                },
                "linkType": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "updatedByUser": {
                    "$ref": "#/definitions/models.User"
                },
                "updatedByUserId": {
                    "type": "integer"
                }
            }
        },
        "models.Login": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "emailVerifications": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.EmailVerification"
                    }
                },
                "isEmailVerified": {
                    "type": "boolean"
                },
                "isPhoneVerified": {
                    "type": "boolean"
                },
                "loginId": {
                    "type": "integer"
                },
                "phoneVerifications": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.PhoneVerification"
                    }
                }
            }
        },
        "models.PhoneNumber": {
            "type": "object"
        },
        "models.PhoneVerification": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "expiresAt": {
                    "type": "string"
                },
                "isVerified": {
                    "type": "boolean"
                },
                "phoneNumber": {
                    "$ref": "#/definitions/models.PhoneNumber"
                },
                "verifiedAt": {
                    "type": "string"
                }
            }
        },
        "models.Spending": {
            "type": "object",
            "properties": {
                "bankAccountId": {
                    "type": "integer"
                },
                "currentAmount": {
                    "type": "integer"
                },
                "dateCreated": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "fundingScheduleId": {
                    "type": "integer"
                },
                "isBehind": {
                    "type": "boolean"
                },
                "isPaused": {
                    "type": "boolean"
                },
                "lastRecurrence": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nextContributionAmount": {
                    "type": "integer"
                },
                "nextRecurrence": {
                    "type": "string"
                },
                "recurrenceRule": {
                    "type": "string"
                },
                "spendingId": {
                    "type": "integer"
                },
                "spendingType": {
                    "type": "integer"
                },
                "targetAmount": {
                    "type": "integer"
                },
                "usedAmount": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "account": {
                    "$ref": "#/definitions/models.Account"
                },
                "accountId": {
                    "type": "integer"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "login": {
                    "$ref": "#/definitions/models.Login"
                },
                "loginId": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "repository.Balances": {
            "type": "object",
            "properties": {
                "available": {
                    "type": "integer"
                },
                "bankAccountId": {
                    "type": "integer"
                },
                "current": {
                    "type": "integer"
                },
                "expenses": {
                    "type": "integer"
                },
                "goals": {
                    "type": "integer"
                },
                "safe": {
                    "type": "integer"
                }
            }
        },
        "swag.BankAccountCreateRequest": {
            "type": "object",
            "properties": {
                "accountSubType": {
                    "description": "Sub Type can have numerous values, but given that the application currently only supports depository the most\ncommon values you will see or use are; checking and savings. Other supported types (albeit untested) are; hsa,\ncd, money market, paypal, prepaid, cash management and ebt.\nMore information on these can be found here: https://plaid.com/docs/api/accounts/#account-type-schema",
                    "type": "string",
                    "example": "checking"
                },
                "accountType": {
                    "description": "Account Type can be; depository, credit, loan, investment or other. At the time of writing this the application\nwill only support depository. Other types may be supported in the future.",
                    "type": "string",
                    "example": "depository"
                },
                "availableBalance": {
                    "description": "The balance available in the account represented as whole cents. This is typically the current balance minus the\ntotal value of all pending transactions. This value is not calculated in the API and is retrieved from Plaid or\nmaintained manually for manual links.",
                    "type": "integer",
                    "example": 102356
                },
                "currentBalance": {
                    "description": "The current balance in the account as whole cents without taking into consideration any pending transactions.",
                    "type": "integer",
                    "example": 102400
                },
                "linkId": {
                    "description": "The numeric Id of the Link this bank account is associated with, if the link is manual then bank bank accounts\ncan be created manually via the API. If the Link is associated with Plaid though then bank accounts can only be\ncreated through the Plaid interface. At the time of writing this there is not a way to add or remove a bank\naccount from an existing Plaid Link.",
                    "type": "integer",
                    "example": 2345
                },
                "mask": {
                    "description": "Last 4 digits of the bank account's account number. We do not store the full bank account number or any other\nsensitive account information.",
                    "type": "string",
                    "example": "9876"
                },
                "name": {
                    "description": "Name of the account, this is different than the ` + "`" + `originalName` + "`" + `. This field can be changed later on while the\n` + "`" + `originalName` + "`" + ` field cannot be changed once the account is created.",
                    "type": "string",
                    "example": "Checking Account"
                },
                "officialName": {
                    "description": "Official name is only used with bank accounts coming from Plaid. It is another name that Plaid uses for an\naccount.",
                    "type": "string",
                    "example": "US Bank - Checking Account"
                },
                "originalName": {
                    "description": "The original name of the bank account from when it was created. This name cannot be changed after the bank\naccount is created. This is primarily due to bank account's coming from a 3rd party provider like Plaid. But to\nreduce the amount of logic in the application the same rule applies for manual links as well.",
                    "type": "string",
                    "example": "Checking Account #1"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "H-Token",
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
	Version:     "0.0",
	Host:        "api.harderthanitneedstobe.com",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Harder Than It Needs To Be's REST API",
	Description: "This is the REST API for our budgeting application.",
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
