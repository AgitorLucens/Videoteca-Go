package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler defines the contract for all handler types
// This implements the Interface Segregation Principle by defining a single, focused interface
// and the Dependency Inversion Principle by depending on abstractions rather than concrete types
type Handler interface {
	Handle(*gin.Context)
}

// Response represents a standardized API response structure
// This implements the Single Responsibility Principle by separating response formatting from business logic
type Response struct {
	Data interface{} `json:"data"`
	Err  string      `json:"error,omitempty"`
}

// NewResponse creates a new response with data
func NewResponse(data interface{}) Response {
	return Response{Data: data}
}

// NewErrorResponse creates a new response with error
func NewErrorResponse(err string) Response {
	return Response{Err: err}
}

// WriteJSON writes a standardized JSON response
func WriteJSON(c *gin.Context, statusCode int, resp Response) {
	c.JSON(statusCode, resp)
}

// WriteSuccess writes a success response with data
func WriteSuccess(c *gin.Context, data interface{}) {
	WriteJSON(c, http.StatusOK, NewResponse(data))
}

// WriteError writes an error response
func WriteError(c *gin.Context, statusCode int, err string) {
	WriteJSON(c, statusCode, NewErrorResponse(err))
}

// WriteBadRequest writes a 400 Bad Request response
func WriteBadRequest(c *gin.Context, err string) {
	WriteError(c, http.StatusBadRequest, err)
}

// WriteInternalServerError writes a 500 Internal Server Error response
func WriteInternalServerError(c *gin.Context, err string) {
	WriteError(c, http.StatusInternalServerError, err)
}
