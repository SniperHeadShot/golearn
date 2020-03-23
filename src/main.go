package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"name": "zy",
		"age":  26,
	}).Info("something occur!")
}
