package router_utils

import (
	api_middleware "github.com/DigiConvent/testd9t/delivery/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetId(ctx *gin.Context) *uuid.UUID {
	id := ctx.Params.ByName("id")
	parsedId := uuid.MustParse(id)
	return &parsedId
}

func GetUserId(ctx *gin.Context) *uuid.UUID {
	id := ctx.GetString(api_middleware.ContextField)
	parsedId := uuid.MustParse(id)
	return &parsedId
}
