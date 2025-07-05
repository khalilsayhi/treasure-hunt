package utils

import (
	"bufio"
	"fmt"
	"github.com/khalilsayhi/treasure-hunt/domain/constants"
	"github.com/pkg/errors"
	"os"
)

func GetBoardInitialConfig(mapPath string) (width, height int, landscapeConfLines []string, err error) {
	file, err := os.Open(mapPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	foundMapConfig := false

	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == 'C' && !foundMapConfig {
			_, err = fmt.Sscanf(line, constants.MapFormat, &width, &height)
			if err != nil {
				return
			}
			foundMapConfig = true
		} else {
			landscapeConfLines = append(landscapeConfLines, line)
		}
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	return
}

func WriteLineListToFile(filePath string, lines []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return errors.Wrapf(constants.ErrorFileCreation, "filepath: %s, error: %+v", filePath, err)
	}
	defer file.Close()

	for _, line := range lines {
		if _, err = file.WriteString(line + "\n"); err != nil {
			return errors.Wrapf(constants.ErrorFileWrite, "filepath: %s, error: %+v", filePath, err)
		}
	}

	return nil
}
