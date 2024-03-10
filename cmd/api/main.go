package main

import (
	"github.com/razorpay/retail-store/database"
	"github.com/razorpay/retail-store/routes"
)

const (
	ServerPort = ":1111"
)

func main() {
	database.Init()

	r := routes.SetupRouter()

	err := r.Run(ServerPort)
	if err != nil {
		panic(err)
	}
}
