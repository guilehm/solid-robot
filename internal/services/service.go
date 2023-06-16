package services

import (
	"bufio"
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"os"
	"time"
)

func processFile(filename string) {
	fmt.Println("processing", filename)
}

type ServiceGroup struct{}

func start() bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		select {
		default:
			fmt.Print("enter the filename to process (or 'exit' to quit): \n")
			input, _ := reader.ReadString('\n')
			input = input[:len(input)-1]

			if input == "exit" {
				return true
			}
			processFile(input)
		}
	}
}

func StartListener(lc fx.Lifecycle, s fx.Shutdowner, logger *zerolog.Logger) {

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				logger.Info().Msg("starting service")

				go func() {
					time.Sleep(5 * time.Millisecond)
					if start() {
						err := s.Shutdown()
						if err != nil {
							logger.Error().Err(err).Msg("failed to shutdown")
						}
					}
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				logger.Info().Msg("stopping service")
				return nil
			},
		},
	)

}

func newServiceGroup() *ServiceGroup {
	return &ServiceGroup{}
}
