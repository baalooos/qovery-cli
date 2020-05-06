package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"qovery.go/io"
)

var environmentStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Deploy and start the current environment",
	Long: `START deploy and turn on an environment. For example:

	qovery environment start`,

	Run: func(cmd *cobra.Command, args []string) {
		if !hasFlagChanged(cmd) {
			BranchName = io.CurrentBranchName()
			qoveryYML, err := io.CurrentQoveryYML()
			if err != nil {
				io.PrintError("No qovery configuration file found")
				os.Exit(1)
			}
			ProjectName = qoveryYML.Application.Project
			ApplicationName = qoveryYML.Application.GetSanitizeName()
		}

		projectId := io.GetProjectByName(ProjectName).Id
		application := io.GetApplicationByName(projectId, BranchName, ApplicationName)
		environment := io.GetEnvironmentByName(projectId, BranchName)

		io.Deploy(projectId, environment.Id, application.Id, application.Repository.CommitId)
		ShowDeploymentMessage()
	},
}

func init() {
	environmentStartCmd.PersistentFlags().StringVarP(&ProjectName, "project", "p", "", "Your project name")
	environmentStartCmd.PersistentFlags().StringVarP(&BranchName, "branch", "b", "", "Your branch name")
	environmentStartCmd.PersistentFlags().StringVarP(&ApplicationName, "application", "a", "", "Your application name")

	environmentCmd.AddCommand(environmentStartCmd)
}
