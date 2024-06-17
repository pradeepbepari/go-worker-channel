package main

import (
	"fmt"
	"log"

	"github.com/pradeepbepari/jsonplaceholder/cmd/api"
)

func main() {
	server := api.NewServer(fmt.Sprintf(":%d", 8000))
	if err := server.Run(); err != nil {
		log.Panic(err)
		return
	}

}
