package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ajstarks/openvg"
)

type Screen struct {
	Name    string         `json:"name"`
	Line1   DisplayLine    `json:"line1"`
	Line2   DisplayLine    `json:"line2"`
	Line3   DisplayLine    `json:"line3"`
	Line4   DisplayLine    `json:"line4"`
	Line5   DisplayLine    `json:"line5"`
	Timeout Timeout        `json:"timeout"`
	Touches []TouchDetails `json:"touches"`
}

type ScreenDetails struct {
	Width  int
	Height int
	W2     openvg.VGfloat
	Images map[string]image.Image
}

type TouchDetails struct {
	X       int            `json:"x"`
	Y       int            `json:"y"`
	Width   int            `json:"width"`
	Height  int            `json:"height"`
	Command CommandDetails `json:"command"`
}

type CommandDetails struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type DisplayLine struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Timeout struct {
	Length        int    `json:"length"`
	ReturnScreen  string `json:"returnScreen"`
	ShowCountDown int    `json:"showCountDown"`
}

type InterruptScreen struct {
	Screens       []Screen
	RecentFrame   *image.Image
	IncomingImage *image.Image
	Lock          sync.Mutex
	LastShown     time.Time
}

type PMImage struct {
	X        int
	Y        int
	Width    int
	Height   int
	Filename string
}

// Load in the screen data from json file
func GetScreenByName(name string) (Screen, error) {
	var ret Screen
	filename := fmt.Sprintf("screens/%s.json", name)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		ret.Line1.Type = "text"
		ret.Line1.Value = "File not found"
		return ret, errors.New("File " + name + ".json not found")
	}

	file, _ := ioutil.ReadFile(filename)

	err := json.Unmarshal([]byte(file), &ret)
	if err != nil {
		return ret, err
	}

	return ret, nil
}

func GetImageFromString(input string) (PMImage, error) {
	ret := PMImage{}
	pieceArray := strings.Split(input, ";")
	if len(pieceArray) < 5 || len(pieceArray) > 5 {
		errString := fmt.Sprintf("Incorrect number of items in image string expected 5 but got %d", len(pieceArray))
		err := errors.New(errString)
		return ret, err
	}
	x, err := strconv.Atoi(pieceArray[0])
	if err != nil {
		return ret, err
	}
	y, err := strconv.Atoi(pieceArray[1])
	if err != nil {
		return ret, err
	}
	width, err := strconv.Atoi(pieceArray[2])
	if err != nil {
		return ret, err
	}
	height, err := strconv.Atoi(pieceArray[3])
	if err != nil {
		return ret, err
	}
	ret.X = x
	ret.Y = y
	ret.Width = width
	ret.Height = height
	ret.Filename = pieceArray[4]
	return ret, nil
}
