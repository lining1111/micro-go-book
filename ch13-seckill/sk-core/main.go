package main

import (
	"micro-go-book/ch13-seckill/sk-core/setup"
)

func main() {

	setup.InitZk()
	setup.InitRedis()
	setup.RunService()

}
