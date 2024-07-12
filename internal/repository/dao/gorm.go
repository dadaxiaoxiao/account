package dao

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type AccountGORMDAO struct {
	db *gorm.DB
}

func NewAccountGORMDAO(db *gorm.DB) AccountDAO {
	return &AccountGORMDAO{
		db: db,
	}
}

func (a *AccountGORMDAO) AddActivities(ctx context.Context, activities ...AccountActivity) error {
	// 这里是事务操作
	return a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 修改余额
		// 添加支付记录
		now := time.Now().UnixMilli()
		for _, act := range activities {
			err := tx.Create(&Account{
				Uid:      act.Uid,
				Account:  act.Account,
				Type:     act.AccountType,
				Balance:  act.Account,
				Currency: act.Currency,
				Ctime:    now,
				Utime:    now,
			}).Clauses(clause.OnConflict{
				DoUpdates: clause.Assignments(map[string]any{
					"utime":   now,
					"balance": gorm.Expr("`balance` + ？", act.Amount),
				}),
			}).Error
			if err != nil {
				return err
			}
		}
		// 批量插入
		return tx.Create(activities).Error
	})
}
