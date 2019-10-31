package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

var log *logrus.Logger

var logger *zap.Logger
var sugar *zap.SugaredLogger

func init() {
	// Logrus
	log = logrus.New()
	log.SetLevel(logrus.WarnLevel)
	log.SetOutput(ioutil.Discard)

	// Zap
	var cfg zap.Config
	json.Unmarshal([]byte(`{
	  "level": "warn",
	  "encoding": "json",
	  "outputPaths": [],
	  "errorOutputPaths": ["stderr"]
	}`), &cfg)
	logger, _ = cfg.Build()
	sugar = logger.Sugar()
}

func main() {
	logger.Sync()
}

func ZapInfo(a string) {
	logger.Info("failed to fetch URL")
}

func ZapInfoWithFields(a string) {
	logger.Info("failed to fetch URL", zap.String("url", a), zap.Int("attempt", 3))
}

func ZapSugarInfo(a string) {
	sugar.Info("failed to fetch URL")
}

func ZapSugarInfoWithFields(a string) {
	sugar.Info("failed to fetch URL", "url", a, "attempt", 3)
}

func LogrusInfo(a string) {
	log.Info("failed to fetch URL")
}

func LogrusInfoWithFields(a string) {
	log.WithFields(logrus.Fields{"url": a, "attempt": 10}).Info("failed to fetch URL")
}

func ZapWarn(a string) {
	logger.Warn("failed to fetch URL")
}

func ZapWarnWithFields(a string) {
	logger.Warn("failed to fetch URL", zap.String("url", a), zap.Int("attempt", 3))
}

func ZapSugarWarn(a string) {
	sugar.Warn("failed to fetch URL")
}

func ZapSugarWarnWithFields(a string) {
	sugar.Warn("failed to fetch URL", "url", a, "attempt", 3)
}

func LogrusWarn(a string) {
	log.Warn("failed to fetch URL")
}

func LogrusWarnWithFields(a string) {
	log.WithFields(logrus.Fields{"url": a, "attempt": 10}).Warn("failed to fetch URL")
}
