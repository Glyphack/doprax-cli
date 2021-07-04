package cmd

import (
	"log"

	"github.com/glyphack/doprax-cli/cmd/project"
	"github.com/glyphack/doprax-cli/internal/config"
	"github.com/glyphack/doprax-cli/pkg/cmdutil"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

func NewRootCmd(f *cmdutil.Factory) *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:   "doprax",
		Short: "Doprax CLI tool",
		Long:  `Doprax CLI is a tool to interact with Doprax dashboard, list current apps or pull and deploy apps and more`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.doprax.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Add commands
	rootCmd.AddCommand(NewCmdLogin(f))
	rootCmd.AddCommand(project.NewCmdProject(f))

	return rootCmd
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	configDir, err := config.GetConfigDir()
	if err != nil {
		log.Panic(err)
	}
	viper.AddConfigPath(configDir)
	viper.SetConfigName(config.ConfigFileName)
	viper.SetConfigType(config.ConfigFileType)

	viper.SetDefault("host", "https://www.doprax.com")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Cannot read config file", err)
	}
}
