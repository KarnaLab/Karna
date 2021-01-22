package cmd

import (
	"os"

	deploy "github.com/karnalab/karna/deploy"
	"github.com/spf13/cobra"
)

var logger *deploy.KarnaLogger
var rootCmd = &cobra.Command{Use: "karna"}

var cmdDeploy = &cobra.Command{
	Use:   "deploy",
	Short: "Use Karna Deployment to deploy your Lambda application.",
	Long: `Karna Deployment will build and deploy your Lambda function 
	on top of your config file.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		logger.Log("Deployment in progress...")
	},
	Run: func(cmd *cobra.Command, args []string) {
		functionName, _ := cmd.Flags().GetString("function-name")
		alias, _ := cmd.Flags().GetString("alias")

		if elapsed, err := deploy.Run(&functionName, &alias); err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		} else {
			logger.Log("Completed in " + elapsed)
		}
	},
}

func init() {
	var functionName string
	var alias string

	cmdDeploy.Flags().StringVarP(&functionName, "function-name", "f", "", "Function to deploy (JSON key into your config file)")
	cmdDeploy.Flags().StringVarP(&alias, "alias", "a", "", "Alias to publish")

	cmdDeploy.MarkFlagRequired("function-name")
	cmdDeploy.MarkFlagRequired("alias")

	rootCmd.AddCommand(cmdDeploy)
}

//Execute => Will register commands && execute the right one.
func Execute() {
	rootCmd.Execute()
}
