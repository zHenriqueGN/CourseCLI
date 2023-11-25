package commands

import (
	"github.com/spf13/cobra"
	"github.com/zHenriqueGN/CourseCLI/internal/cli/cmdtypes"
	"github.com/zHenriqueGN/CourseCLI/internal/database"
)

func NewListCmd(courseDB *database.CourseDB) *cobra.Command {
	var listCmd cobra.Command
	listCmd.Use = "list"
	listCmd.Short = "list the courses"
	listCmd.RunE = runList(courseDB)
	return &listCmd
}

func runList(courseDB *database.CourseDB) cmdtypes.RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		courses, err := courseDB.FindAll()
		if err != nil {
			return err
		}
		for _, course := range courses {
			printCourse(&course)
		}
		return nil
	}
}
