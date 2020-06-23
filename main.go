package main

import (
	"fmt"
	"time"

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

func HandleTouches(t *TouchScreen, input Screen, defaultScreen string) string {
	numTouches, err := t.GetTouchesCount()
	if err != nil {
		fmt.Println(err)
		return defaultScreen
	}
	if numTouches < 1 {
		if input.Timeout == 0 {
			// We don't have a time out so just return
			return defaultScreen
		}

		curTime := time.Now()
		diff := curTime.Sub(t.LastScreenChange).Seconds()

		remain := input.Timeout - int(diff)
		if int(remain) < 0 {
			// We have timed out so return to previous screen
			return input.ReturnScreen
		}
		return defaultScreen
	}

	t.LastScreenChange = time.Now()
	touch, _ := t.GetTouches()

	for i := 0; i < len(input.Touches); i++ {
		hitBox := input.Touches[i]

		if touch.X > hitBox.X && touch.X < hitBox.X+hitBox.Width {
			if touch.Y > hitBox.Y && touch.Y < hitBox.Y+hitBox.Height {

				if hitBox.Command.Type == "menu" {
					return hitBox.Command.Value
				}

			}
		}

	}

	return defaultScreen
}

/*return new screen based on touches if appropriate*/
func DrawScreen(t *TouchScreen, name string, input Screen, s ScreenDetails) string {
	openvg.Start(s.Width, s.Height)      // Start the picture
	openvg.BackgroundColor("black")      // Black background
	openvg.FillColor("rgb(255,255,255)") // White text

	DrawTextLine(s, input.Line1, 400)
	DrawTextLine(s, input.Line2, 320)
	DrawTextLine(s, input.Line3, 240)
	DrawTextLine(s, input.Line4, 160)
	DrawTextLine(s, input.Line5, 80)

	openvg.End() // End the picture

	ret := HandleTouches(t, input, name)

	return ret
}

func main() {
	fmt.Println("Starting Screen App")

	var screenDetails ScreenDetails
	screenDetails.Width, screenDetails.Height = openvg.Init()
	screenDetails.W2 = openvg.VGfloat(screenDetails.Width / 2) // this is to center lines of text

	t := TouchScreen{nil, false, time.Now()}
	t.Init()

	for {
		screen, err := GetScreenByName(cur_screen)
		if err != nil {
			panic(err)
		}
		cur_screen = DrawScreen(&t, cur_screen, screen, screenDetails)
	}
}
