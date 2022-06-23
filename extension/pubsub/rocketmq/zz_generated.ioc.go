//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package rocketmq

import (
	contextx "context"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"

	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
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
			var _ paramInterface = &Param{}
			return &Param{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(paramInterface)
			impl := i.(*Impl)
			return param.New(impl)
		},
	})
	singleton.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &Impl{}
		},
		ParamFactory: func() interface{} {
			var _ paramInterface = &Param{}
			return &Param{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(paramInterface)
			impl := i.(*Impl)
			return param.New(impl)
		},
	})
}

type paramInterface interface {
	New(impl *Impl) (*Impl, error)
}
type impl_ struct {
	Subscribe_   func(topic string, selector consumer.MessageSelector, f func(contextx.Context, ...*primitive.MessageExt) (consumer.ConsumeResult, error)) error
	Unsubscribe_ func(topic string) error
	SendSync_    func(ctx contextx.Context, mq ...*primitive.Message) (*primitive.SendResult, error)
	SendAsync_   func(ctx contextx.Context, mq func(ctx contextx.Context, result *primitive.SendResult, err error), msg ...*primitive.Message) error
	SendOneWay_  func(ctx contextx.Context, mq ...*primitive.Message) error
}

func (i *impl_) Subscribe(topic string, selector consumer.MessageSelector, f func(contextx.Context, ...*primitive.MessageExt) (consumer.ConsumeResult, error)) error {
	return i.Subscribe_(topic, selector, f)
}

func (i *impl_) Unsubscribe(topic string) error {
	return i.Unsubscribe_(topic)
}

func (i *impl_) SendSync(ctx contextx.Context, mq ...*primitive.Message) (*primitive.SendResult, error) {
	return i.SendSync_(ctx, mq...)
}

func (i *impl_) SendAsync(ctx contextx.Context, mq func(ctx contextx.Context, result *primitive.SendResult, err error), msg ...*primitive.Message) error {
	return i.SendAsync_(ctx, mq, msg...)
}

func (i *impl_) SendOneWay(ctx contextx.Context, mq ...*primitive.Message) error {
	return i.SendOneWay_(ctx, mq...)
}

type ImplIOCInterface interface {
	Subscribe(topic string, selector consumer.MessageSelector, f func(contextx.Context, ...*primitive.MessageExt) (consumer.ConsumeResult, error)) error
	Unsubscribe(topic string) error
	SendSync(ctx contextx.Context, mq ...*primitive.Message) (*primitive.SendResult, error)
	SendAsync(ctx contextx.Context, mq func(ctx contextx.Context, result *primitive.SendResult, err error), msg ...*primitive.Message) error
	SendOneWay(ctx contextx.Context, mq ...*primitive.Message) error
}

func GetImpl(p *Param) (*Impl, error) {
	i, err := normal.GetImpl(util.GetSDIDByStructPtr(new(Impl)), p)
	if err != nil {
		return nil, err
	}
	impl := i.(*Impl)
	return impl, nil
}

func GetImplIOCInterface(p *Param) (ImplIOCInterface, error) {
	i, err := normal.GetImplWithProxy(util.GetSDIDByStructPtr(new(Impl)), p)
	if err != nil {
		return nil, err
	}
	impl := i.(ImplIOCInterface)
	return impl, nil
}

func GetImplSingleton(p *Param) (*Impl, error) {
	i, err := singleton.GetImpl(util.GetSDIDByStructPtr(new(Impl)), p)
	if err != nil {
		return nil, err
	}
	impl := i.(*Impl)
	return impl, nil
}

func GetImplIOCInterfaceSingleton(p *Param) (ImplIOCInterface, error) {
	i, err := singleton.GetImplWithProxy(util.GetSDIDByStructPtr(new(Impl)), p)
	if err != nil {
		return nil, err
	}
	impl := i.(ImplIOCInterface)
	return impl, nil
}
