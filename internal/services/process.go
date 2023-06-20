package services

import (
	"bufio"
	"context"
	"fmt"
	"github.com/guilehm/solid-robot/internal/features/models"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"os"
	"time"
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
		Str("now", time.Now().Format("15:04:05.999999999")).
		Str("filename", filename).
		Msg("processing file")

	startTime := time.Now()

	path, err := getFilePath(filename)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}

	lineCount := 0

	const bufferSize = 50000

	// create batches
	batchClientRaw := &pgx.Batch{}
	batchClient := &pgx.Batch{}

	// create channels
	quit := make(chan bool)
	rawMsgChannel := make(chan string, bufferSize)
	clientRawChannel := make(chan models.ClientRaw, bufferSize)
	formatChannel := make(chan models.ClientRaw, bufferSize)
	clientChannel := make(chan models.Client, bufferSize)

	// start workers
	go service.parserWorker(rawMsgChannel, clientRawChannel)
	go service.insertClientRawWorker(ctx, batchClientRaw, clientRawChannel, formatChannel)

	go service.formatterWorker(ctx, formatChannel, clientChannel)
	go service.insertClientWorker(ctx, batchClient, clientChannel, quit)

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

	<-quit

	logger.Info().
		Int("line_count", lineCount).
		Str("filename", filename).
		Str("elapsed_time", time.Since(startTime).String()).
		Msg("finished processing file")

	return nil
}
