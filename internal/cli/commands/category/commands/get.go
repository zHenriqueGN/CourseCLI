package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zHenriqueGN/CourseCLI/internal/cli/cmdtypes"
	"github.com/zHenriqueGN/CourseCLI/internal/database"
	"github.com/zHenriqueGN/CourseCLI/internal/entity"
)

var categoryId string

func NewGetCmd(categoryDB *database.CategoryDB) *cobra.Command {
	var getCmd cobra.Command
	getCmd.Use = "get"
	getCmd.Short = "get a category"
	getCmd.Flags().StringVar(&categoryId, "id", "", "")
	getCmd.MarkFlagRequired("id")
	getCmd.RunE = runGet(categoryDB)
	return &getCmd
}

func runGet(categoryDB *database.CategoryDB) cmdtypes.RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		category, err := categoryDB.GetCategory(categoryId)
		if err != nil {
			return err
		}
		printCategory(category)
		return nil
	}
}

func printCategory(category *entity.Category) {
	lines := []string{
		"____________________________________________",
		"ID: %s",
		"NAME: %s",
		"DESCRIPTION: %s\n",
	}
	joinedLines := strings.Join(lines, "\n")
	fmt.Printf(joinedLines, category.ID, category.Name, category.Description)
}
