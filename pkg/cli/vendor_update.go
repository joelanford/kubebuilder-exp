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

func (c CLI) newVendorUpdateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Update vendor dependencies",
		Long:  `Update vendor dependencies`,
		Example: fmt.Sprintf(`Update the vendor dependencies:
%s update vendor
`, c.commandName),
		RunE: func(cmd *cobra.Command, args []string) error {
			dieIfNoProject()
			fmt.Println("updating vendor dependencies")
			return nil
		},
	}
}
