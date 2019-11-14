package main

import (
	"log"

	"sigs.k8s.io/kubebuilder-exp/pkg/plugin"

	"sigs.k8s.io/kubebuilder-exp/cmd/kubebuilder/cli"

	golangv1 "sigs.k8s.io/kubebuilder-exp/pkg/scaffold/golang/v1"
	golangv2 "sigs.k8s.io/kubebuilder-exp/pkg/scaffold/golang/v2"
)

func main() {
	plugin.Register(&golangv1.Scaffolder{})
	plugin.Register(&golangv2.Scaffolder{})

	c := cli.New()
	if err := c.Run(); err != nil {
		log.Fatal(err)
	}
}
