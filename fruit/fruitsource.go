package fruit

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
)

type FruitSource interface {
	GetNextFruit() (string, string, io.Reader)
}

func NewFruitSource(fruitFileNames []string) FruitSource {
	return &fruitSource{
		callCount: 0,
		fruitNames: []string{
			"Apple",
			"Pineapple",
			"Banana",
			"Eggplant",
		},
		fruitDescriptions: []string{
			"omg, an apple",
			"dude look, a pineapple",
			"these bananas are bananas",
			"\"WTF man that's a fruit?\" Yeah, it is look it up",
		},
		fruitFileNames: fruitFileNames,
	}
}

type fruitSource struct {
	callCount         int
	fruitNames        []string
	fruitDescriptions []string
	fruitFileNames    []string
}

func (s *fruitSource) GetNextFruit() (string, string, io.Reader) {
	index := s.callCount % 4
	s.callCount += 1
	fruitRelativeFilePath := s.fruitFileNames[index]
	fruitFilePath, filePathErr := filepath.Abs(fruitRelativeFilePath)
	if filePathErr != nil {
		log.Fatal(filePathErr)
	}

	fruitFileBytes, err := ioutil.ReadFile(fruitFilePath)
	if err != nil {
		log.Fatal(err)
	}

	fruitImageDataReader := bytes.NewBuffer(fruitFileBytes)

	fruitName := s.fruitNames[index]
	fruitDescription := s.fruitDescriptions[index]
	return fruitName, fruitDescription, fruitImageDataReader
}
