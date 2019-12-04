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

	"github.com/spf13/cobra"
)

// newAlphaCnd returns alpha subcommand which will be mounted
// at the root command by the caller.
func (c CLI) newAlphaCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alpha",
		Short: "Expose commands which are in experimental or early stages of development",
		Long:  `Command group for commands which are either experimental or in early stages of development`,
		Example: fmt.Sprintf(`
# scaffolds webhook server
%s alpha webhook <params>
`, c.commandName),
	}

	cmd.AddCommand(
		c.newWebhookCmd(),
	)
	return cmd
}

func (c CLI) newWebhookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "webhook",
		Short: "Scaffold a webhook server",
		Long: `Scaffold a webhook server if there is no existing server.
Scaffolds webhook handlers based on group, version, kind and other user inputs.
This command is only available for v1 scaffolding project.
`,
		Example: fmt.Sprintf(`	# Create webhook for CRD of group crew, version v1 and kind FirstMate.
	# Set type to be mutating and operations to be create and update.
	%s alpha webhook --group crew --version v1 --kind FirstMate --type=mutating --operations=create,update
`, c.commandName),
		RunE: func(cmd *cobra.Command, args []string) error {
			dieIfNoProject()
			fmt.Printf("Scaffolding alpha webhook for project\n")
			return nil
		},
	}
	return cmd
}
