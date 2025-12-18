/*
Copyright © 2025 Edwin S
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// stampCmd represents the stamp command
var stampCmd = &cobra.Command{
	Use:   "stamp [commit-hash]",
	Short: "Stamp a git commit on the Ethereum blockchain",
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

		return s.Stamp(ctx, commitHash)
	},
}

func init() {
	rootCmd.AddCommand(stampCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stampCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stampCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
