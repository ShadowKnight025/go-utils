package utils

import (
	"os"
	"fmt"
	"context"
	"github.com/redis/go-redis/v9"

)

_CONN_STR := os.Getenv("redcons") or ""

func _init_db(){

	username, password, host,  err := file.read("~/.creds/rdb-general-creds.go")
	if err != nil{
		fmt.Println("Error reading DB creds file: ", err)
	}
	return username, password
}

func new_client(conn_str string) {

	// redis://<user>:<pass>@localhost:6379/<db>
	if _CONN_STR == nil{
		username, password, host = _init_db()
		_CONN_STR = fmt.Sprintf("redis://%s:%s@%s", username, password, host)
	}
	opt, err := redis.ParseURL(_CONN_STR)
	if err != nil {
		fmt.Println("Error creating DB Client: ", err)
		os.Exit(1)
	}

	client := redis.NewClient(opt)
	ctx    := context.Background()
	return client, ctx
}

