package cli

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"github.com/zHenriqueGN/CourseCLI/internal/cli/commands/category"
	"github.com/zHenriqueGN/CourseCLI/internal/database"
)

func Execute() {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	categoryDB := database.NewCategory(db)
	categoryCmd := category.NewCategoryCmd(categoryDB)

	var rootCmd cobra.Command
	rootCmd.Use = "course-cli"
	rootCmd.Short = "course-cli is a CLI for managing courses"
	rootCmd.AddCommand(categoryCmd)

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
