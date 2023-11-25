package course

import (
	"github.com/spf13/cobra"
	"github.com/zHenriqueGN/CourseCLI/internal/cli/commands/course/commands"
	"github.com/zHenriqueGN/CourseCLI/internal/database"
)

func NewCourseCmd(courseDB *database.CourseDB) *cobra.Command {
	createCmd := commands.NewCreateCmd(courseDB)
	listCmd := commands.NewListCmd(courseDB)

	var courseCmd cobra.Command
	courseCmd.Use = "course"
	courseCmd.Short = "course commands"
	courseCmd.AddCommand(createCmd, listCmd)

	return &courseCmd
}
