package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"template/cmd/migration"
	"template/config"
	"template/constant"
)

func Start() {
	cfg := config.ReadConfig()
	// root command
	root := &cobra.Command{}

	// command allowed
	cmds := []*cobra.Command{
		{
			Use:   "adapter:migrate",
			Short: "database migration",
			Run: func(cmd *cobra.Command, args []string) {
				migration.RunMigration(cfg)
			},
		},
		{
			Use:   "api",
			Short: "run api server",
			Run: func(cmd *cobra.Command, args []string) {
				config.InitService(cfg, constant.ServerRest)
			},
		},
		{
			Use:   "grpc",
			Short: "run grpc server",
			Run: func(cmd *cobra.Command, args []string) {
				config.InitService(cfg, constant.ServerGRPC)
			},
		},
	}
	root.AddCommand(cmds...)
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
