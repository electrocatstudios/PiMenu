package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/ajstarks/openvg"
	touchscreen "github.com/electrocatstudios/FTXXXX_Touchscreen_Driver"
	"github.com/electrocatstudios/PiMenu/screenservice"
	"google.golang.org/grpc"
)

var cur_screen = "main"
var imageCache PMImageCache
var interruptScreen InterruptScreen

func GetImageWithOpacity(img image.Image, opacity uint8) image.Image {
	m := image.NewRGBA(img.Bounds())
	mask := image.NewUniform(color.Alpha{opacity})
	draw.DrawMask(m, m.Bounds(), img, image.Point{0, 0}, mask, image.Point{0, 0}, draw.Over)
	return m
}

func DrawLine(s ScreenDetails, dl DisplayLine, offset openvg.VGfloat, opacity *uint8) {
	if dl.Type == "null" || dl.Type == "" {
		// Nothing to do here - just leave blank
		return
	} else if dl.Type == "text" {
		if dl.Color != "" {
			openvg.FillColor(dl.Color)
		} else {
			openvg.FillColor("rgb(255,255,255)")
		}
		openvg.TextMid(s.W2, offset, dl.Value, "sans", 30)
	} else if dl.Type == "data" {
		if dl.Value == "IMAGESERVER" {
			interruptScreen.Lock.Lock()
			if interruptScreen.IncomingImage != nil {

				img := (*interruptScreen.IncomingImage)

				left := openvg.VGfloat(0)
				top := openvg.VGfloat(0)

				openvg.Img(left, top, img)
			} else {
				openvg.TextMid(s.W2, 200, "No Image available", "sans", 30)
			}
			interruptScreen.Lock.Unlock()
		} else {
			if dl.Color != "" {
				openvg.FillColor(dl.Color)
			} else {
				openvg.FillColor("rgb(255,255,255)")
			}
			dataString := GetDataString(dl.Value)
			openvg.TextMid(s.W2, offset, dataString, "sans", 30)
		}
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

		if opacity != nil && *opacity >= uint8(5) {
			img_op := GetImageWithOpacity(image_file, *opacity)
			openvg.Img(left, top, img_op)
		} else if opacity != nil && *opacity < uint8(5) {
			// Don't draw if less than 5 opacity - should be nearly invisible
		} else {
			openvg.Img(left, top, image_file)
		}

	} else if dl.Type == "gif" {
		img, err := GetImageFromString(dl.Value)
		if err != nil {
			fmt.Println(err)
			return
		}

		frame, err := imageCache.GetGifFrame(img.Filename, s)
		// Center the gif
		left := s.W2 - openvg.VGfloat((frame.Bounds().Max.X / 2))
		top := openvg.VGfloat((s.Height / 2) - (frame.Bounds().Max.Y / 2))

		if img.X != 0 {
			left = openvg.VGfloat(img.X)
		}
		if img.Y != 0 {
			top = openvg.VGfloat(img.Y)
		}

		openvg.Img(left, top, frame)
	}
}

func RunCommand(cmd string, wait bool) error {
	cmd_items := strings.Split(cmd, " ")
	cmd_exec := exec.Command(cmd_items[0], strings.Join(cmd_items[1:], " "))
	err := cmd_exec.Start()
	if err != nil {
		fmt.Printf("Failed to start command : %s\n", err)
		return err
	}
	if wait {
		err = cmd_exec.Wait()
		if err != nil {
			fmt.Printf("Failed to execute command: %s\n", err)
		}
	}

	return err
}

func HandleTouches(t *touchscreen.TouchScreen, input Screen, defaultScreen string) (string, *uint8) {
	numTouches, err := t.GetTouchesCount()
	if err != nil {
		fmt.Println(err)
		return defaultScreen, nil
	}

	if numTouches < 1 {
		if input.Timeout.Length == 0 {
			// We don't have a time out so just return
			return defaultScreen, nil
		}

		// See if we have timed out since last touch
		curTime := time.Now()
		diff := curTime.Sub(t.LastScreenChange).Seconds()

		remain := float64(input.Timeout.Length) - diff

		if int(remain) < 0 {
			t.LastScreenChange = time.Now()
			// We have timed out so return to previous screen
			ret_op := uint8(0)
			return input.Timeout.ReturnScreen, &ret_op
		}

		if remain < float64(input.Timeout.ShowCountDown) && input.Timeout.ShowCountDown != 0 {
			seconds_remain := strconv.FormatInt(int64(remain), 10)
			seconds_remain += "s"

			openvg.FillColor("rgb(255,255,255)")
			openvg.Text(10, 420, seconds_remain, "sans", 30)

		}

		// Apply opacity if appropriate
		if input.Background.FadeOut && remain < 0 {
			// Opacity should be zero while we load the new screen
			opacity := uint8(0)
			return defaultScreen, &opacity
		} else if remain < 1 && input.Background.FadeOut {
			opacity := uint8(255 * remain)
			return defaultScreen, &opacity
		} else if diff < 1 && input.Background.FadeIn {
			opacity := uint8(255 * diff)
			return defaultScreen, &opacity
		} else {
			return defaultScreen, nil
		}

	}

	t.LastScreenChange = time.Now()
	touch, _ := t.GetTouches()

	for i := 0; i < len(input.Touches); i++ {
		hitBox := input.Touches[i]

		if touch.X > hitBox.X && touch.X < hitBox.X+hitBox.Width {
			if touch.Y > hitBox.Y && touch.Y < hitBox.Y+hitBox.Height {

				if hitBox.Command.Type == "menu" {
					return hitBox.Command.Value, nil
				} else if hitBox.Command.Type == "command" {
					err := RunCommand(hitBox.Command.Value, false)
					if err != nil {
						return defaultScreen, nil
					}
					if hitBox.Command.ReturnScreen != "" {
						return hitBox.Command.ReturnScreen, nil
					} else {
						return defaultScreen, nil
					}
				}
			}
		}
	}

	return defaultScreen, nil
}

/*return new screen based on touches if appropriate*/
func DrawScreen(t *touchscreen.TouchScreen, name string, input Screen, s ScreenDetails) string {

	openvg.Start(s.Width, s.Height) // Start the picture
	if input.Background.Color != "" {
		openvg.BackgroundColor(input.Background.Color)
	} else {
		openvg.BackgroundColor("black")
	}

	ret, opacity := HandleTouches(t, input, name)

	// Default fill color
	openvg.FillColor("rgb(255,255,255)") // White text

	DrawLine(s, input.Line1, 400, opacity)
	DrawLine(s, input.Line2, 320, opacity)
	DrawLine(s, input.Line3, 240, opacity)
	DrawLine(s, input.Line4, 160, opacity)
	DrawLine(s, input.Line5, 80, opacity)

	openvg.End()

	return ret
}

func monitorService(s *screenservice.Server) {
	// Listen for incoming screens on the buffer
	for {
		if s.NumScreens() != 0 {
			screen := s.GetScreen()
			newScreen := GetScreenFromScreenRequest(screen)

			interruptScreen.Lock.Lock()
			if len(interruptScreen.Screens) == 0 {
				interruptScreen.LastShown = time.Now()
			}
			interruptScreen.Screens = append(interruptScreen.Screens, newScreen)
			interruptScreen.Lock.Unlock()
		}

		if s.HasImage() {
			img := s.GetImage()

			interruptScreen.Lock.Lock()
			interruptScreen.IncomingImage = img
			interruptScreen.Lock.Unlock()
			s.RemoveImage()
		}
		time.Sleep(100 * time.Millisecond)
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

	t := touchscreen.TouchScreen{}
	t.Init(touchscreen.FT62XX)

	t.Debug = true

	go runInterruptServer()

	var prevscreen string
	var screen Screen
	var err error

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

		if cur_screen != prevscreen {

			screen, err = GetScreenByName(cur_screen)
			if err != nil {
				panic(err)
			}
			prevscreen = cur_screen
		}

		cur_screen = DrawScreen(&t, cur_screen, screen, screenDetails)

		time.Sleep(5 * time.Millisecond)
	}
}
