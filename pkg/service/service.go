package service

import "github.com/yerassylbolatov/payment-service/pkg/repository"

type Authorization interface {
}

type Transactions interface {
}

type Service struct {
	Authorization
	Transactions
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
