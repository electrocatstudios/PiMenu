package screenservice

import (
	"bytes"
	"image"
	"sync"

	"golang.org/x/net/context"
)

type Server struct{}

type IncomingImageUpate struct {
	Image *image.Image
	Lock  sync.Mutex
}

var IncomingScreens []ScreenRequest
var IncomingImage IncomingImageUpate

func (s *Server) SendScreen(ctx context.Context, in *ScreenRequest) (*ScreenResponse, error) {
	IncomingScreens = append(IncomingScreens, *in)
	return &ScreenResponse{Status: "ok"}, nil
}

func (s *Server) SendImage(ctx context.Context, in *ScreenImage) (*ScreenResponse, error) {
	img, _, _ := image.Decode(bytes.NewReader(in.ImageData))
	IncomingImage.Lock.Lock()
	IncomingImage.Image = &img
	IncomingImage.Lock.Unlock()
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

func (s *Server) HasImage() bool {
	IncomingImage.Lock.Lock()
	if IncomingImage.Image != nil {
		IncomingImage.Lock.Unlock()
		return true
	}
	IncomingImage.Lock.Unlock()
	return false
}

func (s *Server) GetImage() *image.Image {
	IncomingImage.Lock.Lock()
	ret := IncomingImage.Image
	IncomingImage.Image = nil
	IncomingImage.Lock.Unlock()
	return ret
}
