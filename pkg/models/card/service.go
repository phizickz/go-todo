package card

import (
	"os"
)

func GetCard(id int) {

}

func NewCard(title string, body string) {
	FilePath := "cards.txt"
	fc, err := os.Open(FilePath)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fc.Close(); err != nil {
			panic(err)
		}
	}()

}

// func UpdateCard(id int) {}

// func DeleteCard(id int) {}
