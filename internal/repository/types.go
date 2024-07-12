package repository

import (
	"context"
	"github.com/dadaxiaoxiao/account/internal/domain"
)

type AccountRepository interface {
	AddCredit(ctx context.Context, c domain.Credit) error
}
