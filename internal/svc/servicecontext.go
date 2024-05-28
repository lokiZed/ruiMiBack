package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"ruiMiBack2/internal/config"
	"ruiMiBack2/internal/middleware"
)

type ServiceContext struct {
	Config        config.Config
	JwtMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		JwtMiddleware: middleware.NewJwtMiddleware().Handle,
	}
}
