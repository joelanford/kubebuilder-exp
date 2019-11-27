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

import (
	"sigs.k8s.io/kubebuilder-exp/pkg/plugin/golang/project"

	"github.com/spf13/pflag"
)

type Plugin struct {
	// cli metadata
	commandName string

	// init flags
	fetchDeps          bool
	skipGoVersionCheck bool

	boilerplate project.Boilerplate
	project     project.Project

	// deprecated init flags
	dep     bool
	depFlag *pflag.Flag
	depArgs []string

	// api flags
	resourceFlag, controllerFlag *pflag.Flag
	runMake                      bool
	pattern                      string

	// final result
	// projectScaffolder scaffold.ProjectScaffolder
	// apiScaffolder api.API
}

func (_ Plugin) Version() string {
	return "1"
}

func (p *Plugin) InjectCommandName(name string) {
	p.commandName = name
}

func (_ Plugin) DeprecationWarning() string {
	return "The v1 projects are deprecated and will not be supported beyond Feb 1, 2020.\nSee how to upgrade your project to v2: https://book.kubebuilder.io/migration/guide.html"
}
