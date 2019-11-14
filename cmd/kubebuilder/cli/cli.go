package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	CommandName = "kubebuilder"
)

type CLI struct {
	cmd *cobra.Command
}

func New() *CLI {
	return &CLI{
		newRootCmd(),
	}
}

func (c CLI) Run() error {
	return c.cmd.Execute()
}

func (c *CLI) AddCommand(cmds ...*cobra.Command) error {
	for _, cmd := range cmds {
		for _, subc := range c.cmd.Commands() {
			if cmd.Name() == subc.Name() {
				return fmt.Errorf("command %q already exists", cmd.Name())
			}
		}
		c.cmd.AddCommand(cmd)
	}
	return nil
}

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: CommandName,
	}

	rootCmd.AddCommand(
		newInitCmd(),
		newCreateCmd(),
	)

	return rootCmd
}

func errCmdFunc(err error) func(*cobra.Command, []string) error {
	return func(command *cobra.Command, args []string) error {
		return err
	}
}
