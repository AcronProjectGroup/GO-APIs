package main

import (
	"WebServerAcron/server"

)

func main()  {
	
	server := server.New()

	server.Server(8080)


}