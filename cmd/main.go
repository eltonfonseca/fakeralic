package main

import (
	"fakeralic/pkg/config"
	"fakeralic/pkg/engine"
	"fakeralic/pkg/engine/models"
	"fmt"
	"os"
	"time"

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
	for {
		menu()

		switch getOption() {
		case 1:
			startMonitoring()
		case 2:
			showLogs()
		case 0:
			fmt.Println("Exiting ...")
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
			os.Exit(-1)
		}
	}
}

func menu() {
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit")
}

func getOption() int {
	var option int

	fmt.Scan(&option)

	return option
}

func startMonitoring() {
	logrus.Info("Monitoring ...")

	m := &models.Message{}

	for i := 0; i < c.MonitoringQuantity; i++ {
		if err := engine.Execute(m); err != nil {
			logrus.WithError(err).Error("error on execute engine")
		}
		time.Sleep(time.Duration(c.Delay) * time.Second)
		clear(m)
	}
}

func showLogs() {
	logrus.Info("Show Logs")
}

func clear(m *models.Message) {
	m.Hosts = []models.Host{}
}
