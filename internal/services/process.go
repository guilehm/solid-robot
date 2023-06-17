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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
		service.channelRawMsg <- scanner.Text()
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	logger.Info().
		Int("line_count", lineCount).
		Str("filename", filename).
		Msg("finished processing file")

	return nil
}
