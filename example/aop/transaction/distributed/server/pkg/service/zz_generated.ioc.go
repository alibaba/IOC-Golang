//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package service

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	util "github.com/alibaba/ioc-golang/autowire/util"
	rpc_service "github.com/alibaba/ioc-golang/extension/autowire/rpc/rpc_service"
)

func init() {
	var bankServiceStructDescriptor = &autowire.StructDescriptor{
		Alias: "github.com/alibaba/ioc-golang/example/aop/transaction/distributed/server/pkg/service/api.BankServiceIOCRPCClient",
		Factory: func() interface{} {
			return &BankService{}
		},
		ConstructFunc: func(i interface{}, _ interface{}) (interface{}, error) {
			impl := i.(*BankService)
			var constructFunc BankServiceConstructFunc = InitBankService
			return constructFunc(impl)
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{
				"transaction": map[string]string{
					"AddMoney":    "AddMoneyRollback",
					"RemoveMoney": "RemoveMoneyRollback",
				},
			},
			"autowire": map[string]interface{}{},
		},
	}
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &bankService_{}
		},
	})
	rpc_service.RegisterStructDescriptor(bankServiceStructDescriptor)
}

type BankServiceConstructFunc func(impl *BankService) (*BankService, error)
type bankService_ struct {
	GetMoney_            func(id int) int
	AddMoney_            func(id, num int) error
	AddMoneyRollback_    func(id, num int, errMsg string)
	RemoveMoney_         func(id, num int) error
	RemoveMoneyRollback_ func(id, num int, errMsg string)
}

func (b *bankService_) GetMoney(id int) int {
	return b.GetMoney_(id)
}

func (b *bankService_) AddMoney(id, num int) error {
	return b.AddMoney_(id, num)
}

func (b *bankService_) AddMoneyRollback(id, num int, errMsg string) {
	b.AddMoneyRollback_(id, num, errMsg)
}

func (b *bankService_) RemoveMoney(id, num int) error {
	return b.RemoveMoney_(id, num)
}

func (b *bankService_) RemoveMoneyRollback(id, num int, errMsg string) {
	b.RemoveMoneyRollback_(id, num, errMsg)
}

type BankServiceIOCInterface interface {
	GetMoney(id int) int
	AddMoney(id, num int) error
	AddMoneyRollback(id, num int, errMsg string)
	RemoveMoney(id, num int) error
	RemoveMoneyRollback(id, num int, errMsg string)
}

var _bankServiceSDID string

func GetBankServiceRpc() (*BankService, error) {
	if _bankServiceSDID == "" {
		_bankServiceSDID = util.GetSDIDByStructPtr(new(BankService))
	}
	i, err := rpc_service.GetImpl(_bankServiceSDID)
	if err != nil {
		return nil, err
	}
	impl := i.(*BankService)
	return impl, nil
}

func GetBankServiceIOCInterfaceRpc() (BankServiceIOCInterface, error) {
	if _bankServiceSDID == "" {
		_bankServiceSDID = util.GetSDIDByStructPtr(new(BankService))
	}
	i, err := rpc_service.GetImplWithProxy(_bankServiceSDID)
	if err != nil {
		return nil, err
	}
	impl := i.(BankServiceIOCInterface)
	return impl, nil
}
