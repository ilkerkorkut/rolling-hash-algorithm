package signature

import (
	"os"

	"github.com/ilkerkorkut/rolling-hash-algorithm/internal/logging"
)

func Run(fileName string, signatureFileName string, chunkSize int) {
	file, err := os.Open(fileName)
	if err != nil {
		logging.GetLogger().Fatalf("cannot open file: %s", err)
	}
	defer file.Close()

	sg := NewSignatureGenerator(file, chunkSize, true)

	signature := sg.GenerateSigature()

	logging.GetLogger().Infof("signature generated: %v", signature)

	sgMrsh, err := signature.MarshalJSON()
	if err != nil {
		logging.GetLogger().Fatalf("cannot marshal signature: %s", err)
	}

	logging.GetLogger().Infof("signature marshaled: %s", sgMrsh)

	err = os.WriteFile(signatureFileName, sgMrsh, 0644)
	if err != nil {
		logging.GetLogger().Fatalf("cannot write signature file: %s", err)
	}

	logging.GetLogger().Infof("signature file created: %s", signatureFileName)
}
