package main

import "Quiz3/routers"

func main(){
	routers.StartServer().Run(":8000")
}