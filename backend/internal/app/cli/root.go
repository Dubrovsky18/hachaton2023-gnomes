package cli

import (
	"github.com/Dubrovsky18/hachaton2023-gnomes/pkg/logger"
	"github.com/spf13/cobra"
)

// ExecuteRootCmd prepares all CLI commands
func ExecuteRootCmd() {
	c := cobra.Command{}

	c.AddCommand(NewServeCmd())
	c.AddCommand(NewMigrateCmd())
	c.AddCommand(NewSeedCmd())

	if err := c.Execute(); err != nil {
		logger.Fatal(err.Error())
	}
}
