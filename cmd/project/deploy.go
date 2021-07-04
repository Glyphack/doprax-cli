package project

import (
	"fmt"

	"github.com/glyphack/doprax-cli/api"
	"github.com/glyphack/doprax-cli/pkg/cmdutil"
	"github.com/spf13/cobra"
)

type DeployOptions struct {
	projectTitle string
	operation    string
	client       *api.Client
}

func NewCmdDeploy(f *cmdutil.Factory) *cobra.Command {
	deployCmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy project",
		Long: "Deploy project with given operation",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := f.APIClient()
			if err != nil {
				return err
			}
			opts := DeployOptions{operation: args[0], projectTitle: args[1], client: client}
			return DeployRun(&opts)
		},
	}

	return deployCmd
}

func DeployRun(opts *DeployOptions) error {
	project, err := opts.client.GetProjectByTitle(opts.projectTitle)
	if err != nil {
		return err
	}
	response, err := opts.client.DeployProject(project, &api.DeployProjectInput{Operation: opts.operation})
	if err != nil {
		return err
	}
	if response.Success == false {
		return fmt.Errorf("Deploy failed")
	}

	fmt.Printf(response.Msg)
	return nil
}
