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

func NewListCmd(categoryDB *database.CategoryDB) *cobra.Command {
	var listCmd cobra.Command
	listCmd.Use = "list"
	listCmd.Short = "list the categories"
	listCmd.Flags().StringVar(&categoryId, "category-id", "", "category id")
	listCmd.RunE = runList(categoryDB)
	return &listCmd
}

func runList(categoryDB *database.CategoryDB) cmdtypes.RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		if categoryId != "" {
			category, err := categoryDB.GetCategory(categoryId)
			if err != nil {
				return err
			}
			printCategory(category)
			return nil
		}
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
