package middleware

import (
	"github.com/oldme-git/36hour/app/gateway/internal/service"
)

type sMiddleware struct {
}

func init() {
	service.RegisterMiddleware(New())
}

func New() *sMiddleware {
	return &sMiddleware{}
}
