package ulogging

import (
	"log"
	"os"
)

func Basic_Logging(filename string){
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil{
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("Log initialized")
}
