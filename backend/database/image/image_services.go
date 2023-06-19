package image

import (
	"gorm.io/gorm"
	"strconv"
)

type Image struct {
	gorm.Model
	Name        string
	Type        string
	Size        int
	BinaryValue string
}

var image = &Image{}

func BinaryToDecimal(binary string) (int, error) {
	decimal, err := strconv.ParseInt(binary, 2, 0)
	if err != nil {
		return 0, err
	}
	return int(decimal), nil
}
