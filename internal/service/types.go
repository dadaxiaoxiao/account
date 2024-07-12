package service

import (
	"context"
	"github.com/dadaxiaoxiao/account/internal/domain"
)

type AccountService interface {
	Credit(ctx context.Context, credit domain.Credit) error
}
