package migration

import (
	"flag"
	"fmt"
	"github.com/pressly/goose/v3"
	"log"
	"os"
	driver "template/adapter"
	"template/config"
)

var (
	flags         = flag.NewFlagSet("adapter:migrate", flag.ExitOnError)
	usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    down                 Roll back the version by 1
    reset                Roll back all migrations
`
	dir = flags.String("dir destination", "adapter/migration", "directory with migration destination")
)

// RunMigration running auto migration
func RunMigration(cfg config.Config) {
	// assign var to flags
	flags.Usage = usage
	flags.Parse(os.Args[2:])

	args := flags.Args()
	if len(args) == 0 {
		flags.Usage()
		return
	}

	command := args[0]
	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.SetDialect("mysql"); err != nil {
		log.Fatalf("Failed to set dialect: %v", err)
	}
	dbSrc, err := driver.NewDatabase(cfg.DB)
	if err != nil {
		log.Fatalf(err.Error())
	}

	//close connection
	defer func() {
		if err := dbSrc.Close(); err != nil {
			log.Fatalf("adapter migrate: failed to close DB: %v\n", err)
		}
	}()

	// running migration in destination folder
	if err := goose.Run(command, dbSrc, *dir, arguments...); err != nil {
		log.Fatalf("adapter migrate run: %v", err)
	}

}

// usage print list of command
func usage() {
	fmt.Println(usageCommands)
}
