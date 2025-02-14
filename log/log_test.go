package log

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func Test(t *testing.T) {
	logrus.Info("hello world")

	InitDefaultLogger("[OKOK]")
	logrus.Info("hello world")
	logrus.Warn("hello world")
	logrus.Error("hello world")
	logrus.Println("hello world")
}
