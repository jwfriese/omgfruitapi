package fruit

import (
	"bufio"
	"encoding/base64"
	"log"
	"os"
	"path/filepath"
)

type FruitSource interface {
	GetNextFruit() (string, string, string)
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

func (s *fruitSource) GetNextFruit() (string, string, string) {
	index := s.callCount % 4
	s.callCount += 1
	fruitRelativeFilePath := s.fruitFileNames[index]
	fruitFilePath, filePathErr := filepath.Abs(fruitRelativeFilePath)
	if filePathErr != nil {
		log.Fatal(filePathErr)
	}

	fruitFile, err := os.Open(fruitFilePath)
	if err != nil {
		log.Fatal(err)
	}

	fruitFileInfo, _ := fruitFile.Stat()
	var size int64 = fruitFileInfo.Size()
	fruitFileBase64Bytes := make([]byte, size)
	fruitFileReader := bufio.NewReader(fruitFile)
	fruitFileReader.Read(fruitFileBase64Bytes)

	fruitFileBase64EncodedString := base64.StdEncoding.EncodeToString(fruitFileBase64Bytes)

	fruitName := s.fruitNames[index]
	fruitDescription := s.fruitDescriptions[index]
	return fruitName, fruitDescription, fruitFileBase64EncodedString
}
