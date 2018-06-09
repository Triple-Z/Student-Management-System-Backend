package app

import "log"

func checkError(err error) bool {
	if err != nil {
		log.Fatal(err.Error())
		return true
	}
	return false
}
