package main

import (
	"github.com/myapp/noname/api/presentation/router"
)

func main() {
	r := router.Init()
	r.Swagger()
	r.Routes()
	r.Run()
}
