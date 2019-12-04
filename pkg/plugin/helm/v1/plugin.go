/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

type Plugin struct {
	// cli metadata
	commandName string

	// project flags
	helmChart     string
	helmChartRepo string

	apiVersion string
	kind       string
}

func (_ Plugin) Version() string {
	return "helm:1"
}

func (p *Plugin) InjectCommandName(name string) {
	p.commandName = name
}
