// Code generated by iocli

package api

import (
	"github.com/alibaba/ioc-golang/autowire"
	"github.com/alibaba/ioc-golang/autowire/normal"
	"github.com/alibaba/ioc-golang/example/autowire_rpc/client/test/dto"
	"github.com/alibaba/ioc-golang/extension/autowire/rpc/rpc_client"
)

func init() {
	rpc_client.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &complexServiceIOCRPCClient{}
		},
	})
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &complexServiceIOCRPCClient_{}
		},
	})
}

type complexServiceIOCRPCClient_ struct {
	RPCBasicType_                  func(name string, age int, age32 int32, age64 int64, ageF32 float32, ageF64 float64, namePtr *string, agePtr *int, age32Ptr *int32, age64Ptr *int64, ageF32Ptr *float32, ageF64Ptr *float64) (string, int, int32, int64, float32, float64, *string, *int, *int32, *int64, *float32, *float64)
	RPCWithoutParamAndReturnValue_ func()
	RPCWithoutParam_               func() (*dto.User, error)
	RPCWithoutReturnValue_         func(user *dto.User)
	RPCWithCustomValue_            func(customStruct dto.CustomStruct, customStruct2 *dto.CustomStruct) (dto.CustomStruct, *dto.CustomStruct)
	RPCWithError_                  func() (*dto.User, error)
	RPCWithParamCustomMethod_      func(customStruct dto.CustomStruct) dto.User
}

func (c *complexServiceIOCRPCClient_) RPCBasicType(name string, age int, age32 int32, age64 int64, ageF32 float32, ageF64 float64, namePtr *string, agePtr *int, age32Ptr *int32, age64Ptr *int64, ageF32Ptr *float32, ageF64Ptr *float64) (string, int, int32, int64, float32, float64, *string, *int, *int32, *int64, *float32, *float64) {
	return c.RPCBasicType_(name, age, age32, age64, ageF32, ageF64, namePtr, agePtr, age32Ptr, age64Ptr, ageF32Ptr, ageF64Ptr)
}
func (c *complexServiceIOCRPCClient_) RPCWithoutParamAndReturnValue() {
	c.RPCWithoutParamAndReturnValue_()
}
func (c *complexServiceIOCRPCClient_) RPCWithoutParam() (*dto.User, error) {
	return c.RPCWithoutParam_()
}
func (c *complexServiceIOCRPCClient_) RPCWithoutReturnValue(user *dto.User) {
	c.RPCWithoutReturnValue_(user)
}
func (c *complexServiceIOCRPCClient_) RPCWithCustomValue(customStruct dto.CustomStruct, customStruct2 *dto.CustomStruct) (dto.CustomStruct, *dto.CustomStruct) {
	return c.RPCWithCustomValue_(customStruct, customStruct2)
}
func (c *complexServiceIOCRPCClient_) RPCWithError() (*dto.User, error) {
	return c.RPCWithError_()
}
func (c *complexServiceIOCRPCClient_) RPCWithParamCustomMethod(customStruct dto.CustomStruct) dto.User {
	return c.RPCWithParamCustomMethod_(customStruct)
}

type ComplexServiceIOCRPCClient interface {
	RPCBasicType(name string, age int, age32 int32, age64 int64, ageF32 float32, ageF64 float64, namePtr *string, agePtr *int, age32Ptr *int32, age64Ptr *int64, ageF32Ptr *float32, ageF64Ptr *float64) (string, int, int32, int64, float32, float64, *string, *int, *int32, *int64, *float32, *float64)
	RPCWithoutParamAndReturnValue()
	RPCWithoutParam() (*dto.User, error)
	RPCWithoutReturnValue(user *dto.User)
	RPCWithCustomValue(customStruct dto.CustomStruct, customStruct2 *dto.CustomStruct) (dto.CustomStruct, *dto.CustomStruct)
	RPCWithError() (*dto.User, error)
	RPCWithParamCustomMethod(customStruct dto.CustomStruct) dto.User
}

type complexServiceIOCRPCClient struct {
	RPCBasicType                  func(name string, age int, age32 int32, age64 int64, ageF32 float32, ageF64 float64, namePtr *string, agePtr *int, age32Ptr *int32, age64Ptr *int64, ageF32Ptr *float32, ageF64Ptr *float64) (string, int, int32, int64, float32, float64, *string, *int, *int32, *int64, *float32, *float64)
	RPCWithoutParamAndReturnValue func()
	RPCWithoutParam               func() (*dto.User, error)
	RPCWithoutReturnValue         func(user *dto.User)
	RPCWithCustomValue            func(customStruct dto.CustomStruct, customStruct2 *dto.CustomStruct) (dto.CustomStruct, *dto.CustomStruct)
	RPCWithError                  func() (*dto.User, error)
	RPCWithParamCustomMethod      func(customStruct dto.CustomStruct) dto.User
}
