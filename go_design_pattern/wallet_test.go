package go_design_pattern_test

import (
	"testing"
	"time"
)

// 两部分 一部分跟应用内的虚拟钱包账户打交道 一部分跟银行账户打交道
// 虚拟钱包系统和三方支付系统

// 钱包 充值 提现 支付 查询余额 查询交易流水
// 第三方支付 +余额 -余额 +- 余额 查询余额 查询交易流水
// 交易流水    [交易流水ID] [交易时间] [交易金额] [交易类型](充值 提现 支付) [入账钱包账号] [出账钱包账号]

type virtualWallet struct{
	id string
	createTime time.Time
	balance int64
	GetBalance func()int64
	Debit func(int64)int64
	Credit func(int64) int64
}

type VirtualWalletService struct {
	vw *virtualWallet
}

func (v *VirtualWalletService)GetBalance()int64{
	v.vw.GetBalance =  func() int64{
		return 1
	}
	return v.vw.GetBalance()
}

func Test_VirtualWallet(t *testing.T){
	vw := &virtualWallet{}
	vws := &VirtualWalletService{vw}
	t.Log(vws.GetBalance())
}