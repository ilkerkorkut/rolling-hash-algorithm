package delta

import (
	"os"

	"github.com/ilkerkorkut/rolling-hash-algorithm/internal/logging"
	"github.com/ilkerkorkut/rolling-hash-algorithm/internal/signature"
)

func Run(signatureFileName string, updatedFileName string, deltaFileName string, chunkSize int) {
	sigFile, err := os.ReadFile(signatureFileName)
	if err != nil {
		logging.GetLogger().Fatalf("cannot read signature file: %s", err)
	}

	sig, err := signature.UnmarshalJSON(sigFile)
	if err != nil {
		logging.GetLogger().Fatalf("cannot unmarshal signature: %s", err)
	}

	updatedFile, err := os.Open(updatedFileName)
	if err != nil {
		logging.GetLogger().Fatalf("cannot open updated file: %s", err)
	}
	defer updatedFile.Close()

	deltaFile, err := os.Create(deltaFileName)
	if err != nil {
		logging.GetLogger().Fatalf("cannot create delta file: %s", err)
	}
	defer deltaFile.Close()

	dp := NewDeltaProcessor(sig, updatedFile, true)
	d := dp.BuildDelta()

	deltaBytes, err := d.MarshalJSON()
	if err != nil {
		logging.GetLogger().Fatalf("cannot marshal delta: %s", err)
	}

	err = os.WriteFile(deltaFileName, deltaBytes, 0644)
	if err != nil {
		logging.GetLogger().Fatalf("cannot write delta file: %s", err)
	}

	logging.GetLogger().Infof("delta file created: %s", deltaFileName)

	GenerateDeltaReport(d)
}
