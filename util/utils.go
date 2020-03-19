package util

import "log"

func Panic(err error) {
	if err != nil {
		Panic(err)
	}
}

func FailedOnError(err error, msg string)  {
	if err != nil {
		log.Printf("%s: %v", msg, err)
	}
}