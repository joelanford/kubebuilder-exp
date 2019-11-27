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

package projutil

import (
	"io/ioutil"
	"os"

	"github.com/alecthomas/gometalinter/_linters/src/gopkg.in/yaml.v2"
	"sigs.k8s.io/kubebuilder-exp/pkg/plugin/golang/project"
	"sigs.k8s.io/kubebuilder-exp/pkg/scaffold/input"
)

// LoadProjectFile reads the project file and deserializes it into a Project
func LoadProjectFile(path string) (input.ProjectFile, error) {
	in, err := ioutil.ReadFile(path) // nolint: gosec
	if err != nil {
		return input.ProjectFile{}, err
	}
	p := input.ProjectFile{}
	err = yaml.Unmarshal(in, &p)
	if err != nil {
		return input.ProjectFile{}, err
	}
	if p.Version == "" {
		// older kubebuilder project does not have scaffolding version
		// specified, so default it to Version1
		p.Version = project.Version1
	}
	return p, nil
}

func IsExistingProject() bool {
	_, err := os.Stat("PROJECT")
	if err != nil {
		return false
	}
	return true
}
