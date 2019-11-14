package project

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/pflag"

	"github.com/alecthomas/gometalinter/_linters/src/gopkg.in/yaml.v2"
	"sigs.k8s.io/kubebuilder-exp/pkg/scaffold/golang/project"
	"sigs.k8s.io/kubebuilder-exp/pkg/scaffold/input"
)

var ErrProjectNotFound = errors.New("project not found")

func GetVersionFromFlag(defaultVersion string) (projectVersion string, isHelp bool) {
	fs := pflag.NewFlagSet("project-version", pflag.ExitOnError)
	fs.ParseErrorsWhitelist = pflag.ParseErrorsWhitelist{UnknownFlags: true}

	fs.StringVar(&projectVersion, "project-version", defaultVersion, "project version")
	fs.BoolVarP(&isHelp, "help", "h", false, "print help")

	err := fs.Parse(os.Args[1:])
	isHelp = err != nil || (isHelp && !fs.Lookup("project-version").Changed)
	return projectVersion, isHelp
}

func GetVersionFromFile() (string, error) {
	if _, err := os.Stat("PROJECT"); os.IsNotExist(err) {
		return "", ErrProjectNotFound
	}
	projectInfo, err := LoadProjectFile("PROJECT")
	if err != nil {
		return "", fmt.Errorf("failed to read the PROJECT file: %w", err)
	}
	if projectInfo.Version == "" {
		return "", ErrProjectNotFound
	}
	return projectInfo.Version, nil
}

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
