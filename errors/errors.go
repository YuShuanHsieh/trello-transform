package errors

import (
	"errors"
	"fmt"
	"log"
)

func NewFromStr(msg string) error {
	return errors.New(msg)
}

func NewFromFormat(msg string, args ...interface{}) error {
	return fmt.Errorf(msg, args...)
}

func LogFromFormat(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}

func Log(msg string) {
	log.Println(msg)
}
