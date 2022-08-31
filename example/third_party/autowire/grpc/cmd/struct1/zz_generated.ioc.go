//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package struct1

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
	var struct1StructDescriptor = &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &Struct1{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &struct1_{}
		},
	})
	singleton.RegisterStructDescriptor(struct1StructDescriptor)
}

type struct1_ struct {
	Hello_ func(name string) string
}

func (s *struct1_) Hello(name string) string {
	return s.Hello_(name)
}

type Struct1IOCInterface interface {
	Hello(name string) string
}

var _struct1SDID string

func GetStruct1Singleton() (*Struct1, error) {
	if _struct1SDID == "" {
		_struct1SDID = util.GetSDIDByStructPtr(new(Struct1))
	}
	i, err := singleton.GetImpl(_struct1SDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*Struct1)
	return impl, nil
}

func GetStruct1IOCInterfaceSingleton() (Struct1IOCInterface, error) {
	if _struct1SDID == "" {
		_struct1SDID = util.GetSDIDByStructPtr(new(Struct1))
	}
	i, err := singleton.GetImplWithProxy(_struct1SDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(Struct1IOCInterface)
	return impl, nil
}

type ThisStruct1 struct {
}

func (t *ThisStruct1) This() Struct1IOCInterface {
	thisPtr, _ := GetStruct1IOCInterfaceSingleton()
	return thisPtr
}
