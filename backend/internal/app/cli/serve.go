package cli

import (
	"context"
	"github.com/Dubrovsky18/hachaton2023-gnomes/internal/app"
	"github.com/Dubrovsky18/hachaton2023-gnomes/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

// NewServeCmd starts new application instance
func NewServeCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "serve",
		Aliases: []string{"s"},
		Short:   "Start server",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("starting application")

			sigchan := make(chan os.Signal, 1)
			signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			application, err := app.InitializeApplication()
			if err != nil {
				logger.Fatal("can not initialize application", "error", err)
			}

			cliMode := false
			application.Start(ctx, cliMode)

			logger.Info("started")
			<-sigchan

			logger.Info("stop application", "error", application.Stop())

			time.Sleep(time.Second * cliCmdExecFinishDelaySeconds)
			logger.Info("finished")
		},
	}
}
