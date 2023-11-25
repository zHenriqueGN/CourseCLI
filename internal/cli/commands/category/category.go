package category

import (
	"github.com/spf13/cobra"
	"github.com/zHenriqueGN/CourseCLI/internal/cli/commands/category/commands"
	"github.com/zHenriqueGN/CourseCLI/internal/database"
)

func NewCategoryCmd(categoryDB *database.CategoryDB) *cobra.Command {
	createCmd := commands.NewCreateCmd(categoryDB)
	listCmd := commands.NewListCmd(categoryDB)

	var categoryCmd cobra.Command
	categoryCmd.Use = "category"
	categoryCmd.Short = "category commands"
	categoryCmd.AddCommand(createCmd, listCmd)

	return &categoryCmd
}
