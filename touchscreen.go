package main

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/exp/io/i2c"
)

const FT62XX_ADDR = 0x38
const FT62XX_REG_VENDID = 0xA8
const FT62XX_TOUCH_COUNT_REG = 0x02
const FT62XX_TOUCH_VAL_REG = 0x03

type TouchScreen struct {
	Device           *i2c.Device
	Touched          bool
	LastScreenChange time.Time
}

type TouchPoint struct {
	X int
	Y int
}

func (touchscreen *TouchScreen) Init() error {

	if touchscreen.Device != nil {
		return errors.New("Device must be unitiliased")
	}

	d, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, FT62XX_ADDR)
	if err != nil {
		panic(err)
	}
	touchscreen.Device = d

	buf := make([]byte, 1)
	err = touchscreen.Device.ReadReg(0xA8, buf)

	if err != nil {
		fmt.Println("Error readying register")
		return errors.New("Failed to read in the Vendor Reg")
	}
	if buf[0] != 0x11 {
		fmt.Println("Failed to get the right Vendor ID")
		return errors.New("Failed to match the expected Vendor ID")
	}

	threshold_buf := make([]byte, 1)
	threshold_buf[0] = 0x80
	err = touchscreen.Device.WriteReg(0x80, threshold_buf)
	if err != nil {
		fmt.Println("Failed to write to register")
		return errors.New("Failed to set the threshold")
	}

	touchscreen.Touched = false

	return nil
}

func (t *TouchScreen) GetTouchesCount() (int, error) {
	if t.Device == nil {
		return 0, errors.New("Touchscreen not initialised yet")
	}
	buf := make([]byte, 1)
	err := t.Device.ReadReg(FT62XX_TOUCH_COUNT_REG, buf)

	touchCount := int(buf[0])
	if touchCount > 2 {
		touchCount = 0
	}
	if t.Touched == true && touchCount == 0 {
		t.Touched = false
		return 0, err
	}

	if t.Touched == true {
		// We have already sent this touch - not again please
		return 0, err
	}

	if touchCount > 0 && touchCount < 3 {
		t.Touched = true
	}
	return touchCount, err
}

func (t *TouchScreen) GetTouches() (TouchPoint, error) {
	ret := TouchPoint{X: 0, Y: 0}
	buf := make([]byte, 4)
	err := t.Device.ReadReg(FT62XX_TOUCH_VAL_REG, buf)
	if err != nil {
		return ret, errors.New("Failed to read in the touch registers")
	}

	// Mask off the MSB and stick it to the LSB
	ret.X = int(buf[0]) & 0xf
	ret.X = ret.X << 8
	ret.X = ret.X | int(buf[1])

	ret.Y = int(buf[2] & 0xf)
	ret.Y = ret.Y << 8
	ret.Y = ret.Y | int(buf[3])

	return ret, nil
}
