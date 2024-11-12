package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Привет. Это консольное to-do приложение!",
	Long:  "В этом консольном приложении ты можешь добавлять, изменять и удалять задачи",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Для помощи используй флаг --help")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
