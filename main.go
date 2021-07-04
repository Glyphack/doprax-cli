package main

import (
	"github.com/glyphack/doprax-cli/cmd"
	"github.com/glyphack/doprax-cli/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(cmd.NewRootCmd(cmdutil.New()).Execute())
}
