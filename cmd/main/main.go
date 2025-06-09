package main

import(
	"fmt"
	"github.com/ShadowKnight025/go-utils/uredis"
)

func main(){

	fmt.Print("init utils")
	client, ctx := uredis.New_Client("https://localhost:8099")
	fmt.Println(client)
	fmt.Println(ctx)
}

