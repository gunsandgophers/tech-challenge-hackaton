package controllers

import (
	"fmt"
	httpserver "tech-challenge-hackaton/internal/infra/http"
)

type ErrorJSONResponse struct {
	Message   string `json:"message,omitempty"`
	ErrorCode int    `json:"error_code,omitempty"`
}

type JSONResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func sendError(ctx httpserver.HTTPContext, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, ErrorJSONResponse{
		Message:   msg,
		ErrorCode: code,
	})
}

func sendSuccess(ctx httpserver.HTTPContext, code int, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, JSONResponse{
		Message: fmt.Sprintf("operation: %s successfull", op),
		Data:    data,
	})
}
