package fruit

import (
	"encoding/json"
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
	Base64Image string `json:"image"`
}

func (h *fruitHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fruitName, fruitDescription, fruitImageBase64DataString := h.fruitSource.GetNextFruit()
	fruit := &fruit{
		Name:        fruitName,
		Description: fruitDescription,
		Base64Image: fruitImageBase64DataString,
	}

	fruitBytes, jsonErr := json.Marshal(fruit)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(fruitBytes)
}
