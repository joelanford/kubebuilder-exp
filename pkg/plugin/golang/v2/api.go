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

package v2

import (
	"fmt"
	"os"

	"sigs.k8s.io/kubebuilder-exp/pkg/plugin"

	"github.com/spf13/pflag"
)

var _ plugin.CreateAPIPlugin = &Plugin{}

func (_ Plugin) CreateAPIDescription() string {
	return `Scaffold a Kubernetes API by creating a Resource definition and / or a Controller.

create resource will prompt the user for if it should scaffold the Resource and / or Controller.  To only
scaffold a Controller for an existing Resource, select "n" for Resource.  To only define
the schema for a Resource without writing a Controller, select "n" for Controller.

After the scaffold is written, api will run make on the project.
`
}

func (p Plugin) CreateAPIExample() string {
	return fmt.Sprintf(`  # Create a frigates API with Group: ship, Version: v1beta1 and Kind: Frigate
  %s create api --group ship --version v1beta1 --kind Frigate
  
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
`, p.commandName)
}

func (p *Plugin) BindCreateAPIFlags(fs *pflag.FlagSet) {
	fs.BoolVar(&p.runMake, "make", true,
		"if true, run make after generating files")
	//fs.BoolVar(&p.apiScaffolder.DoResource, "resource", true,
	//	"if set, generate the resource without prompting the user")
	p.resourceFlag = fs.Lookup("resource")
	//fs.BoolVar(&p.apiScaffolder.DoController, "controller", true,
	//	"if set, generate the controller without prompting the user")
	p.controllerFlag = fs.Lookup("controller")
	if os.Getenv("KUBEBUILDER_ENABLE_PLUGINS") != "" {
		fs.StringVar(&p.pattern, "pattern", "",
			"generates an API following an extension pattern (addon)")
	}
	//p.apiScaffolder.Resource = resourceForFlags(fs)
}

func (p Plugin) CreateAPI() error {
	fmt.Printf("Scaffolding API for project version %q\n", p.Version())
	return nil
}
