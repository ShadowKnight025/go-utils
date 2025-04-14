package utils

import (
	"os"
	"fmt"
	"github.com/redis/go-redis/v9"
	"utils/file"
)


func _init_db()(username string, password string){

	username, password, err := file.read("~/.creds/rdb-general-creds.go")
	if err != nil{
		fmt.Println("Error reading DB creds file: ", err)
	}
	return username, password
}

func new_client(conn_str string) {

	// redis://<user>:<pass>@localhost:6379/<db>
	opt, err := redis.ParseURL(conn_str)
	if err != nil {
		fmt.Println("Error creating DB Client: ", err)
	}

	client := redis.NewClient(opt)
	return client
}

