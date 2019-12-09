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
	"fmt"

	"sigs.k8s.io/kubebuilder-exp/pkg/plugin"

	"github.com/spf13/pflag"
)

var _ plugin.CreateAPIPlugin = &Plugin{}

func (_ Plugin) CreateAPIDescription() string {
	return `Scaffold a Kubernetes API based on a Helm chart.

create api will generate, copy, or fetch a new Helm chart into the project
and map it to the specified API version and kind. A default role will be
generated based on the resources defined in the Helm chart'p default
manifest.
`
}

func (p Plugin) CreateAPIExample() string {
	return fmt.Sprintf(`  # Scaffold a project using a sample chart
  %s create api --api-version=example.com/v1alpha1 --kind=MyApp

  # Scaffold a project using the stable/tomcat chart
  %s create api --helm-chart=stable/tomcat
`, p.commandName, p.commandName)
}

func (p *Plugin) BindCreateAPIFlags(fs *pflag.FlagSet) {
	fs.StringVar(&p.helmChart, "helm-chart", "", "Helm chart")
	fs.StringVar(&p.helmChartRepo, "helm-chart-repo", "", "Helm chart repo")
	fs.StringVar(&p.apiVersion, "apiVersion", "", "Kubernetes API version (e.g. example.com/v1alpha1)")
	fs.StringVar(&p.kind, "kind", "", "Kubernetes Kind (e.g. MyApp)")
}

func (p Plugin) CreateAPI() error {
	fmt.Printf("Scaffolding API for project version %q\n", p.Version())
	return nil
}
