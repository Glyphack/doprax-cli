package project

import (
	"github.com/glyphack/doprax-cli/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmdProject(f *cmdutil.Factory) *cobra.Command {
	projectCmd := &cobra.Command{
		Use:   "project",
		Short: "Manage projects",
		Long:  "Work with projects",
	}

	projectCmd.AddCommand(NewCmdPull(f))
	projectCmd.AddCommand(NewCmdList(f))
	projectCmd.AddCommand(NewCmdDeploy(f))

	return projectCmd
}
