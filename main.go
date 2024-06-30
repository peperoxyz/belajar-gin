package main

import "belajar-gin/routers"

func main() {
	var PORT = ":9090"

	// start server
	routers.StartServer().Run(PORT)
}