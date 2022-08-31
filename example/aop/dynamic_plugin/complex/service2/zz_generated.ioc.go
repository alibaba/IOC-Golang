//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package service2

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
	"github.com/alibaba/ioc-golang/example/aop/dynamic_plugin/complex/service1"
)

func init() {
	var service2StructDescriptor = &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &Service2{}
		},
		ConstructFunc: func(i interface{}, _ interface{}) (interface{}, error) {
			impl := i.(*Service2)
			var constructFunc Service2ConstructFunc = constructService2
			return constructFunc(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &service2_{}
		},
	})
	singleton.RegisterStructDescriptor(service2StructDescriptor)
}

type Service2ConstructFunc func(impl *Service2) (*Service2, error)
type service2_ struct {
	LoadData_             func() string
	SetData_              func(val string)
	GetName_              func() string
	SetName_              func(name string)
	GetService1Normal_    func() service1.Service1IOCInterface
	GetService1Singleton_ func() service1.Service1IOCInterface
}

func (s *service2_) LoadData() string {
	return s.LoadData_()
}

func (s *service2_) SetData(val string) {
	s.SetData_(val)
}

func (s *service2_) GetName() string {
	return s.GetName_()
}

func (s *service2_) SetName(name string) {
	s.SetName_(name)
}

func (s *service2_) GetService1Normal() service1.Service1IOCInterface {
	return s.GetService1Normal_()
}

func (s *service2_) GetService1Singleton() service1.Service1IOCInterface {
	return s.GetService1Singleton_()
}

type Service2IOCInterface interface {
	LoadData() string
	SetData(val string)
	GetName() string
	SetName(name string)
	GetService1Normal() service1.Service1IOCInterface
	GetService1Singleton() service1.Service1IOCInterface
}

var _service2SDID string

func GetService2Singleton() (*Service2, error) {
	if _service2SDID == "" {
		_service2SDID = util.GetSDIDByStructPtr(new(Service2))
	}
	i, err := singleton.GetImpl(_service2SDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*Service2)
	return impl, nil
}

func GetService2IOCInterfaceSingleton() (Service2IOCInterface, error) {
	if _service2SDID == "" {
		_service2SDID = util.GetSDIDByStructPtr(new(Service2))
	}
	i, err := singleton.GetImplWithProxy(_service2SDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(Service2IOCInterface)
	return impl, nil
}

type ThisService2 struct {
}

func (t *ThisService2) This() Service2IOCInterface {
	thisPtr, _ := GetService2IOCInterfaceSingleton()
	return thisPtr
}
