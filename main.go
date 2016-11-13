package main

import (
	"log"
	"net/http"

	"github.com/jwfriese/omgfruitapi/fruit"
)

func main() {
	fruitFileNames := []string{
		"fruit/images/apple.png",
		"fruit/images/pineapple.png",
		"fruit/images/bananas.png",
		"fruit/images/eggplant.png",
	}
	http.Handle("/fruit", fruit.GetFruitHandler(fruit.NewFruitSource(fruitFileNames)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
