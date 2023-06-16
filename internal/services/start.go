package services

import (
	"bufio"
	"context"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"os"
	"strings"
	"time"
)

func start(service *ServiceGroup, ctx context.Context, logger *zerolog.Logger) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		select {
		default:
			logger.Info().Msg("ENTER THE FILENAME TO PROCESS (or 'exit' to quit):")
			input, _ := reader.ReadString('\n')
			input = input[:len(input)-1]

			if strings.ToLower(input) == "exit" {
				return true
			}
			err := service.Process(logger, ctx, input)
			if err != nil {
				logger.Error().Err(err).Msg("error calling service.Process")
				return false
			}
		}
	}

}

func StartListener(lc fx.Lifecycle, s fx.Shutdowner, logger *zerolog.Logger, service *ServiceGroup) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				logger.Info().Msg("starting service")

				go func() {
					time.Sleep(5 * time.Millisecond)
					if start(service, ctx, logger) {
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
