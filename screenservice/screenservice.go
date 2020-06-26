package screenservice

import (
	fmt "fmt"

	"golang.org/x/net/context"
)

type Server struct{}

func (s *Server) SendScreen(ctx context.Context, in *ScreenRequest) (*ScreenResponse, error) {
	fmt.Println("Got a new screen in")
	fmt.Println(in.Line1)
	return &ScreenResponse{Status: "ok"}, nil
}
