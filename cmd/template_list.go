package cmd

import (
	"github.com/spf13/cobra"
	"qovery.go/util"
)

var templateListCmd = &cobra.Command{
	Use:   "list",
	Short: "List templates",
	Long: `LIST show all available templates. For example:

	qovery template list`,

	Run: func(cmd *cobra.Command, args []string) {
		table := util.GetTable()
		table.SetHeader([]string{"name", "description"})

		templates := util.ListAvailableTemplates()

		if len(templates) == 0 {
			table.Append([]string{"", ""})
		} else {
			for _, t := range templates {
				table.Append([]string{t.Name, t.Description})
			}
		}

		table.Render()
	},
}

func init() {
	templateCmd.AddCommand(templateListCmd)
}
