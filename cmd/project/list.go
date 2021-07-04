package project

import (
	"fmt"

	"github.com/glyphack/doprax-cli/api"
	"github.com/glyphack/doprax-cli/pkg/cmdutil"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	client *api.Client
}

func NewCmdList(f *cmdutil.Factory) *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List projects",
		Long:  "List available projects in doprax dashboard",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := f.APIClient()
			if err != nil {
				return err
			}
			opts := ListOptions{client: client}
			return ListRun(&opts)
		},
	}

	return listCmd
}

func ListRun(opts *ListOptions) error {
	projects, err := opts.client.GetProjects()
	if err != nil {
		return err
	}
	for _, project := range *projects {
		fmt.Println(project.Title)
	}

	return nil
}
