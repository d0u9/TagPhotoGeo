package app

import (
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
	"TagPhotoGeo/pkg/config"
)

type App struct {
	cmd			*cobra.Command
	cmdOpts		*config.CmdLineOptions
}

func New() *App {
	app := &App {
		cmdOpts:	&config.CmdLineOptions {},
	}

	app.cmd = &cobra.Command {
		Use:   filepath.Base(os.Args[0]),
		Short: "",
		Long: ``,
		Run: func(cmd *cobra.Command, args []string) {
			app.doit()
		},
	}

	app.cmdSetup()

	return app
}

func (app *App) Run() {
	app.cmd.Execute()
}

func (app *App) doit() {
	log.Warn("Start App")
	if err := config.UpdateFromCmdLineOptions(app.cmdOpts); err != nil {
		log.Error(err.Error())
		goto exit
	}

	return

exit:
	os.Exit(1)
}

func (app *App) cmdSetup() {
	flags := app.cmd.PersistentFlags()
	flags.StringVar(&app.cmdOpts.LogLevel, "log-level", "info", "log level")
}
