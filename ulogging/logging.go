package logging

import (
	"log"
	"os"
)

func basic_logging(filename string){
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil{
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("Log initialized")
}
