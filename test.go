package main

type PayLog struct {
	ID        uint
	Amount    float64 // 充值金额
	FromType  int     // 1 现金 2 公司转账  3 微信 4 支付包
	Account   Account //
	AccountID uint    // 操作人员
}
type Account struct {
	ID        uint
	Company   Company
	CompanyID uint
}
type Company struct {
	ID uint
}
