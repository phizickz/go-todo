package card

import (
	"fmt"
	"os"
)

// func GetCard(id int) {

// }

func writeToFile(file string, data string) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		os.Create(file)
	}

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	_, err = f.WriteString(data)
	if err != nil {
		panic(err)
	}
}

func NewCard(title string, body string) {
	FilePath := "cards.txt"

	writeToFile(FilePath, fmt.Sprintf("title:%s,body:%s;", title, body))

}

// func UpdateCard(id int) {}

// func DeleteCard(id int) {}
