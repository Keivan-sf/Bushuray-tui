package utils

import (
	"fmt"
	"log"
)

func LogEverywhere(msg string) {
	fmt.Println(msg)
	log.Println(msg)
}
