package model

import (
	"bytes"
	"encoding/base64"
	"golang_task/database"
	"image"
	"image/jpeg"
	"image/png"
	"strings"

	"github.com/disintegration/imaging"
	"gorm.io/gorm"
)

type StoredImage struct {
	gorm.Model
	NEGATIVE string `json:"negative_image"gorm:"not null;default:null"`
	POSITIVE string `json:"positive_image"gorm:"not null;default:null"`
}

func (input *InputImage) CreateEntity() (StoredImage, error) {
	var ret StoredImage

	ret.POSITIVE = input.Base64Image

	data_start_index := strings.Index(input.Base64Image, "base64,")
	image_data := input.Base64Image[data_start_index+7:]

	image_bytes, err := base64.StdEncoding.DecodeString(image_data)

	mimeType := input.Base64Image[0 : data_start_index+7]

	if err != nil {
		return ret, err
	}

	// create Image object from byte array
	im, _, err := image.Decode(bytes.NewReader(image_bytes))
	if err != nil {
		return ret, err
	}
	// invert colors of image
	inverted_image := imaging.Invert(im)

	// encode inverted image based on original format
	buf := new(bytes.Buffer)
	switch mimeType {
	case "data:image/jpeg;base64,":
		ret.NEGATIVE += "data:image/jpeg;base64,"
		err := jpeg.Encode(buf, inverted_image, nil)
		if err != nil {
			return ret, err
		}
	case "data:image/png;base64,":
		ret.NEGATIVE += "data:image/png;base64,"
		png.Encode(buf, inverted_image)
	}

	// png.Encode(buf, inverted_image)
	inverted_image_bytes := buf.Bytes()
	ret.NEGATIVE += base64.StdEncoding.EncodeToString(inverted_image_bytes)

	return ret, nil

}

func (entity *StoredImage) SaveEntity() (InputImage, error) {
	err := database.Database.Create(&entity).Error
	if err != nil {
		return InputImage{}, err
	}
	var ret InputImage
	ret.Base64Image = entity.NEGATIVE
	return ret, nil
}

func GetLastThree() ([]StoredImage, error) {
	var storedImages []StoredImage
	request := database.Database.Raw("SELECT * FROM stored_images ORDER BY updated_at DESC LIMIT 3;")

	request.Scan(&storedImages)
	err := request.Error

	return storedImages, err
}
