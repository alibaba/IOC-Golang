//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package call

import (
	"github.com/sirupsen/logrus"

	"github.com/alibaba/ioc-golang/aop"
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
	debugLogContextStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &debugLogContext{}
		},
		ParamFactory: func() interface{} {
			var _ debugLogContextParamInterface = &debugLogContextParam{}
			return &debugLogContextParam{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(debugLogContextParamInterface)
			impl := i.(*debugLogContext)
			return param.init(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	normal.RegisterStructDescriptor(debugLogContextStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &globalLogrusIOCCtxHook_{}
		},
	})
	globalLogrusIOCCtxHookStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &GlobalLogrusIOCCtxHook{}
		},
		ParamFactory: func() interface{} {
			var _ globalLogrusIOCCtxHookParamInterface = &globalLogrusIOCCtxHookParam{}
			return &globalLogrusIOCCtxHookParam{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(globalLogrusIOCCtxHookParamInterface)
			impl := i.(*GlobalLogrusIOCCtxHook)
			return param.newLogrusIOCCtxHook(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	singleton.RegisterStructDescriptor(globalLogrusIOCCtxHookStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &logInterceptor_{}
		},
	})
	logInterceptorStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &logInterceptor{}
		},
		ParamFactory: func() interface{} {
			var _ logInterceptorParamsInterface = &logInterceptorParams{}
			return &logInterceptorParams{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(logInterceptorParamsInterface)
			impl := i.(*logInterceptor)
			return param.initLogInterceptor(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	singleton.RegisterStructDescriptor(logInterceptorStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &logGoRoutineInterceptorFacadeCtx_{}
		},
	})
	logGoRoutineInterceptorFacadeCtxStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &logGoRoutineInterceptorFacadeCtx{}
		},
		ParamFactory: func() interface{} {
			var _ logGoRoutineInterceptorFacadeCtxParamInterface = &logGoRoutineInterceptorFacadeCtxParam{}
			return &logGoRoutineInterceptorFacadeCtxParam{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(logGoRoutineInterceptorFacadeCtxParamInterface)
			impl := i.(*logGoRoutineInterceptorFacadeCtx)
			return param.initLogGoRoutineInterceptorFacadeCtx(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	normal.RegisterStructDescriptor(logGoRoutineInterceptorFacadeCtxStructDescriptor)
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &invocationCtxNotifyHook_{}
		},
	})
	invocationCtxNotifyHookStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &invocationCtxNotifyHook{}
		},
		ParamFactory: func() interface{} {
			var _ invocationCtxNotifyHookParamInterface = &invocationCtxNotifyHookParam{}
			return &invocationCtxNotifyHookParam{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(invocationCtxNotifyHookParamInterface)
			impl := i.(*invocationCtxNotifyHook)
			return param.initInvocationCtxNotifyHook(impl)
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
	}
	singleton.RegisterStructDescriptor(invocationCtxNotifyHookStructDescriptor)
	invocationCtxLogsGeneratorStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &InvocationCtxLogsGenerator{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	singleton.RegisterStructDescriptor(invocationCtxLogsGeneratorStructDescriptor)
	logServiceImplStructDescriptor := &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &logServiceImpl{}
		},
		Metadata: map[string]interface{}{
			"aop":      map[string]interface{}{},
			"autowire": map[string]interface{}{},
		},
		DisableProxy: true,
	}
	singleton.RegisterStructDescriptor(logServiceImplStructDescriptor)
}

type logGoRoutineInterceptorFacadeCtxParamInterface interface {
	initLogGoRoutineInterceptorFacadeCtx(impl *logGoRoutineInterceptorFacadeCtx) (*logGoRoutineInterceptorFacadeCtx, error)
}
type invocationCtxNotifyHookParamInterface interface {
	initInvocationCtxNotifyHook(impl *invocationCtxNotifyHook) (*invocationCtxNotifyHook, error)
}
type debugLogContextParamInterface interface {
	init(impl *debugLogContext) (*debugLogContext, error)
}
type globalLogrusIOCCtxHookParamInterface interface {
	newLogrusIOCCtxHook(impl *GlobalLogrusIOCCtxHook) (*GlobalLogrusIOCCtxHook, error)
}
type logInterceptorParamsInterface interface {
	initLogInterceptor(impl *logInterceptor) (*logInterceptor, error)
}
type globalLogrusIOCCtxHook_ struct {
	Levels_      func() []logrus.Level
	Fire_        func(entry *logrus.Entry) error
	SetLogLevel_ func(level uint32)
}

func (g *globalLogrusIOCCtxHook_) Levels() []logrus.Level {
	return g.Levels_()
}

func (g *globalLogrusIOCCtxHook_) Fire(entry *logrus.Entry) error {
	return g.Fire_(entry)
}

func (g *globalLogrusIOCCtxHook_) SetLogLevel(level uint32) {
	g.SetLogLevel_(level)
}

type logInterceptor_ struct {
	BeforeInvoke_           func(ctx *aop.InvocationContext)
	AfterInvoke_            func(ctx *aop.InvocationContext)
	WatchLogs_              func(logCtx *debugLogContext)
	StopWatch_              func()
	NotifyEntry_            func(entry *logrus.Entry, hookType string)
	GetInvocationCtxLogger_ func() *logrus.Logger
}

func (l *logInterceptor_) BeforeInvoke(ctx *aop.InvocationContext) {
	l.BeforeInvoke_(ctx)
}

func (l *logInterceptor_) AfterInvoke(ctx *aop.InvocationContext) {
	l.AfterInvoke_(ctx)
}

func (l *logInterceptor_) WatchLogs(logCtx *debugLogContext) {
	l.WatchLogs_(logCtx)
}

func (l *logInterceptor_) StopWatch() {
	l.StopWatch_()
}

func (l *logInterceptor_) NotifyEntry(entry *logrus.Entry, hookType string) {
	l.NotifyEntry_(entry, hookType)
}

func (l *logInterceptor_) GetInvocationCtxLogger() *logrus.Logger {
	return l.GetInvocationCtxLogger_()
}

type logGoRoutineInterceptorFacadeCtx_ struct {
	pushEntry_    func(entry *logrus.Entry, hookType string)
	BeforeInvoke_ func(ctx *aop.InvocationContext)
	AfterInvoke_  func(ctx *aop.InvocationContext)
	Type_         func() string
}

func (l *logGoRoutineInterceptorFacadeCtx_) pushEntry(entry *logrus.Entry, hookType string) {
	l.pushEntry_(entry, hookType)
}

func (l *logGoRoutineInterceptorFacadeCtx_) BeforeInvoke(ctx *aop.InvocationContext) {
	l.BeforeInvoke_(ctx)
}

func (l *logGoRoutineInterceptorFacadeCtx_) AfterInvoke(ctx *aop.InvocationContext) {
	l.AfterInvoke_(ctx)
}

func (l *logGoRoutineInterceptorFacadeCtx_) Type() string {
	return l.Type_()
}

type invocationCtxNotifyHook_ struct {
	Levels_ func() []logrus.Level
	Fire_   func(entry *logrus.Entry) error
}

func (i *invocationCtxNotifyHook_) Levels() []logrus.Level {
	return i.Levels_()
}

func (i *invocationCtxNotifyHook_) Fire(entry *logrus.Entry) error {
	return i.Fire_(entry)
}

type GlobalLogrusIOCCtxHookIOCInterface interface {
	Levels() []logrus.Level
	Fire(entry *logrus.Entry) error
	SetLogLevel(level uint32)
}

type logInterceptorIOCInterface interface {
	BeforeInvoke(ctx *aop.InvocationContext)
	AfterInvoke(ctx *aop.InvocationContext)
	WatchLogs(logCtx *debugLogContext)
	StopWatch()
	NotifyEntry(entry *logrus.Entry, hookType string)
	GetInvocationCtxLogger() *logrus.Logger
}

type logGoRoutineInterceptorFacadeCtxIOCInterface interface {
	pushEntry(entry *logrus.Entry, hookType string)
	BeforeInvoke(ctx *aop.InvocationContext)
	AfterInvoke(ctx *aop.InvocationContext)
	Type() string
}

type invocationCtxNotifyHookIOCInterface interface {
	Levels() []logrus.Level
	Fire(entry *logrus.Entry) error
}

var _debugLogContextSDID string

func GetdebugLogContext(p *debugLogContextParam) (*debugLogContext, error) {
	if _debugLogContextSDID == "" {
		_debugLogContextSDID = util.GetSDIDByStructPtr(new(debugLogContext))
	}
	i, err := normal.GetImpl(_debugLogContextSDID, p)
	if err != nil {
		return nil, err
	}
	impl := i.(*debugLogContext)
	return impl, nil
}

var _globalLogrusIOCCtxHookSDID string

func GetGlobalLogrusIOCCtxHookSingleton(p *globalLogrusIOCCtxHookParam) (*GlobalLogrusIOCCtxHook, error) {
	if _globalLogrusIOCCtxHookSDID == "" {
		_globalLogrusIOCCtxHookSDID = util.GetSDIDByStructPtr(new(GlobalLogrusIOCCtxHook))
	}
	i, err := singleton.GetImpl(_globalLogrusIOCCtxHookSDID, p)
	if err != nil {
		return nil, err
	}
	impl := i.(*GlobalLogrusIOCCtxHook)
	return impl, nil
}

func GetGlobalLogrusIOCCtxHookIOCInterfaceSingleton(p *globalLogrusIOCCtxHookParam) (GlobalLogrusIOCCtxHookIOCInterface, error) {
	if _globalLogrusIOCCtxHookSDID == "" {
		_globalLogrusIOCCtxHookSDID = util.GetSDIDByStructPtr(new(GlobalLogrusIOCCtxHook))
	}
	i, err := singleton.GetImplWithProxy(_globalLogrusIOCCtxHookSDID, p)
	if err != nil {
		return nil, err
	}
	impl := i.(GlobalLogrusIOCCtxHookIOCInterface)
	return impl, nil
}

type ThisGlobalLogrusIOCCtxHook struct {
}

func (t *ThisGlobalLogrusIOCCtxHook) This() GlobalLogrusIOCCtxHookIOCInterface {
	thisPtr, _ := GetGlobalLogrusIOCCtxHookIOCInterfaceSingleton(nil)
	return thisPtr
}

var _logInterceptorSDID string

func GetlogInterceptorSingleton(p *logInterceptorParams) (*logInterceptor, error) {
	if _logInterceptorSDID == "" {
		_logInterceptorSDID = util.GetSDIDByStructPtr(new(logInterceptor))
	}
	i, err := singleton.GetImpl(_logInterceptorSDID, p)
	if err != nil {
		return nil, err
	}
	impl := i.(*logInterceptor)
	return impl, nil
}

func GetlogInterceptorIOCInterfaceSingleton(p *logInterceptorParams) (logInterceptorIOCInterface, error) {
	if _logInterceptorSDID == "" {
		_logInterceptorSDID = util.GetSDIDByStructPtr(new(logInterceptor))
	}
	i, err := singleton.GetImplWithProxy(_logInterceptorSDID, p)
	if err != nil {
		return nil, err
	}
	impl := i.(logInterceptorIOCInterface)
	return impl, nil
}

type ThislogInterceptor struct {
}

func (t *ThislogInterceptor) This() logInterceptorIOCInterface {
	thisPtr, _ := GetlogInterceptorIOCInterfaceSingleton(nil)
	return thisPtr
}

var _logGoRoutineInterceptorFacadeCtxSDID string

func GetlogGoRoutineInterceptorFacadeCtx(p *logGoRoutineInterceptorFacadeCtxParam) (*logGoRoutineInterceptorFacadeCtx, error) {
	if _logGoRoutineInterceptorFacadeCtxSDID == "" {
		_logGoRoutineInterceptorFacadeCtxSDID = util.GetSDIDByStructPtr(new(logGoRoutineInterceptorFacadeCtx))
	}
	i, err := normal.GetImpl(_logGoRoutineInterceptorFacadeCtxSDID, p)
	if err != nil {
		return nil, err
	}
	impl := i.(*logGoRoutineInterceptorFacadeCtx)
	return impl, nil
}

func GetlogGoRoutineInterceptorFacadeCtxIOCInterface(p *logGoRoutineInterceptorFacadeCtxParam) (logGoRoutineInterceptorFacadeCtxIOCInterface, error) {
	if _logGoRoutineInterceptorFacadeCtxSDID == "" {
		_logGoRoutineInterceptorFacadeCtxSDID = util.GetSDIDByStructPtr(new(logGoRoutineInterceptorFacadeCtx))
	}
	i, err := normal.GetImplWithProxy(_logGoRoutineInterceptorFacadeCtxSDID, p)
	if err != nil {
		return nil, err
	}
	impl := i.(logGoRoutineInterceptorFacadeCtxIOCInterface)
	return impl, nil
}

var _invocationCtxNotifyHookSDID string

func GetinvocationCtxNotifyHookSingleton(p *invocationCtxNotifyHookParam) (*invocationCtxNotifyHook, error) {
	if _invocationCtxNotifyHookSDID == "" {
		_invocationCtxNotifyHookSDID = util.GetSDIDByStructPtr(new(invocationCtxNotifyHook))
	}
	i, err := singleton.GetImpl(_invocationCtxNotifyHookSDID, p)
	if err != nil {
		return nil, err
	}
	impl := i.(*invocationCtxNotifyHook)
	return impl, nil
}

func GetinvocationCtxNotifyHookIOCInterfaceSingleton(p *invocationCtxNotifyHookParam) (invocationCtxNotifyHookIOCInterface, error) {
	if _invocationCtxNotifyHookSDID == "" {
		_invocationCtxNotifyHookSDID = util.GetSDIDByStructPtr(new(invocationCtxNotifyHook))
	}
	i, err := singleton.GetImplWithProxy(_invocationCtxNotifyHookSDID, p)
	if err != nil {
		return nil, err
	}
	impl := i.(invocationCtxNotifyHookIOCInterface)
	return impl, nil
}

type ThisinvocationCtxNotifyHook struct {
}

func (t *ThisinvocationCtxNotifyHook) This() invocationCtxNotifyHookIOCInterface {
	thisPtr, _ := GetinvocationCtxNotifyHookIOCInterfaceSingleton(nil)
	return thisPtr
}

var _invocationCtxLogsGeneratorSDID string

func GetInvocationCtxLogsGeneratorSingleton() (*InvocationCtxLogsGenerator, error) {
	if _invocationCtxLogsGeneratorSDID == "" {
		_invocationCtxLogsGeneratorSDID = util.GetSDIDByStructPtr(new(InvocationCtxLogsGenerator))
	}
	i, err := singleton.GetImpl(_invocationCtxLogsGeneratorSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*InvocationCtxLogsGenerator)
	return impl, nil
}

var _logServiceImplSDID string

func GetlogServiceImplSingleton() (*logServiceImpl, error) {
	if _logServiceImplSDID == "" {
		_logServiceImplSDID = util.GetSDIDByStructPtr(new(logServiceImpl))
	}
	i, err := singleton.GetImpl(_logServiceImplSDID, nil)
	if err != nil {
		return nil, err
	}
	impl := i.(*logServiceImpl)
	return impl, nil
}
