package cmd

import (
	"github.com/ilkerkorkut/rolling-hash-algorithm/internal/logging"
	"github.com/ilkerkorkut/rolling-hash-algorithm/internal/signature"
	"github.com/spf13/cobra"
)

func IsValid(hf string) bool {
	switch hf {
	case "sha1":
		return true
	case "md5":
		return true
	default:
		return false
	}
}

var signatureCmd = &cobra.Command{
	Use:   "signature",
	Short: "Signature of the file",
	Args: func(cmd *cobra.Command, args []string) error {
		if chunkSize == 0 {
			chunkSize = 4
			logging.GetLogger().Infof("chunk size is not set, using default value: %d", chunkSize)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		signature.Run(
			cmd.Flag("file").Value.String(),
			cmd.Flag("signature-file").Value.String(),
			chunkSize,
		)
	},
}

func init() {
	rootCmd.AddCommand(signatureCmd)

	signatureCmd.Flags().StringP("file", "f", "", "File to be signed")
	signatureCmd.Flags().StringP("signature-file", "s", "", "Signature file")
	signatureCmd.MarkFlagsRequiredTogether("file", "signature-file")

	signatureCmd.Flags().IntVarP(&chunkSize, "chunk-size", "c", 0, "Chunk size")
}
