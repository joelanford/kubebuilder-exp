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

func (p Plugin) InitExample() string {
	return fmt.Sprintf(`  # Scaffold a project using the apache2 license with "The Kubernetes authors" as owners
  %s init --project-version=1 --domain example.org --license apache2 --owner "The Kubernetes authors"
`, p.commandName)
}

func (p *Plugin) BindInitFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&p.skipGoVersionCheck, "skip-go-version-check", false, "if specified, skip checking the Go version")

	// dependency args
	fs.BoolVar(&p.fetchDeps, "fetch-deps", true, "ensure dependencies are downloaded")

	// deprecated dependency args
	fs.BoolVar(&p.dep, "dep", true, "if specified, determines whether dep will be used.")
	p.depFlag = fs.Lookup("dep")
	fs.StringArrayVar(&p.depArgs, "depArgs", nil, "Additional arguments for dep")
	fs.MarkDeprecated("dep", "use the fetch-deps flag instead")
	fs.MarkDeprecated("depArgs", "will be removed with version 1 scaffolding")

	// boilerplate args
	fs.StringVar(&p.boilerplate.Path, "path", "", "path for boilerplate")
	fs.StringVar(&p.boilerplate.License, "license", "apache2", "license to use to boilerplate.  May be one of apache2,none")
	fs.StringVar(&p.boilerplate.Owner, "owner", "", "Owner to add to the copyright")

	// project args
	fs.StringVar(&p.project.Repo, "repo", "", "name to use for go module, e.g. github.com/user/repo.  "+
		"defaults to the go package of the current working directory.")
	fs.StringVar(&p.project.Domain, "domain", "my.domain", "domain for groups")
}

func (p Plugin) Init() error {
	fmt.Printf("Scaffolding project for project version %q\n", p.Version())
	return nil
}
