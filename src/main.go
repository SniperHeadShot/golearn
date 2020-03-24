package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"name": "zy",
		"age":  26,
	}).Info("something occur!")

	test()
}

func test(s ...string) {
	len := len(s)
	fmt.Println(len)
}
