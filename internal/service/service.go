package service

import (
	"github.com/zxylon/olaf-layout-advanced/internal/repository"
	"github.com/zxylon/olaf-layout-advanced/pkg/jwt"
	"github.com/zxylon/olaf-layout-advanced/pkg/log"
	"github.com/zxylon/olaf-layout-advanced/pkg/sid"
)

type Service struct {
	logger *log.Logger
	sid    *sid.Sid
	jwt    *jwt.JWT
	tm     repository.Transaction
}

func NewService(
	tm repository.Transaction,
	logger *log.Logger,
	sid *sid.Sid,
	jwt *jwt.JWT,
) *Service {
	return &Service{
		logger: logger,
		sid:    sid,
		jwt:    jwt,
		tm:     tm,
	}
}
