package utils

import (
	"encoding/csv"
	"os"

	"github.com/c0d3G33k/GitWave/pkg/common"
)

func ValidateWorkDirectory() {

	if common.WorkDirectory != "/tmp/" {
		if common.WorkDirectory[len(common.WorkDirectory)-1] != '/' {
			common.WorkDirectory += "/"
		}
	}
}

func CheckGitHubPAT() bool {
	return os.Getenv("GITHUB_PAT") != ""
}

func ReadCSVFile(filePath string) (data [][]string, err error) {
	file, errFile := os.Open(filePath)

	if errFile != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, errRead := reader.ReadAll()

	if errRead != nil {
		return nil, err
	}

	return records, nil
}
