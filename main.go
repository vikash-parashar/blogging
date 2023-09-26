// main.go
package main

import (
	"blogging/config"
	"blogging/routes"
	"log"
)

func main() {
	config.InitDB()
	r := routes.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalln("error : failed to start application")
		log.Println(err)
	}
}
