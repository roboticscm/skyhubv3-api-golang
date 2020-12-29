package slog

import (
	"fmt"
	"log"
)

func Compaq(obj ...interface{}) {
	if len(obj) == 2 {
		log.Printf("%v: %v\n", obj[0], obj[1])
	} else if len(obj) == 1 {
		log.Printf("%v\n", obj[0])
	} else {
		log.Printf("%v\n", obj)
	}
}

func Detail(obj ...interface{}) {
	if len(obj) == 2 {
		log.Printf("%v: %#v\n", obj[0], obj[1])
	} else if len(obj) == 1 {
		log.Printf("%#v\n", obj[0])
	} else {
		log.Printf("%#v\n", obj)
	}
}

func Fatal(obj ...interface{}) {
	if len(obj) > 0 && obj[0] != nil {
		if len(obj) == 2 {
			log.Fatal(fmt.Sprintf("%v: %v\n", obj[0], obj[1]))
		} else {
			log.Fatal(fmt.Sprintf("%v\n", obj))
		}
	}

}
