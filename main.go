package main

import (
	"fmt"

	"github.com/klurpicolo/finalexam/routers"
)

func main() {
	fmt.Println("Server start")
	r := routers.GetRouter()
	r.Run(":2019")
}
