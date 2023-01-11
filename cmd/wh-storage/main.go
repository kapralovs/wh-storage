package main

import (
	"log"

	"github.com/kapralovs/wh-storage/internal/app"
)

func main() {
	a := app.New()
	log.Fatal(a.Run())
}
