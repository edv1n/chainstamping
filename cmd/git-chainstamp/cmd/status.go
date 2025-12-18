/*
Copyright © 2025 Edwin S
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status [commit-hash]",
	Short: "Get the stamping status of a git commit",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		commitHash := "HEAD"
		if len(args) > 0 {
			commitHash = args[0]
		}

		s, err := newStamper()
		if err != nil {
			return fmt.Errorf("Failed to create stamper: %w", err)
		}

		return s.Status(ctx, commitHash)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
