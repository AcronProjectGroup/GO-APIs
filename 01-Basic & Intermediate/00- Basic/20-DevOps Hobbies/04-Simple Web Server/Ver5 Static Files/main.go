package main

import (
	"WebServerAcron/server"

)

func main()  {
	
	server := server.New()

	server.Serve(8080)


}


/* How to curl?
curl -d '{"height": 185, "weight": 65}' -X POST localhost:8080/bmi


*/