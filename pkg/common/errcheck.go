package common

import "log"

func Errcheck(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
