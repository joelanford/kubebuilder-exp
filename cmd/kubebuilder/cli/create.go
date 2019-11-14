package cli

import (
	"fmt"

	"sigs.k8s.io/kubebuilder-exp/pkg/plugin"
	"sigs.k8s.io/kubebuilder-exp/pkg/project"

	"github.com/spf13/cobra"
)

type createCommandNotSupportedError struct {
	ProjectVersion string
	CommandName    string
}

func (err createCommandNotSupportedError) Error() string {
	return fmt.Sprintf("plugin for project version %q does not support command \"create %s\"", err.ProjectVersion, err.CommandName)
}

func newCreateCmd() *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Scaffold a Kubernetes API or webhook",
	}

	createCmd.AddCommand(
		newCreateAPICmd(),
		newCreateWebhookCmd(),
	)

	return createCmd
}

func newCreateAPICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api",
		Short: "Scaffold a Kubernetes API",
		Long: `Scaffold a Kubernetes API

This command must be run in the root directory of a project that supports API scaffolding.
`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	if s, projectVersion, err := loadPluginForProject(); err != nil {
		cmd.RunE = errCmdFunc(err)
	} else if as, ok := s.(plugin.APIScaffolder); !ok {
		err := createCommandNotSupportedError{ProjectVersion: projectVersion, CommandName: cmd.Name()}
		cmd.RunE = errCmdFunc(err)
	} else {
		as.BindAPIFlags(cmd.Flags())
		cmd.Long = as.APIHelp()
		cmd.Example = as.APIExample()
		cmd.RunE = func(cmd *cobra.Command, args []string) error {
			if err := as.ScaffoldAPI(); err != nil {
				return fmt.Errorf("failed to scaffold API for project with version %q: %w", projectVersion, err)
			}
			return nil
		}
	}
	return cmd
}

func newCreateWebhookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "webhook",
		Short: "Scaffold a Kubernetes webhook",
		Long: `Scaffold a Kubernetes webhook

This command must be run in the root directory of a project that supports webhook scaffolding.
`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	if s, projectVersion, err := loadPluginForProject(); err != nil {
		cmd.RunE = errCmdFunc(err)
	} else if ws, ok := s.(plugin.WebhookScaffolder); !ok {
		err := createCommandNotSupportedError{ProjectVersion: projectVersion, CommandName: cmd.Name()}
		cmd.RunE = errCmdFunc(err)
	} else {
		ws.BindWebhookFlags(cmd.Flags())
		cmd.Long = ws.WebhookHelp()
		cmd.Example = ws.WebhookExample()
		cmd.RunE = func(cmd *cobra.Command, args []string) error {
			if err := ws.ScaffoldWebhook(); err != nil {
				return fmt.Errorf("failed to scaffold webhook for project with version %q: %w", projectVersion, err)
			}
			return nil
		}
	}
	return cmd
}

func loadPluginForProject() (plugin.Plugin, string, error) {
	projectVersion, err := project.GetVersionFromFile()
	if err != nil {
		return nil, "", err
	}

	p, ok := plugin.Lookup(projectVersion)
	if !ok {
		return nil, "", fmt.Errorf("failed to lookup plugin for project version %q", projectVersion)
	}
	return p, projectVersion, nil
}
