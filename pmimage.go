package main

import (
	"errors"
	"fmt"
	"image"
	"os"

	"github.com/nfnt/resize"
)

type PMImageCache struct {
	Images map[string]image.Image
}

func (imgCache *PMImageCache) LoadImage(fn string, width int, height int) error {
	filename := fmt.Sprintf("images/%s", fn)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println(err)
		errString := fmt.Sprintf("File %s does not exist in the images folder", fn)
		err := errors.New(errString)
		return err
	}

	imgfile, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer imgfile.Close()

	tmpImage, _, err := image.Decode(imgfile)
	if err != nil {
		return err
	}

	tmpImage = resize.Resize(uint(width), uint(height), tmpImage, resize.Lanczos3)

	if imgCache.Images == nil {
		imgCache.Images = make(map[string]image.Image)
	}

	imgCache.Images[fn] = tmpImage
	return nil
}

func (imgCache *PMImageCache) GetImage(imageName string, width int, height int) (image.Image, error) {
	if _, ok := imgCache.Images[imageName]; ok {
		return imgCache.Images[imageName], nil
	}

	err := imgCache.LoadImage(imageName, width, height)
	if err != nil {
		return nil, err
	}
	return imgCache.Images[imageName], nil
}
