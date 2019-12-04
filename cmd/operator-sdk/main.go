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

package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"sigs.k8s.io/kubebuilder-exp/pkg/cli"
	golangv1 "sigs.k8s.io/kubebuilder-exp/pkg/plugin/golang/v1"
	golangv2 "sigs.k8s.io/kubebuilder-exp/pkg/plugin/golang/v2"
	helmv1 "sigs.k8s.io/kubebuilder-exp/pkg/plugin/helm/v1"
)

func main() {
	c, err := cli.New(
		cli.WithCommandName("operator-sdk"),
		cli.WithDefaultProjectVersion("2"),
		cli.WithExtraCommands(newOLMCmd()),
		cli.WithPlugins(
			&golangv1.Plugin{},
			&golangv2.Plugin{},
			&helmv1.Plugin{},
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := c.Run(); err != nil {
		log.Fatal(err)
	}
}

func newOLMCmd() *cobra.Command {
	olmCmd := &cobra.Command{
		Use:   "olm",
		Short: "Install OLM, generate CSVs, or run your operator with OLM",
	}
	olmCmd.AddCommand(
		&cobra.Command{
			Use:   "install",
			Short: "Install Operator Lifecycle Manager in your cluster",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("not implemented")
			},
		},
		&cobra.Command{
			Use:   "gen-csv",
			Short: "Generate a ClusterServiceVersion for your operator",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("not implemented")
			},
		},
		&cobra.Command{
			Use:   "up",
			Short: "Run your operator using your CSV with OLM",
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("not implemented")
			},
		},
	)
	return olmCmd
}
