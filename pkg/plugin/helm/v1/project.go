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

var _ plugin.InitPlugin = &Plugin{}

func (_ Plugin) InitDescription() string {
	return `Initialize a new project based on a Helm chart.

Writes the following files:
- a watches.yaml file to configure mappings of GVKs to Helm charts 
- a PROJECT file with the domain and repo
- a Makefile to build the project
- a Kustomization.yaml for customizating manifests
- a Patch file for customizing image for manager manifests
- a Patch file for enabling prometheus metrics
- a copy of the initial helm chart to be used in the operator
`
}

func (p Plugin) InitExample() string {
	return fmt.Sprintf(`  # Scaffold a project using a sample chart
  %s init --project-version=helm:1 --api-version=example.com/v1alpha1 --kind=MyApp

  # Scaffold a project using the stable/tomcat chart
  %s init --project-version=helm:1 --helm-chart=stable/tomcat
`, p.commandName, p.commandName)
}

func (p *Plugin) BindInitFlags(fs *pflag.FlagSet) {
	fs.StringVar(&p.helmChart, "helm-chart", "", "Helm chart")
	fs.StringVar(&p.helmChartRepo, "helm-chart-repo", "", "Helm chart repo")
	fs.StringVar(&p.apiVersion, "apiVersion", "", "Kubernetes API version (e.g. example.com/v1alpha1)")
	fs.StringVar(&p.kind, "kind", "", "Kubernetes Kind (e.g. MyApp)")
}

func (p Plugin) Init() error {
	fmt.Printf("Scaffolding project for project version %q\n", p.Version())
	return nil
}
