package services

import (
	"bufio"
	"context"
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

func (service *ServiceGroup) Process(logger *zerolog.Logger, ctx context.Context, filename string) error {
	logger.Info().Msg("processing " + filename)

	path, err := getFilePath(filename)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	return nil
}
