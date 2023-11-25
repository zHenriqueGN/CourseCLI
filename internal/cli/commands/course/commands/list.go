package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zHenriqueGN/CourseCLI/internal/cli/cmdtypes"
	"github.com/zHenriqueGN/CourseCLI/internal/database"
	"github.com/zHenriqueGN/CourseCLI/internal/entity"
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

func printCourse(course *entity.Course) {
	lines := []string{
		"____________________________________________",
		"ID: %s",
		"NAME: %s",
		"DESCRIPTION: %s",
		"CATEGORY ID: %s\n",
	}
	joinedLines := strings.Join(lines, "\n")
	fmt.Printf(joinedLines, course.ID, course.Name, course.Description, course.CategoryID)
}
