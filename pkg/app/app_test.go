package app

import (
	"testing"
	log "github.com/sirupsen/logrus"
)

const (
	packName =			"app"
)

func TestApp(t *testing.T) {
	funcName := "TestApp"
	log.Warnf("TEST %s::%s start\n", packName, funcName)

	New().Run()
}
