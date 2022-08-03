package helper

import (
	"os"

	"github.com/incu6us/goimports-reviser/v3/pkg/module"
	"github.com/incu6us/goimports-reviser/v3/reviser"
)

func DetermineProjectName(projectName, filePath string) (string, error) {
	if filePath == reviser.StandardInput {
		var err error
		filePath, err = os.Getwd()
		if err != nil {
			return "", err
		}
	}

	return module.DetermineProjectName(projectName, filePath)
}