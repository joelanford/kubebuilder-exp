package v1

import (
	"fmt"
	"strings"

	"sigs.k8s.io/kubebuilder-exp/cmd/kubebuilder/cli"

	"sigs.k8s.io/kubebuilder-exp/pkg/plugin"

	"github.com/spf13/pflag"
)

var _ plugin.ProjectScaffolder = &Scaffolder{}
var _ plugin.APIScaffolder = &Scaffolder{}

type Scaffolder struct {
	// project flags
	helmChart     string
	helmChartRepo string

	apiVersion string
	kind       string
}

func (s Scaffolder) ProjectHelp() string {
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

func (s Scaffolder) ProjectExample() string {
	exampleTmpl := `  # Scaffold a project using a sample chart
  COMMAND_NAME init --project-version=helm:1 --api-version=example.com/v1alpha1 --kind=MyApp

  # Scaffold a project using the stable/tomcat chart
  COMMAND_NAME init --project-version=helm:1 --helm-chart=stable/tomcat
`
	return strings.ReplaceAll(exampleTmpl, "COMMAND_NAME", cli.CommandName)
}

func (s Scaffolder) ScaffoldProject() error {
	fmt.Printf("Scaffolding project for project version %q\n", s.Version())
	return nil
}

func (s Scaffolder) Version() string {
	return "helm:1"
}

func (s *Scaffolder) BindProjectFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.helmChart, "helm-chart", "", "Helm chart")
	fs.StringVar(&s.helmChartRepo, "helm-chart-repo", "", "Helm chart repo")
	fs.StringVar(&s.apiVersion, "apiVersion", "", "Kubernetes API version (e.g. example.com/v1alpha1)")
	fs.StringVar(&s.apiVersion, "kind", "", "Kubernetes Kind (e.g. MyApp)")
}

func (s Scaffolder) APIHelp() string {
	return `Scaffold a Kubernetes API based on a Helm chart.

create api will generate, copy, or fetch a new Helm chart into the project
and map it to the specified API version and kind. A default role will be
generated based on the resources defined in the Helm chart's default
manifest.
`
}

func (s Scaffolder) APIExample() string {
	exampleTmpl := `  # Scaffold a project using a sample chart
  COMMAND_NAME create api --api-version=example.com/v1alpha1 --kind=MyApp

  # Scaffold a project using the stable/tomcat chart
  COMMAND_NAME create api --helm-chart=stable/tomcat
`
	return strings.ReplaceAll(exampleTmpl, "COMMAND_NAME", cli.CommandName)
}

func (s Scaffolder) ScaffoldAPI() error {
	fmt.Printf("Scaffolding API for project version %q\n", s.Version())
	return nil
}

func (s *Scaffolder) BindAPIFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.helmChart, "helm-chart", "", "Helm chart")
	fs.StringVar(&s.helmChartRepo, "helm-chart-repo", "", "Helm chart repo")
	fs.StringVar(&s.apiVersion, "apiVersion", "", "Kubernetes API version (e.g. example.com/v1alpha1)")
	fs.StringVar(&s.apiVersion, "kind", "", "Kubernetes Kind (e.g. MyApp)")
}
