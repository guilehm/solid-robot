package services

import (
	"bufio"
	"context"
	"fmt"
	"github.com/guilehm/solid-robot/internal/features/models"
	"github.com/jackc/pgx/v5"
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

func (service *ServiceGroup) Process(ctx context.Context, logger *zerolog.Logger, filename string) error {
	logger.Info().
		Str("filename", filename).
		Msg("processing file")

	path, err := getFilePath(filename)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}

	lineCount := 0
	batch := &pgx.Batch{}

	clientRawChannel := make(chan models.ClientRaw)
	go service.insertClientRaw(ctx, batch, clientRawChannel)

	rawMsgChannel := make(chan string)
	go service.parser(ctx, rawMsgChannel, clientRawChannel)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan() {
		lineCount++
		rawMsgChannel <- scanner.Text()
	}
	close(rawMsgChannel)

	if scanner.Err() != nil {
		return scanner.Err()
	}

	logger.Info().
		Int("line_count", lineCount).
		Str("filename", filename).
		Msg("finished processing file")

	return nil
}
