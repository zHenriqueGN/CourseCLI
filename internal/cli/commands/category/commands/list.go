package commands

import (
	"github.com/spf13/cobra"
	"github.com/zHenriqueGN/CourseCLI/internal/cli/cmdtypes"
	"github.com/zHenriqueGN/CourseCLI/internal/database"
)

func NewListCmd(categoryDB *database.CategoryDB) *cobra.Command {
	var listCmd cobra.Command
	listCmd.Use = "list"
	listCmd.Short = "list the categories"
	listCmd.RunE = runList(categoryDB)
	return &listCmd
}

func runList(categoryDB *database.CategoryDB) cmdtypes.RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		categories, err := categoryDB.FindAll()
		if err != nil {
			return err
		}
		for _, category := range categories {
			printCategory(category)
		}
		return nil
	}
}
