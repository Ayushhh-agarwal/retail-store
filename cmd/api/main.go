package main

import (
	"github.com/razorpay/retail-store/cmd/boot"
	"github.com/razorpay/retail-store/database"
	"github.com/razorpay/retail-store/internal/mutex"
	"github.com/razorpay/retail-store/routes"
)

const (
	ServerPort = ":1111"
)

func main() {
	database.Init()

	mutex.InitializeRedis(mutex.GetRedisConfig())
	mutex.Mutex = mutex.InitializeMutex(mutex.GetRedisClient())
	r := routes.SetupRouter()

	boot.LoadAllModules()

	err := r.Run(ServerPort)
	if err != nil {
		panic(err)
	}
}
