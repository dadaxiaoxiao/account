package repository

import (
	"context"
	"github.com/dadaxiaoxiao/account/internal/domain"
	"github.com/dadaxiaoxiao/account/internal/repository/dao"
	"time"
)

type accountRepository struct {
	dao dao.AccountDAO
}

func NewAccountRepository(dao dao.AccountDAO) AccountRepository {
	return &accountRepository{
		dao: dao,
	}
}

func (a *accountRepository) AddCredit(ctx context.Context, c domain.Credit) error {
	activities := make([]dao.AccountActivity, len(c.Items))
	now := time.Now().UnixMilli()
	for _, item := range c.Items {
		activities = append(activities, dao.AccountActivity{
			Uid:         item.Uid,
			Biz:         c.Biz,
			BizId:       c.BizId,
			Account:     item.Account,
			AccountType: item.AccountType.AsUint8(),
			Amount:      item.Amt,
			Currency:    item.Currency,
			Ctime:       now,
			Utime:       now,
		})
	}
	// 添加流水记录
	return a.dao.AddActivities(ctx, activities...)
}
