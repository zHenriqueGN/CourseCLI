package commands

import (
	"github.com/spf13/cobra"
	"github.com/zHenriqueGN/CourseCLI/internal/cli/cmdtypes"
	"github.com/zHenriqueGN/CourseCLI/internal/database"
	"github.com/zHenriqueGN/CourseCLI/internal/entity"
)

var (
	name        string
	description string
)

func NewCreateCmd(categoryDB *database.CategoryDB) *cobra.Command {
	var createCmd cobra.Command
	createCmd.Use = "create"
	createCmd.Short = "create a new category"
	createCmd.Flags().StringVarP(&name, "name", "n", "", "category name")
	createCmd.Flags().StringVarP(&description, "description", "d", "", "category description")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("description")
	createCmd.RunE = runCreate(categoryDB)
	return &createCmd
}

func runCreate(categoryDB *database.CategoryDB) cmdtypes.RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		category := entity.NewCategory(name, description)
		err := categoryDB.Create(category)
		if err != nil {
			return err
		}
		printCategory(category)
		return nil
	}
}
