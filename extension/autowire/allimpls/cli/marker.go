/*
 * Copyright (c) 2022, Alibaba Group;
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cli

import "sigs.k8s.io/controller-tools/pkg/markers"

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:proxy=false
// +ioc:autowire:allimpls:interface=github.com/alibaba/ioc-golang/iocli/gen/marker.DefinitionGetter

type iocGolangAutowireAllImplsInterfaceMarker struct {
}

func (m *iocGolangAutowireAllImplsInterfaceMarker) GetMarkerDefinition() *markers.Definition {
	return markers.Must(markers.MakeDefinition(allimplsInterfaceAnnotation, markers.DescribesType, ""))
}

// +ioc:autowire=true
// +ioc:autowire:type=allimpls
// +ioc:autowire:proxy=false
// +ioc:autowire:allimpls:interface=github.com/alibaba/ioc-golang/iocli/gen/marker.DefinitionGetter

type iocGolangAutowireAllImplsTypeMarker struct {
}

func (m *iocGolangAutowireAllImplsTypeMarker) GetMarkerDefinition() *markers.Definition {
	return markers.Must(markers.MakeDefinition(allimplsAutowireTypeAnnotation, markers.DescribesType, ""))
}
