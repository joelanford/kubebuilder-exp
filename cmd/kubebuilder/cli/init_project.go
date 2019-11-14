package cli

import (
	"fmt"
	"strconv"
	"strings"

	"sigs.k8s.io/kubebuilder-exp/pkg/plugin"

	"github.com/spf13/cobra"
	"sigs.k8s.io/kubebuilder-exp/pkg/project"
)

var (
	DefaultProjectVersion = "2"
)

/*

 */
func newInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a new project",
		Long: `Initialize a new project.

For further help about a specific project version, set --project-version.
`,
		Run:     func(cmd *cobra.Command, args []string) {},
		Example: getInitProjectVersionExamples(),
	}
	// Register --project-version on the dynamically created command
	// so that it shows up in help and does not cause a parse error.
	cmd.Flags().String("project-version", DefaultProjectVersion, fmt.Sprintf("project version, (%s)", strings.Join(getProjectVersions(), ", ")))

	projectVersion, isBaseHelp := project.GetVersionFromFlag(DefaultProjectVersion)

	if !isBaseHelp {
		if s, ok := plugin.Lookup(projectVersion); !ok {
			err := fmt.Errorf("no scaffolder for project version %q", projectVersion)
			cmd.RunE = errCmdFunc(err)
		} else if ps, ok := s.(plugin.ProjectScaffolder); !ok {
			err := fmt.Errorf("scaffolder for project version %q does not support project initialization", projectVersion)
			cmd.RunE = errCmdFunc(err)
		} else {
			ps.BindProjectFlags(cmd.Flags())
			cmd.Long = ps.ProjectHelp()
			cmd.Example = ps.ProjectExample()
			cmd.RunE = func(cmd *cobra.Command, args []string) error {
				if err := ps.ScaffoldProject(); err != nil {
					return fmt.Errorf("failed to scaffold project with version %q: %v", projectVersion, err)
				}
				return nil
			}
		}
	}

	return cmd
}

func getProjectVersions() (projectVersions []string) {
	for _, s := range plugin.List() {
		if ps, ok := s.(plugin.ProjectScaffolder); ok {
			projectVersions = append(projectVersions, strconv.Quote(ps.Version()))
		}
	}
	return projectVersions
}

func getInitProjectVersionExamples() string {
	exampleBuilder := strings.Builder{}
	for _, sv := range getProjectVersions() {
		exampleBuilder.WriteString(fmt.Sprintf(`  # Help for initializing a project with version %s
  %s init --project-version=%s -h

`, sv, CommandName, sv))
	}
	return strings.TrimSuffix(exampleBuilder.String(), "\n")
}
