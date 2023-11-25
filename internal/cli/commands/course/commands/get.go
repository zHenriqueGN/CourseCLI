package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zHenriqueGN/CourseCLI/internal/cli/cmdtypes"
	"github.com/zHenriqueGN/CourseCLI/internal/database"
	"github.com/zHenriqueGN/CourseCLI/internal/entity"
)

var courseId string

func NewGetCmd(courseDB *database.CourseDB) *cobra.Command {
	var getCmd cobra.Command
	getCmd.Use = "get"
	getCmd.Short = "get a course"
	getCmd.Flags().StringVar(&courseId, "id", "", "course id")
	getCmd.MarkFlagRequired("id")
	getCmd.RunE = runGet(courseDB)
	return &getCmd
}

func runGet(courseDB *database.CourseDB) cmdtypes.RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		course, err := courseDB.FindByID(courseId)
		if err != nil {
			return err
		}
		printCourse(course)
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
