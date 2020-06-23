package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"io/ioutil"
	"os"

	"github.com/ajstarks/openvg"
)

type Screen struct {
	Name    string      `json:"name"`
	Line1   DisplayLine `json:"line1"`
	Line2   DisplayLine `json:"line2"`
	Line3   DisplayLine `json:"line3"`
	Line4   DisplayLine `json:"line4"`
	Line5   DisplayLine `json:"line5"`
	Timeout int         `json:"timeout"`
}

type ScreenDetails struct {
	Width  int
	Height int
	W2     openvg.VGfloat
	Images map[string]image.Image
}

type DisplayLine struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

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
