package cmd

import (
	"github.com/ilkerkorkut/rolling-hash-algorithm/internal/delta"
	"github.com/ilkerkorkut/rolling-hash-algorithm/internal/logging"
	"github.com/spf13/cobra"
)

var deltaCmd = &cobra.Command{
	Use:   "delta",
	Short: "Delta of the files",
	Args: func(cmd *cobra.Command, args []string) error {
		if chunkSize == 0 {
			chunkSize = 4
			logging.GetLogger().Infof("chunk size is not set, using default value: %d", chunkSize)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		delta.Run(
			cmd.Flag("signature-file").Value.String(),
			cmd.Flag("new-file").Value.String(),
			cmd.Flag("delta-file").Value.String(),
			chunkSize,
		)
	},
}

func init() {
	rootCmd.AddCommand(deltaCmd)

	deltaCmd.Flags().StringP("signature-file", "s", "", "Signature file")
	deltaCmd.Flags().StringP("new-file", "n", "", "New file")
	deltaCmd.Flags().StringP("delta-file", "d", "", "Delta file")
	deltaCmd.MarkFlagsRequiredTogether("signature-file", "new-file", "delta-file")

	deltaCmd.Flags().IntVarP(&chunkSize, "chunk-size", "c", 0, "Chunk size")
}
