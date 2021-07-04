package project

import (
	"fmt"

	"github.com/glyphack/doprax-cli/api"
	"github.com/glyphack/doprax-cli/pkg/cmdutil"
	"github.com/spf13/cobra"
)

type PullOptions struct {
	projectTitle string
	client       *api.Client
}

func NewCmdPull(f *cmdutil.Factory) *cobra.Command {
	pullCmd := &cobra.Command{
		Use:   "pull",
		Short: "Pull project code from selected source",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := f.APIClient()
			if err != nil {
				return err
			}
			opts := PullOptions{projectTitle: args[0], client: client}
			return PullRun(&opts)
		},
	}

	return pullCmd
}

func PullRun(opts *PullOptions) error {
	project, err := opts.client.GetProjectByTitle(opts.projectTitle)
	if err != nil {
		return err
	}
	response, err := opts.client.PullProject(project)
	if err != nil {
		return err
	}
	if response.Success == false {
		return fmt.Errorf("Pull failed")
	}

	fmt.Printf("Pull initiated")
	return nil
}
