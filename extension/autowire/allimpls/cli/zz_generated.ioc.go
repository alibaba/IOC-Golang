//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package cli

import (
	"sigs.k8s.io/controller-tools/pkg/markers"

	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	"github.com/alibaba/ioc-golang/extension/autowire/allimpls"
	"github.com/alibaba/ioc-golang/iocli/gen/generator/plugin"
	marker "github.com/alibaba/ioc-golang/iocli/gen/marker"
)

func init() {
	var allImplsCodeGenerationPluginStructDescriptor = &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &allImplsCodeGenerationPlugin{}
		},
		ConstructFunc: func(i interface{}, _ interface{}) (interface{}, error) {
			impl := i.(*allImplsCodeGenerationPlugin)
			var constructFunc allImplsCodeGenerationPluginConstructFunc = create
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
			return &allImplsCodeGenerationPlugin_{}
		},
	})
	allimpls.RegisterStructDescriptor(allImplsCodeGenerationPluginStructDescriptor)
	var iocGolangAutowireAllImplsTypeMarkerStructDescriptor = &autowire.StructDescriptor{
		Factory: func() interface{} {
			return &iocGolangAutowireAllImplsTypeMarker{}
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
	allimpls.RegisterStructDescriptor(iocGolangAutowireAllImplsTypeMarkerStructDescriptor)
}

type allImplsCodeGenerationPluginConstructFunc func(impl *allImplsCodeGenerationPlugin) (*allImplsCodeGenerationPlugin, error)
type allImplsCodeGenerationPlugin_ struct {
	Name_                           func() string
	Type_                           func() plugin.Type
	Init_                           func(markers markers.MarkerValues)
	GenerateSDMetadataForOneStruct_ func(c plugin.CodeWriter)
	GenerateInFileForOneStruct_     func(c plugin.CodeWriter)
}

func (a *allImplsCodeGenerationPlugin_) Name() string {
	return a.Name_()
}

func (a *allImplsCodeGenerationPlugin_) Type() plugin.Type {
	return a.Type_()
}

func (a *allImplsCodeGenerationPlugin_) Init(markers markers.MarkerValues) {
	a.Init_(markers)
}

func (a *allImplsCodeGenerationPlugin_) GenerateSDMetadataForOneStruct(c plugin.CodeWriter) {
	a.GenerateSDMetadataForOneStruct_(c)
}

func (a *allImplsCodeGenerationPlugin_) GenerateInFileForOneStruct(c plugin.CodeWriter) {
	a.GenerateInFileForOneStruct_(c)
}

type allImplsCodeGenerationPluginIOCInterface interface {
	Name() string
	Type() plugin.Type
	Init(markers markers.MarkerValues)
	GenerateSDMetadataForOneStruct(c plugin.CodeWriter)
	GenerateInFileForOneStruct(c plugin.CodeWriter)
}

var _allImplsCodeGenerationPluginSDID string
var _iocGolangAutowireAllImplsTypeMarkerSDID string
