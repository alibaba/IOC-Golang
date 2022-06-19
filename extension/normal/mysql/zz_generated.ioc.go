//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli

package mysql

import (
	"github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	util "github.com/alibaba/ioc-golang/autowire/util"
	"gorm.io/gorm"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &impl_{}
		},
	})
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &Impl{}
		},
		ParamFactory: func() interface{} {
			var _ configInterface = &Config{}
			return &Config{}
		},
		ParamLoader: &paramLoader{},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(configInterface)
			impl := i.(*Impl)
			return param.New(impl)
		},
	})
}

type configInterface interface {
	New(impl *Impl) (*Impl, error)
}
type impl_ struct {
	GetDB_       func() *gorm.DB
	SelectWhere_ func(queryStr string, result interface{}, args ...interface{}) error
	Insert_      func(toInsertLines UserDefinedModel) error
	Update_      func(queryStr, field string, target interface{}, args ...interface{}) error
	Delete_      func(toDeleteTarget UserDefinedModel) error
	First_       func(queryStr string, findTarget UserDefinedModel, args ...interface{}) error
}

func (i *impl_) GetDB() *gorm.DB {
	return i.GetDB_()
}
func (i *impl_) SelectWhere(queryStr string, result interface{}, args ...interface{}) error {
	return i.SelectWhere_(queryStr, result, args...)
}
func (i *impl_) Insert(toInsertLines UserDefinedModel) error {
	return i.Insert_(toInsertLines)
}
func (i *impl_) Update(queryStr, field string, target interface{}, args ...interface{}) error {
	return i.Update_(queryStr, field, target, args...)
}
func (i *impl_) Delete(toDeleteTarget UserDefinedModel) error {
	return i.Delete_(toDeleteTarget)
}
func (i *impl_) First(queryStr string, findTarget UserDefinedModel, args ...interface{}) error {
	return i.First_(queryStr, findTarget, args...)
}

type ImplIOCInterface interface {
	GetDB() *gorm.DB
	SelectWhere(queryStr string, result interface{}, args ...interface{}) error
	Insert(toInsertLines UserDefinedModel) error
	Update(queryStr, field string, target interface{}, args ...interface{}) error
	Delete(toDeleteTarget UserDefinedModel) error
	First(queryStr string, findTarget UserDefinedModel, args ...interface{}) error
}

func GetImpl(p *Config) (*Impl, error) {
	i, err := normal.GetImpl(util.GetSDIDByStructPtr(new(Impl)), p)
	if err != nil {
		return nil, err
	}
	impl := i.(*Impl)
	return impl, nil
}
func GetImplIOCInterface(p *Config) (ImplIOCInterface, error) {
	i, err := normal.GetImplWithProxy(util.GetSDIDByStructPtr(new(Impl)), p)
	if err != nil {
		return nil, err
	}
	impl := i.(ImplIOCInterface)
	return impl, nil
}
