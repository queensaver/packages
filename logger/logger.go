package logger

// This is a logging module that enforces structured logging. 

import (
    "github.com/sirupsen/logrus"
)

var (
    log *logrus.Logger
)

func init() {
    log = logrus.New()
    log.Formatter = &logrus.JSONFormatter{}
    //log.Formatter = &logrus.TextFormatter{}
    log.SetLevel(logrus.DebugLevel)
    // TODO: Remove. This is a huge performance penalty.
    // log.SetReportCaller(true)
}

// TODO: Context logging: https://notes.burke.libbey.me/context-and-logging/

func Debug(ip string, msg string) {
  log.WithFields(logrus.Fields{
    "ip": ip,
  }).Debug(msg)
}

func Info(ip string, msg string) {
  log.WithFields(logrus.Fields{
    "ip": ip,
  }).Info(msg)
}

func Error(ip string, err error) {
  log.WithFields(logrus.Fields{
    "ip": ip,
  }).Error(err)
}

var (

    // ConfigError ...
    ConfigError = "%v type=config.error"

    // HTTPError ...
    HTTPError = "%v type=http.error"

    // HTTPWarn ...
    HTTPWarn = "%v type=http.warn"

    // HTTPInfo ...
    HTTPInfo = "%v type=http.info"
)
