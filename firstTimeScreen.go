package main

import (
	"time"

	"github.com/ajstarks/openvg"
	touchscreen "github.com/electrocatstudios/FTXXXX_Touchscreen_Driver"
)

func RunFirstTimeScreen(s ScreenDetails, t *touchscreen.TouchScreen) bool {
	for {

		openvg.Start(s.Width, s.Height) // Start the picture
		openvg.BackgroundColor("black")

		count, _ := t.GetTouchesCount()
		if count > 0 {
			break
		}

		// Default fill color
		openvg.FillColor("rgb(255,255,255)") // White text

		line1 := DisplayLine{Type: "text", Value: "Welcome", Color: ""}
		line2 := DisplayLine{Type: "text", Value: "Touch Screen to Perform Setup", Color: ""}

		DrawLine(s, line1, 400, nil)
		DrawLine(s, line2, 160, nil)

		openvg.End()

		time.Sleep(5 * time.Millisecond)
	}

	// Start the command to get the screens - TODO: Check the file exists
	RunCommand("/bin/bash /home/pi/menecs/get_screens.sh", false)

	for {
		openvg.Start(s.Width, s.Height) // Start the picture
		openvg.BackgroundColor("black")

		// Default fill color
		openvg.FillColor("rgb(255,255,255)") // White text

		line1 := DisplayLine{Type: "text", Value: "Downloading", Color: ""}
		line2 := DisplayLine{Type: "text", Value: "App will restart when complete", Color: ""}

		DrawLine(s, line1, 400, nil)
		DrawLine(s, line2, 160, nil)

		openvg.End()

		time.Sleep(5 * time.Millisecond)
	}
}
