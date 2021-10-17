package main

import "github.com/sirupsen/logrus"

func main() {
	logger := logrus.StandardLogger()

	logger.Infof("backend started")
}
