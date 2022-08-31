//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package cli

import (
	"sigs.k8s.io/controller-tools/pkg/markers"

	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	allimpls "github.com/alibaba/ioc-golang/extension/autowire/allimpls"
	"github.com/alibaba/ioc-golang/iocli/gen/generator/plugin"
	marker "github.com/alibaba/ioc-golang/iocli/gen/marker"
)

func init() {
	var transactionFunctionMarkerStructDescriptor = &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &transactionFunctionMarker{}
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"common": map[string]interface{}{
					"implements": []interface{}{
						new(marker.DefinitionGetter),
					},
				},
			},
		},
		DisableProxy: true,
	}
	allimpls.RegisterStructDescriptor(transactionFunctionMarkerStructDescriptor)
	var txCodeGenerationPluginStructDescriptor = &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &txCodeGenerationPlugin{}
		},
		ConstructFunc: func(i interface{}, _ interface{}) (interface{}, error) {
			impl := i.(*txCodeGenerationPlugin)
			var constructFunc txCodeGenerationPluginConstructFunc = create
			return constructFunc(impl)
		},
		Metadata: map[string]interface{}{
			"aop": map[string]interface{}{},
			"autowire": map[string]interface{}{
				"allimpls": map[string]interface{}{
					"autowireType": "normal",
				},
				"common": map[string]interface{}{
					"implements": []interface{}{
						new(plugin.CodeGeneratorPluginForOneStruct),
					},
				},
			},
		},
	}
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &txCodeGenerationPlugin_{}
		},
	})
	allimpls.RegisterStructDescriptor(txCodeGenerationPluginStructDescriptor)
}

type txCodeGenerationPluginConstructFunc func(impl *txCodeGenerationPlugin) (*txCodeGenerationPlugin, error)
type txCodeGenerationPlugin_ struct {
	Name_                           func() string
	Type_                           func() plugin.Type
	Init_                           func(markers markers.MarkerValues)
	GenerateSDMetadataForOneStruct_ func(c plugin.CodeWriter)
	GenerateInFileForOneStruct_     func(c plugin.CodeWriter)
}

func (t *txCodeGenerationPlugin_) Name() string {
	return t.Name_()
}

func (t *txCodeGenerationPlugin_) Type() plugin.Type {
	return t.Type_()
}

func (t *txCodeGenerationPlugin_) Init(markers markers.MarkerValues) {
	t.Init_(markers)
}

func (t *txCodeGenerationPlugin_) GenerateSDMetadataForOneStruct(c plugin.CodeWriter) {
	t.GenerateSDMetadataForOneStruct_(c)
}

func (t *txCodeGenerationPlugin_) GenerateInFileForOneStruct(c plugin.CodeWriter) {
	t.GenerateInFileForOneStruct_(c)
}

type txCodeGenerationPluginIOCInterface interface {
	Name() string
	Type() plugin.Type
	Init(markers markers.MarkerValues)
	GenerateSDMetadataForOneStruct(c plugin.CodeWriter)
	GenerateInFileForOneStruct(c plugin.CodeWriter)
}

var _transactionFunctionMarkerSDID string
var _txCodeGenerationPluginSDID string
