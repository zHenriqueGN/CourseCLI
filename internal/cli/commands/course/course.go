package course

import (
	"github.com/spf13/cobra"
	"github.com/zHenriqueGN/CourseCLI/internal/database"
)

func NewCourseCmd(courseDB *database.CourseDB) *cobra.Command {

	var courseCmd cobra.Command
	courseCmd.Use = "course"
	courseCmd.Short = "course commands"

	return &courseCmd
}
