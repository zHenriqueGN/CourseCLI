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
	categoryId  string
)

func NewCreateCmd(courseDB *database.CourseDB) *cobra.Command {
	var createCmd cobra.Command
	createCmd.Use = "create"
	createCmd.Short = "create a course"
	createCmd.Flags().StringVarP(&name, "name", "n", "", "course name")
	createCmd.Flags().StringVarP(&description, "description", "d", "", "course description")
	createCmd.Flags().StringVar(&categoryId, "category-id", "", "category id")
	createCmd.MarkFlagRequired("name")
	createCmd.MarkFlagRequired("description")
	createCmd.MarkFlagRequired("category-id")
	createCmd.RunE = runCreate(courseDB)
	return &createCmd
}

func runCreate(courseDB *database.CourseDB) cmdtypes.RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		course := entity.NewCourse(name, description, categoryId)
		err := courseDB.Create(course)
		if err != nil {
			return err
		}
		printCourse(course)
		return nil
	}
}
