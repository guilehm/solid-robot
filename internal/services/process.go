package services

import (
	"bufio"
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

func getFilePath(filename string) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/internal/tmp/%s", path, filename), nil
}

func (service *ServiceGroup) Process(logger *zerolog.Logger, filename string) error {
	logger.Info().Msg("processing " + filename)

	path, err := getFilePath(filename)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}

	lineCount := 0

	bulkLines := make([]string, 0, service.bulkAmount)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
		line := scanner.Text()

		bulkLines = append(bulkLines, line)

		if lineCount%service.bulkAmount == 0 {
			logger.Info().
				Int("line_count", lineCount).
				Msg("publishing message raw data message")

			service.rawMsgChannel <- bulkLines
			bulkLines = make([]string, 0, service.bulkAmount)
		}
	}

	// send remaining lines
	if lineCount%service.bulkAmount != 0 {
		logger.Info().
			Int("line_count", lineCount).
			Msg("publishing message raw data remainders message")
		service.rawMsgChannel <- bulkLines
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	return nil
}
