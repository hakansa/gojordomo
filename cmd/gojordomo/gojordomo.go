package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hakansa/entegrator/pkg/version"
	"github.com/sirupsen/logrus"
)

var (
	configFileFlag = flag.String("config.file", "config.yml", "Path to the configuration file.")
	logFileFlag    = flag.String("log.file", "license_controller.log", "Path to the log file.")
	versionFlag    = flag.Bool("version", false, "Show version information.")
	debugFlag      = flag.Bool("debug", false, "Show debug information.")
)

func init() {
	// Parse command-line flags
	flag.Parse()

	// Log settings
	if *debugFlag {
		logrus.SetReportCaller(true)
		logrus.SetLevel(logrus.TraceLevel)
	} else {
		logrus.SetReportCaller(false)
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logFile, err := os.OpenFile(*logFileFlag, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.WithError(err).Fatal("Could not open log file")
	}

	logrus.SetOutput(logFile)
}

func main() {
	// Show version information
	if *versionFlag {
		fmt.Fprintln(os.Stdout, version.Print("gojordomo"))
		os.Exit(0)
	}

	// Load configuration file
	_, err := ioutil.ReadFile(*configFileFlag)
	if err != nil {
		logrus.WithError(err).Fatal("Could not load configuration")
	}
}
