package cmd

import (
	"os"

	"github.com/ilkerkorkut/rolling-hash-algorithm/internal/logging"
	"github.com/ilkerkorkut/rolling-hash-algorithm/internal/options"
	"github.com/ilkerkorkut/rolling-hash-algorithm/internal/version"
	"github.com/spf13/cobra"
)

var chunkSize int

var (
	//nolint:unused
	opts *options.RHAOptions
	ver  = version.Get()
)

var rootCmd = &cobra.Command{
	Use:     "rha",
	Short:   "Rolling hash algorithm for file comparison",
	Long:    ``,
	Version: ver.GitVersion,
	Run: func(cmd *cobra.Command, args []string) {
		logging.GetLogger().
			WithField("appVersion", ver.GitVersion).
			WithField("goVersion", ver.GoVersion).
			WithField("goOS", ver.GoOs).
			WithField("goArch", ver.GoArch).
			WithField("gitCommit", ver.GitCommit).
			WithField("buildDate", ver.BuildDate).
			Info("RHA is started")
	},
}

func init() {
	opts = options.GetRHAOptions()
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
