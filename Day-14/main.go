package main

import "Day-14/routers"

func main(){
	var PORT = ":8000"

	routers.StartServer().Run(PORT)
}