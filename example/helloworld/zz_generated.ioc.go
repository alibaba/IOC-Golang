//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to update this file

package main

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &app_{}
		},
	})
	singleton.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &App{}
		},
	})
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceImpl1_{}
		},
	})
	singleton.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &ServiceImpl1{}
		},
	})
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceImpl2_{}
		},
	})
	singleton.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &ServiceImpl2{}
		},
	})
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceStruct_{}
		},
	})
	singleton.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &ServiceStruct{}
		},
	})
}

type app_ struct {
	Run_ func()
}

func (a *app_) Run() {
	a.Run_()
}

type serviceImpl1_ struct {
	GetHelloString_ func(name string) string
}

func (s *serviceImpl1_) GetHelloString(name string) string {
	return s.GetHelloString_(name)
}

type serviceImpl2_ struct {
	GetHelloString_ func(name string) string
}

func (s *serviceImpl2_) GetHelloString(name string) string {
	return s.GetHelloString_(name)
}

type serviceStruct_ struct {
	GetString_ func(name string) string
}

func (s *serviceStruct_) GetString(name string) string {
	return s.GetString_(name)
}

type AppIOCInterface interface {
	Run()
}
type ServiceImpl1IOCInterface interface {
	GetHelloString(name string) string
}
type ServiceImpl2IOCInterface interface {
	GetHelloString(name string) string
}
type ServiceStructIOCInterface interface {
	GetString(name string) string
}

func GetApp() (*App, error) {
	i, err := singleton.GetImpl(util.GetSDIDByStructPtr(new(App)), nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*App)
	return impl, nil
}
func GetAppIOCInterface() (AppIOCInterface, error) {
	i, err := singleton.GetImplWithProxy(util.GetSDIDByStructPtr(new(App)), nil)
	if err != nil {
		return nil, err
	}
	impl := i.(AppIOCInterface)
	return impl, nil
}
func GetServiceImpl1() (*ServiceImpl1, error) {
	i, err := singleton.GetImpl(util.GetSDIDByStructPtr(new(ServiceImpl1)), nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*ServiceImpl1)
	return impl, nil
}
func GetServiceImpl1IOCInterface() (ServiceImpl1IOCInterface, error) {
	i, err := singleton.GetImplWithProxy(util.GetSDIDByStructPtr(new(ServiceImpl1)), nil)
	if err != nil {
		return nil, err
	}
	impl := i.(ServiceImpl1IOCInterface)
	return impl, nil
}
func GetServiceImpl2() (*ServiceImpl2, error) {
	i, err := singleton.GetImpl(util.GetSDIDByStructPtr(new(ServiceImpl2)), nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*ServiceImpl2)
	return impl, nil
}
func GetServiceImpl2IOCInterface() (ServiceImpl2IOCInterface, error) {
	i, err := singleton.GetImplWithProxy(util.GetSDIDByStructPtr(new(ServiceImpl2)), nil)
	if err != nil {
		return nil, err
	}
	impl := i.(ServiceImpl2IOCInterface)
	return impl, nil
}
func GetServiceStruct() (*ServiceStruct, error) {
	i, err := singleton.GetImpl(util.GetSDIDByStructPtr(new(ServiceStruct)), nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*ServiceStruct)
	return impl, nil
}
func GetServiceStructIOCInterface() (ServiceStructIOCInterface, error) {
	i, err := singleton.GetImplWithProxy(util.GetSDIDByStructPtr(new(ServiceStruct)), nil)
	if err != nil {
		return nil, err
	}
	impl := i.(ServiceStructIOCInterface)
	return impl, nil
}
