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

package cli

import (
	"fmt"

	"sigs.k8s.io/kubebuilder-exp/pkg/plugin/golang/project"

	"sigs.k8s.io/kubebuilder-exp/pkg/plugin"

	"github.com/spf13/cobra"
)

func (c CLI) newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Scaffold a Kubernetes API or webhook",
		Long:  `Scaffold a Kubernetes API or webhook.`,
	}
	cmd.AddCommand(c.newCreateAPICmd())

	foundProject, projectVersion := getProjectVersion()
	if !foundProject || projectVersion != project.Version1 {
		cmd.AddCommand(c.newCreateWebhookCmd())
	}

	return cmd
}

func (c CLI) newCreateAPICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api",
		Short: "Scaffold a Kubernetes API",
		Long: `Scaffold a Kubernetes API.
`,
		RunE: errCmdFunc(
			fmt.Errorf("create api subcommand requires an existing project"),
		),
	}

	found, projectVersion := getProjectVersion()
	if !found {
		msg := `For project-specific information, run this command in the root directory of a
project.
`
		cmd.Long = fmt.Sprintf("%s\n%s", cmd.Long, msg)
		return cmd
	}

	// Lookup the plugin for projectVersion and bind it to the command.
	c.bindCreateAPIPlugin(cmd, projectVersion)
	return cmd
}

func (c CLI) bindCreateAPIPlugin(cmd *cobra.Command, projectVersion string) {
	p, ok := c.plugins[projectVersion]
	if !ok {
		err := fmt.Errorf("unknown project version %q", projectVersion)
		cmdErr(cmd, err)
		return
	}
	as, ok := p.(plugin.APIScaffolder)
	if !ok {
		err := fmt.Errorf("plugin for project version %q does not support API creation", projectVersion)
		cmdErr(cmd, err)
		return
	}

	as.BindAPIFlags(cmd.Flags())
	cmd.Long = as.APIHelp()
	cmd.Example = as.APIExample()
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if err := as.ScaffoldAPI(); err != nil {
			return fmt.Errorf("failed to create api for project with version %q: %v", projectVersion, err)
		}
		return nil
	}
}
func (c CLI) newCreateWebhookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "webhook",
		Short: "Scaffold a webhook for an API resource",
		Long: `Scaffold a webhook for an API resource.
`,
		RunE: errCmdFunc(
			fmt.Errorf("create webhook subcommand requires an existing project"),
		),
	}

	found, projectVersion := getProjectVersion()
	if !found {
		msg := `For project-specific information, run this command in the root directory of a
project.
`
		cmd.Long = fmt.Sprintf("%s\n%s", cmd.Long, msg)
		return cmd
	}

	// Lookup the plugin for projectVersion and bind it to the command.
	c.bindCreateWebhookPlugin(cmd, projectVersion)
	return cmd
}

func (c CLI) bindCreateWebhookPlugin(cmd *cobra.Command, projectVersion string) {
	p, ok := c.plugins[projectVersion]
	if !ok {
		err := fmt.Errorf("unknown project version %q", projectVersion)
		cmdErr(cmd, err)
		return
	}
	ws, ok := p.(plugin.WebhookScaffolder)
	if !ok {
		err := fmt.Errorf("plugin for project version %q does not support webhook creation", projectVersion)
		cmdErr(cmd, err)
		return
	}

	ws.BindWebhookFlags(cmd.Flags())
	cmd.Long = ws.WebhookHelp()
	cmd.Example = ws.WebhookExample()
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if err := ws.ScaffoldWebhook(); err != nil {
			return fmt.Errorf("failed to create webhook for project with version %q: %v", projectVersion, err)
		}
		return nil
	}
}
