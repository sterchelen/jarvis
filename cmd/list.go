package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	inventoryPath string
)

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listRepositoryCmd)

	listRepositoryCmd.Flags().
		StringVarP(&inventoryPath, "inventory", "i", "", "Inventory path")
	viper.BindPFlag("inventory_path", listRepositoryCmd.Flags().Lookup("inventory"))
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List command",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var listRepositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "List ansible repository",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
