/*
 * ----------------------------------------------------------------------------
 * "THE PIZZA-WARE LICENSE" (Revision 42):
 * <whoami@dev-urandom.eu> wrote this file.  As long as you retain this notice you
 * can do whatever you want with this stuff. If we meet some day, and you think
 * this stuff is worth it, you can buy me a pizza in return.
 * ----------------------------------------------------------------------------
 */

package libdcled

import (
	"errors"
	"fmt"
	libusb "github.com/OneOfOne/go-libusb"
	"os"
	"time"
)

const (
	vendor_id  = 0x1D34
	product_id = 0x0013
)

type Point struct {
	x int
	y int
}

type display struct {
	buffer *buffer
	dev    *libusb.Device
}

func newDisplay() (*display, error) {
	libusb.Init()

	dev := libusb.Open(vendor_id, product_id)
	if dev == nil {
		fmt.Fprintf(os.Stderr, "%s\n", "Could not find Device")
		return nil, errors.New("Could not find Device")
	}

	return &display{buffer: &buffer{}, dev: dev}, nil
}

func (dsp *display) updateDisplay(rate time.Duration, kill chan bool) {
	for {
		select {
		case k := <-kill:
			if k {
				return
			}
		default:
			if Debug {
				fmt.Fprintf(os.Stderr, "%s\n----------\n", "dsp.updateDisplay")
			}
			dsp.buffer.mu.Lock()
			for i := 0; i <= 6; i += 2 {
				packet, err := dsp.buffer.getPacket(i)
				if err != nil {
					fmt.Printf("%v - update display error\n", err)
					continue
				}
				if Debug {
					fmt.Fprintf(os.Stderr, "%v\n", packet)
				}
				dsp.dev.ControlMsg(0x21, 0x09, 0, 0, packet)
			}
			dsp.buffer.mu.Unlock()
			if Debug {
				fmt.Fprint(os.Stderr, "----------\n")
			}

			time.Sleep(rate * time.Millisecond)
		}
	}
}

func (dsp *display) scrollDisplay(p <-chan []Point, kill chan bool) {
	points := <-p
	l := 0

	for i := range points {
		points[i].x += 21
	}

	for {
		select {
		case k := <-kill:
			if k {
				return
			}
		case points = <-p: // new content
			for i := range points {
				points[i].x += 22
			}
		default:
			last_x := 0

			for i := range points {
				points[i].x -= 1
				if last_x < points[i].x {
					last_x = points[i].x
				}
			}
			l += 1

			if Debug {
				fmt.Fprintf(os.Stderr, "%s\n----------\n", "dsp.scrollDisplay")
				fmt.Fprintf(os.Stderr, "%v\n", points)
				fmt.Fprint(os.Stderr, "----------\n")
			}

			dsp.buffer.clear()
			dsp.buffer.drawPoints(points)

			if last_x <= 0 {
				for i := range points {
					points[i].x += l
				}
				l = 0
			}
		}
		time.Sleep(60 * time.Millisecond)
	}
}
