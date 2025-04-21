package utils

import (
	"log"
	"os"
)

func basic_logging(filename string){
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.OWRONLY, 0666)
	if err != nil{
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("Log initialized")
}
