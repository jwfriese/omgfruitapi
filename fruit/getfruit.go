package fruit

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetFruitHandler(fruitSource FruitSource) http.Handler {
	return &fruitHandler{
		fruitSource: fruitSource,
	}
}

type fruitHandler struct {
	fruitIndex  int
	fruitSource FruitSource
}

type fruit struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func (h *fruitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fruitName, fruitDescription, fruitImageReader := h.fruitSource.GetNextFruit()
	fruitImage, readErr := ioutil.ReadAll(fruitImageReader)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fruit := &fruit{
		Name:        fruitName,
		Description: fruitDescription,
		Image:       string(fruitImage),
	}

	fruitBytes, jsonErr := json.Marshal(fruit)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(fruitBytes)
}
