package controllers

import (
	"fmt"
	httpserver "tech-challenge-hackaton/internal/infra/http"
)

func sendError(ctx httpserver.HTTPContext, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, httpserver.Payload{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx httpserver.HTTPContext, code int, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, httpserver.Payload{
		"message": fmt.Sprintf("operation: %s successfull", op),
		"data":    data,
	})
}
