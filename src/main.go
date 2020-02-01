package main

import (
	"fmt"
	"image"
	"log"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

//WIDTH is the width of each image you want
const WIDTH = 224

//HEIGHT is the width of each image you want
const HEIGHT = 224

func main() {
	// open directory
	imagePath := "./files"
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatalf("Failed opening directory: %s", err)
	}
	defer file.Close()

	flag := false // flag = true when detected wrong format or dimension

	// scan images one by one
	imageList, _ := file.Readdirnames(0) // 0 to read all files and folders
	for _, name := range imageList {
		// open image
		currentImage, imgErr := os.Open("./files/" + name)
		defer file.Close()
		if imgErr != nil {
			log.Fatalf("Failed opening image: %s", imgErr)
		}

		// decode the image
		decodedImage, format, dImgErr := image.DecodeConfig(currentImage)
		if dImgErr != nil {
			log.Fatalf("Failed opening directory: %s", dImgErr)
		}
		// fmt.Println("Width:", decodedImage.Width, "Height:", decodedImage.Height)

		// check format
		if format != "jpeg" {
			fmt.Println("Wrong format: " + name)
			flag = true
		}

		// check dimension
		if decodedImage.Width != WIDTH || decodedImage.Height != HEIGHT {
			fmt.Println("Wrong dimension: " + name)
			flag = true
		}
	}

	// print this message when all dimensions and foramts are correct
	if flag == false {
		fmt.Println("Dimensions and Formats of all images are correct!")
	}
}
