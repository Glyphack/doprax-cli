package cmd

import (
	"fmt"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/glyphack/doprax-cli/api"
	"github.com/glyphack/doprax-cli/internal/config"
	"github.com/glyphack/doprax-cli/pkg/cmdutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCmdLogin(f *cmdutil.Factory) *cobra.Command {
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "Login into Doprax account",
		Long:  `Login into Doprax account requires API key, username and password it will save obtained token in config file`,
		Run: func(cmd *cobra.Command, args []string) {
			apiKey, err := cmd.Flags().GetString("api-key")
			if apiKey == "" || err != nil {
				apiKey = viper.GetString("api-key")
			}
			if apiKey == "" {
				apiKey = getApiKeyFromPrompt()
			}

			// verify api key is correct
			client := api.NewHttpClient(&api.Config{ApiKey: apiKey})
			_, err = client.GetProjects()
			if err != nil {
				fmt.Printf("Api key is incorrect %s", err.Error())
				return
			}
			viper.Set("api-key", apiKey)

			email, err := cmd.Flags().GetString("email")
			password, err := cmd.Flags().GetString("password")

			if email == "" || password == "" {
				email, password = getEmailPassword()
			}
			result, _ := client.Login(&api.LoginData{Email: email, Password: password})
			if result.Success {
				fmt.Printf("%s Welcome %s\n", result.SuccessMsg, result.Username)
				viper.Set("cca", result.Cca)
			} else {
				fmt.Printf("Login failed %s\n", result.ErrorMsg)
			}

			err = config.WriteConfigFile()
			if err != nil {
				fmt.Printf("Cannot save API key")
			}

		},
	}

	loginCmd.Flags().String("api-key", "", "API key")
	loginCmd.Flags().String("email", "", "Email")
	loginCmd.Flags().String("password", "", "Password")

	return loginCmd
}

func getApiKeyFromPrompt() string {
	promptMessage := &survey.Password{
		Message: "Please enter doprax api key",
	}
	apiKey := ""
	survey.AskOne(promptMessage, &apiKey)
	return apiKey
}

func getEmailPassword() (string, string) {
	emailPromptMessage := &survey.Input{
		Message: "Please enter your email",
	}
	email := ""
	survey.AskOne(emailPromptMessage, &email)
	passPromptMessage := &survey.Password{
		Message: "Please enter your password",
	}
	password := ""
	survey.AskOne(passPromptMessage, &password)
	return email, password
}
