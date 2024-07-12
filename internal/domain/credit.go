package domain

// Credit 账单
type Credit struct {
	// 业务唯一标识
	Biz   string
	BizId int64
	Items []CreditItem
}

// CreditItem 账单明细
type CreditItem struct {
	// 被打赏人id
	Uid int64
	// 对外暴露的账号（这里是打赏账号）
	// 一个人会存在多个账号
	Account int64
	// 账号类型
	AccountType AccountType
	Amt         int64
	// 货币类型
	Currency string
}

// AccountType 账号类型
type AccountType uint8

func (a AccountType) AsUint8() uint8 {
	return uint8(a)
}

// 账单
// 用户收到打赏
// 平台系统抽成
const (
	AccountTypeUnknown = iota
	// AccountTypeReward 个人赞赏账号
	AccountTypeReward
	// AccountTypeSystem 平台分成账号
	AccountTypeSystem
)
