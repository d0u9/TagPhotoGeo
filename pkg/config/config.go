package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

var (
	configImpl	*Config
)

type CmdLineOptions struct {
	LogLevel		string
}

type Config struct {
	LogLevel		log.Level
}

func init() {
	configImpl = &Config {}
}

func UpdateFromCmdLineOptions(opts *CmdLineOptions) error {
	if level, err := log.ParseLevel(opts.LogLevel); err != nil {
		return fmt.Errorf("Command line option [LogLevel=%s] is invalid", opts.LogLevel)
	} else {
		configImpl.LogLevel = level
	}

	return nil
}
