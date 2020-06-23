package main

import (
	"fmt"

	"github.com/ajstarks/openvg"
)

var cur_screen = "main"

func DrawTextLine(s ScreenDetails, dl DisplayLine, offset openvg.VGfloat) {
	if dl.Type == "null" {
		// Nothing to do here - just leave blank
		return
	} else if dl.Type == "text" {
		openvg.TextMid(s.W2, offset, dl.Value, "sans", 30)
	} else if dl.Type == "data" {
		dataString := GetDataString(dl.Value)
		openvg.TextMid(s.W2, offset, dataString, "sans", 30)
	}
}

/*return new screen based on touches if appropriate*/
func DrawScreen(name string, input Screen, s ScreenDetails) string {
	openvg.Start(s.Width, s.Height)      // Start the picture
	openvg.BackgroundColor("black")      // Black background
	openvg.FillColor("rgb(255,255,255)") // White text

	DrawTextLine(s, input.Line1, 400)
	DrawTextLine(s, input.Line2, 320)
	DrawTextLine(s, input.Line3, 240)
	DrawTextLine(s, input.Line4, 160)
	DrawTextLine(s, input.Line5, 80)

	openvg.End() // End the picture
	return name  // TODO: Respond to touches appropriately
}

func main() {
	fmt.Println("Starting Screen App")

	var screenDetails ScreenDetails
	screenDetails.Width, screenDetails.Height = openvg.Init()
	screenDetails.W2 = openvg.VGfloat(screenDetails.Width / 2) // this is to center bits

	for {
		screen, err := GetScreenByName(cur_screen)
		if err != nil {
			panic(err)
		}
		DrawScreen(cur_screen, screen, screenDetails)
	}
}
