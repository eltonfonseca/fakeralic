package main

import (
	"fakeralic/pkg/config"
	"os"

	joonix "github.com/joonix/log"
	"github.com/sirupsen/logrus"
)

var c = config.New()

func init() {
	logrus.SetOutput(os.Stdout)
	if c.Environment != "development" {
		logrus.SetFormatter(joonix.NewFormatter())
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func main() {
	logrus.WithField("Environment", c.Environment).Info("Started Function")
	logrus.WithField("EnvExample", c.EnvExemple).Info("Example Environment Variable")
}
