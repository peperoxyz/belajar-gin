package main

import "belajar-gin/routers"

func main() {
	var PORT = ":9090"

	routers.StartServer().Run(PORT)
}