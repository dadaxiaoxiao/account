package dao

import "context"

type AccountDAO interface {
	AddActivities(ctx context.Context, activities ...AccountActivity) error
}

// Account 账单本体
type Account struct {
	Id int64 `gorm:"primaryKey,autoIncrement" bson:"id,omitempty"`
	// 对应的用户的 ID, 如果是系统账号，它是 0
	// 一个用户会有多个账号
	Uid int64 `gorm:"uniqueIndex:account_uid"`
	// 对外使用的账号
	Account int64 `gorm:"uniqueIndex:account_uid"`
	Type    uint8 `gorm:"uniqueIndex:account_uid"`

	// 这里可以冗余额外字段

	// 可用余额
	Balance  int64
	Currency string

	Ctime int64
	Utime int64
}

// AccountActivity 账单流水
//
// 记录每一次账单的变动情况
type AccountActivity struct {
	Id  int64 `gorm:"primaryKey,autoIncrement" bson:"id,omitempty"`
	Uid int64 `gorm:"index:account_uid"`
	// 有些设计会只用一个单独的 txn_id 来标记
	Biz   string
	BizId int64
	// account 对外使用账号
	// 一个用户会有多个账号
	Account     int64 `gorm:"index:account_uid"`
	AccountType uint8 `gorm:"index:account_uid"`
	// 调整的金额 正数为+
	Amount int64
	// 币种
	Currency string

	Ctime int64
	Utime int64
}
