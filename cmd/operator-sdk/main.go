package main

import (
	"fmt"
	"log"

	"sigs.k8s.io/kubebuilder-exp/cmd/kubebuilder/cli"

	"github.com/spf13/cobra"

	"sigs.k8s.io/kubebuilder-exp/pkg/plugin"

	helmv1 "sigs.k8s.io/kubebuilder-exp/cmd/operator-sdk/helm/v1"
	golangv1 "sigs.k8s.io/kubebuilder-exp/pkg/scaffold/golang/v1"
	golangv2 "sigs.k8s.io/kubebuilder-exp/pkg/scaffold/golang/v2"
)

type commandHelp struct {
	name string
	desc string
}

func main() {
	cli.CommandName = "operator-sdk"
	plugin.Register(&golangv1.Scaffolder{})
	plugin.Register(&golangv2.Scaffolder{})
	plugin.Register(&helmv1.Scaffolder{})

	c := cli.New()
	err := c.AddCommand(
		newOLMCmd(),
	)
	if err != nil {
		panic(fmt.Errorf("bug found in operator-sdk (%v); please file an issue!", err))
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
