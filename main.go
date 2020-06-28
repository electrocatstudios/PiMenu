package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/ajstarks/openvg"
	"github.com/electrocatstudios/PiMenu/screenservice"
	"google.golang.org/grpc"
)

var cur_screen = "main"
var imageCache PMImageCache
var interruptScreen InterruptScreen

func DrawLine(s ScreenDetails, dl DisplayLine, offset openvg.VGfloat) {
	if dl.Type == "null" || dl.Type == "" {
		// Nothing to do here - just leave blank
		return
	} else if dl.Type == "text" {
		openvg.TextMid(s.W2, offset, dl.Value, "sans", 30)
	} else if dl.Type == "data" {
		dataString := GetDataString(dl.Value)
		openvg.TextMid(s.W2, offset, dataString, "sans", 30)
	} else if dl.Type == "image" {
		img, err := GetImageFromString(dl.Value)

		if err != nil {
			fmt.Println(err)
			return
		}

		force_height := s.Height
		if img.Height != 0 {
			force_height = img.Height
		}

		image_file, err := imageCache.GetImage(img.Filename, 0, force_height)

		if err != nil {
			fmt.Println(err)
			return
		}
		left := openvg.VGfloat(0)
		top := openvg.VGfloat(img.Y)

		if img.X != 0 {
			// We aren't centering this one - use given x val
			left = openvg.VGfloat(img.X)
		} else {
			// Center the image
			image_width := openvg.VGfloat(image_file.Bounds().Max.X)
			left = openvg.VGfloat(s.W2 - (image_width / 2))
		}

		openvg.Img(left, top, image_file)
	}
}

func HandleTouches(t *TouchScreen, input Screen, defaultScreen string) string {
	numTouches, err := t.GetTouchesCount()
	if err != nil {
		fmt.Println(err)
		return defaultScreen
	}
	if numTouches < 1 {
		if input.Timeout.Length == 0 {
			// We don't have a time out so just return
			return defaultScreen
		}

		// See if we have timed out since last touch
		curTime := time.Now()
		diff := curTime.Sub(t.LastScreenChange).Seconds()

		remain := input.Timeout.Length - int(diff)

		if int(remain) < 0 {
			t.LastScreenChange = time.Now()
			// We have timed out so return to previous screen
			return input.Timeout.ReturnScreen
		}

		if remain < input.Timeout.ShowCountDown {
			seconds_remain := strconv.FormatInt(int64(remain), 10)
			seconds_remain += "s"

			openvg.Text(10, 420, seconds_remain, "sans", 30)

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
	DrawLine(s, input.Line1, 400)
	DrawLine(s, input.Line2, 320)
	DrawLine(s, input.Line3, 240)
	DrawLine(s, input.Line4, 160)
	DrawLine(s, input.Line5, 80)

	var ret string
	ret = HandleTouches(t, input, name)

	openvg.End()

	return ret
}

func monitorService(s *screenservice.Server) {
	// Listen for incoming screens on the buffer
	for {
		if s.NumScreens() != 0 {
			fmt.Println("Found screen")
			screen := s.GetScreen()
			var newScreen Screen
			newScreen.Line1 = DisplayLine{Type: "text", Value: screen.Line1}
			newScreen.Line2 = DisplayLine{Type: "text", Value: screen.Line2}
			newScreen.Line3 = DisplayLine{Type: "text", Value: screen.Line3}
			newScreen.Line4 = DisplayLine{Type: "text", Value: screen.Line4}
			newScreen.Line5 = DisplayLine{Type: "text", Value: screen.Line5}
			newScreen.Timeout = Timeout{Length: int(screen.Length), ShowCountDown: int(screen.ShowCountdown)}
			interruptScreen.Lock.Lock()
			if len(interruptScreen.Screens) == 0 {
				interruptScreen.LastShown = time.Now()
			}
			interruptScreen.Screens = append(interruptScreen.Screens, newScreen)
			interruptScreen.Lock.Unlock()
		}
		time.Sleep(1 * time.Second)
	}
}

func runInterruptServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := screenservice.Server{}
	grpcServer := grpc.NewServer()

	go monitorService(&s)

	screenservice.RegisterScreenServerServer(grpcServer, &s)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func main() {
	fmt.Println("Starting Screen App")

	var screenDetails ScreenDetails
	screenDetails.Width, screenDetails.Height = openvg.Init()
	screenDetails.W2 = openvg.VGfloat(screenDetails.Width / 2) // this is to center lines of text

	t := TouchScreen{nil, false, time.Now()}
	t.Init()

	go runInterruptServer()

	for {
		bFoundInterrupt := false
		interruptScreen.Lock.Lock()

		if len(interruptScreen.Screens) > 0 {
			bFoundInterrupt = true

			DrawScreen(&t, "interrupt", interruptScreen.Screens[0], screenDetails)
			curTime := time.Now()

			diff := curTime.Sub(interruptScreen.LastShown).Seconds()

			if int(diff) > interruptScreen.Screens[0].Timeout.Length {
				interruptScreen.Screens = interruptScreen.Screens[1:]
				interruptScreen.LastShown = time.Now()
			}
		}

		interruptScreen.Lock.Unlock()
		if bFoundInterrupt {
			continue
		}

		screen, err := GetScreenByName(cur_screen)
		if err != nil {
			panic(err)
		}
		cur_screen = DrawScreen(&t, cur_screen, screen, screenDetails)
	}
}
