package iam_router

import (
	iam_service "github.com/DigiConvent/testd9t/pkg/iam/service"
	"github.com/gin-gonic/gin"
)

type IamRouterInterface interface {
	LoginWithTelegram(ctx *gin.Context)
	LoginWithCredentials(ctx *gin.Context)
}

type IamRouter struct {
	iamService iam_service.IAMServiceInterface
}

func NewIamRouter(iamService iam_service.IAMServiceInterface) IamRouterInterface {
	return &IamRouter{
		iamService: iamService,
	}
}
