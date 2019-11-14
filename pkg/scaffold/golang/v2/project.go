package v2

import (
	"fmt"
	"os"
	"strings"

	"sigs.k8s.io/kubebuilder-exp/cmd/kubebuilder/cli"

	"sigs.k8s.io/kubebuilder-exp/pkg/plugin"

	"github.com/spf13/pflag"
	"sigs.k8s.io/kubebuilder-exp/pkg/scaffold/golang/project"
)

var _ plugin.ProjectScaffolder = &Scaffolder{}
var _ plugin.APIScaffolder = &Scaffolder{}

type Scaffolder struct {
	// project flags
	fetchDeps          bool
	skipGoVersionCheck bool

	boilerplate project.Boilerplate
	project     project.Project

	// api flags
	resourceFlag, controllerFlag *pflag.Flag
	runMake                      bool
	pattern                      string

	// final result
	// projectScaffolder scaffold.ProjectScaffolder
	// apiScaffolder api.API
}

func (s Scaffolder) ProjectHelp() string {
	return `Initialize a new project including vendor/ directory and Go package directories.

Writes the following files:
- a boilerplate license file
- a PROJECT file with the domain and repo
- a Makefile to build the project
- a go.mod with project dependencies
- a Kustomization.yaml for customizating manifests
- a Patch file for customizing image for manager manifests
- a Patch file for enabling prometheus metrics
- a cmd/manager/main.go to run

project will prompt the user to run 'dep ensure' after writing the project files.
`
}

func (s Scaffolder) ProjectExample() string {
	exampleTmpl := `  # Scaffold a project using the apache2 license with "The Kubernetes authors" as owners
  COMMAND_NAME init --project-version=2 --domain example.org --license apache2 --owner "The Kubernetes authors"
`
	return strings.ReplaceAll(exampleTmpl, "COMMAND_NAME", cli.CommandName)
}

func (s Scaffolder) ScaffoldProject() error {
	fmt.Printf("Scaffolding project for project version %q\n", s.Version())
	return nil
}

func (s Scaffolder) Version() string {
	return "2"
}

func (s *Scaffolder) BindProjectFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&s.skipGoVersionCheck, "skip-go-version-check", false, "if specified, skip checking the Go version")

	// dependency args
	fs.BoolVar(&s.fetchDeps, "fetch-deps", true, "ensure dependencies are downloaded")

	// boilerplate args
	fs.StringVar(&s.boilerplate.Path, "path", "", "path for boilerplate")
	fs.StringVar(&s.boilerplate.License, "license", "apache2", "license to use to boilerplate.  May be one of apache2,none")
	fs.StringVar(&s.boilerplate.Owner, "owner", "", "Owner to add to the copyright")

	// project args
	fs.StringVar(&s.project.Repo, "repo", "", "name to use for go module, e.g. github.com/user/repo.  "+
		"defaults to the go package of the current working directory.")
	fs.StringVar(&s.project.Domain, "domain", "my.domain", "domain for groups")
}

func (s Scaffolder) APIHelp() string {
	return `Scaffold a Kubernetes API by creating a Resource definition and / or a Controller.

create resource will prompt the user for if it should scaffold the Resource and / or Controller.  To only
scaffold a Controller for an existing Resource, select "n" for Resource.  To only define
the schema for a Resource without writing a Controller, select "n" for Controller.

After the scaffold is written, api will run make on the project.
`
}

func (s Scaffolder) APIExample() string {
	exampleTmpl := `  # Create a frigates API with Group: ship, Version: v1beta1 and Kind: Frigate
  COMMAND_NAME create api --group ship --version v1beta1 --kind Frigate
  
  # Edit the API Scheme
  nano api/v1beta1/frigate_types.go

  # Edit the Controller
  nano controllers/frigate/frigate_controller.go

  # Edit the Controller Test
  nano controllers/frigate/frigate_controller_test.go

  # Install CRDs into the Kubernetes cluster using kubectl apply
  make install

  # Regenerate code and run against the Kubernetes cluster configured by ~/.kube/config
  make run
`
	return strings.ReplaceAll(exampleTmpl, "COMMAND_NAME", cli.CommandName)

}

func (s Scaffolder) ScaffoldAPI() error {
	fmt.Printf("Scaffolding API for project version %q\n", s.Version())
	return nil
}

func (s *Scaffolder) BindAPIFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&s.runMake, "make", true,
		"if true, run make after generating files")
	//fs.BoolVar(&s.apiScaffolder.DoResource, "resource", true,
	//	"if set, generate the resource without prompting the user")
	s.resourceFlag = fs.Lookup("resource")
	//fs.BoolVar(&s.apiScaffolder.DoController, "controller", true,
	//	"if set, generate the controller without prompting the user")
	s.controllerFlag = fs.Lookup("controller")
	if os.Getenv("KUBEBUILDER_ENABLE_PLUGINS") != "" {
		fs.StringVar(&s.pattern, "pattern", "",
			"generates an API following an extension pattern (addon)")
	}
	//s.apiScaffolder.Resource = resourceForFlags(fs)
}
