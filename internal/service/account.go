package service

import (
	"context"
	"github.com/dadaxiaoxiao/account/internal/domain"
	"github.com/dadaxiaoxiao/account/internal/repository"
)

type accountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) AccountService {
	return &accountService{
		repo: repo,
	}
}

func (a *accountService) Credit(ctx context.Context, credit domain.Credit) error {
	return a.repo.AddCredit(ctx, credit)
}
