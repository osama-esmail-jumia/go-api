package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
	"unicode"
)

var validationCode = "400"

type Response struct {
	Data    interface{} `json:"data"`
	Links   interface{} `json:"links"`
	Meta    interface{} `json:"meta"`
	Jsonapi interface{} `json:"jsonapi"`
}

type ErrorResponse struct {
	Errors  []Error     `json:"errors"`
	Links   interface{} `json:"links"`
	Meta    interface{} `json:"meta"`
	Jsonapi interface{} `json:"jsonapi"`
}

type Error struct {
	Code    string `json:"code"`
	Source  string `json:"source"`
	Message string `json:"message"`
}

func HandleResponse(c *gin.Context, code int, response Response) {
	c.JSON(code, response)
}

func HandleResponseError(c *gin.Context, err error) {
	switch cErr := err.(type) {
	case validator.ValidationErrors:
		errs := handleValidationErrors(cErr)
		c.JSON(
			http.StatusUnprocessableEntity, ErrorResponse{
				Errors: errs,
			},
		)
	}
}

func handleValidationErrors(errors validator.ValidationErrors) []Error {
	var errorResponse = make([]Error, 0, len(errors))
	
	for _, err := range errors {
		e := validationErrorToText(err)
		errorResponse = append(errorResponse, e)
	}
	
	return errorResponse
}

func validationErrorToText(e validator.FieldError) Error {
	var detail string
	switch e.Tag() {
	case "required":
		detail = fmt.Sprintf("%s is required", getFieldName(e.Field()))
	default:
		detail = "value is not valid"
	}
	return Error{
		Code:    validationCode,
		Source:  e.Field(),
		Message: detail,
	}
}

func getFieldName(name string) string {
	newName := ""
	for _, c := range name {
		if unicode.IsUpper(c) {
			newName += " "
		}
		newName += string(c)
	}
	
	return strings.TrimSpace(newName)
}
