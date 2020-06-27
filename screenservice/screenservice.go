package screenservice

import (
	fmt "fmt"

	"golang.org/x/net/context"
)

type Server struct{}

var IncomingScreens []ScreenRequest

func (s *Server) SendScreen(ctx context.Context, in *ScreenRequest) (*ScreenResponse, error) {
	fmt.Println("Got a new screen in")
	fmt.Println(in.Line1)
	var newScreen ScreenRequest
	newScreen.Line1 = in.Line1
	newScreen.Line2 = in.Line2
	newScreen.Line3 = in.Line3
	newScreen.Line4 = in.Line4
	newScreen.Line5 = in.Line5
	newScreen.Length = in.Length
	newScreen.ShowCountdown = in.ShowCountdown

	IncomingScreens = append(IncomingScreens, newScreen)

	fmt.Printf("%d screens in list\n", len(IncomingScreens))

	return &ScreenResponse{Status: "ok"}, nil
}

func (s *Server) NumScreens() int {
	return len(IncomingScreens)
}

func (s *Server) GetScreen() *ScreenRequest {
	if len(IncomingScreens) < 1 {
		return nil
	}
	var ret *ScreenRequest
	ret = &IncomingScreens[0]
	IncomingScreens = IncomingScreens[1:]

	return ret
}