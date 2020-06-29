package main

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"os"

	"github.com/nfnt/resize"
)

type PMImageCache struct {
	Images map[string]image.Image
	Gifs   map[string]*PMGif
}

type PMGif struct {
	Frames        []image.Image
	PreviousFrame int
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

func (imgCache *PMImageCache) LoadGif(fn string, s ScreenDetails) error {
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

	gif_out, err := gif.DecodeAll(imgfile)
	if err != nil {
		return err
	}

	// Create the PMGif object in the map
	if imgCache.Gifs == nil {
		imgCache.Gifs = make(map[string]*PMGif)
	}
	var nextPMGif PMGif

	nextPMGif.PreviousFrame = 0

	for _, srcImage := range gif_out.Image {
		nextPMGif.Frames = append(nextPMGif.Frames, srcImage)
	}

	fmt.Println(len(nextPMGif.Frames))
	imgCache.Gifs[fn] = &nextPMGif

	// TODO: Scale the image
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

func (imgCache *PMImageCache) GetGifFrame(imageName string, s ScreenDetails) (image.Image, error) {
	if _, ok := imgCache.Gifs[imageName]; ok {
		return imgCache.Gifs[imageName].GetNextFrame()
	}

	// Not yet cached - load in
	err := imgCache.LoadGif(imageName, s)

	if err != nil {
		fmt.Println(err)
		var blank image.Image
		return blank, err
	} else {
		ret := imgCache.Gifs[imageName]
		return ret.GetNextFrame()
	}
}

func (gif *PMGif) GetNextFrame() (image.Image, error) {
	gif.PreviousFrame += 1
	if gif.PreviousFrame > len(gif.Frames)-1 {
		gif.PreviousFrame = 0
		return gif.Frames[0], nil
	}
	return gif.Frames[gif.PreviousFrame], nil
}
