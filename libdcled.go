/*
 * ----------------------------------------------------------------------------
 * "THE PIZZA-WARE LICENSE" (Revision 42):
 * <whoami@dev-urandom.eu> wrote this file.  As long as you retain this notice you
 * can do whatever you want with this stuff. If we meet some day, and you think
 * this stuff is worth it, you can buy me a pizza in return.
 * ----------------------------------------------------------------------------
 */

package libdcled

var Debug = false

type DcLed struct {
	dsp         *display
	scroll_chan chan []Point
	kill_scroll chan bool
	kill_update chan bool
	scroll      bool
}

func NewDcLed() (*DcLed, error) {
	dsp, err := newDisplay()
	if err != nil {
		return nil, err
	}

	dsp.buffer.clear()
	k := make(chan bool)
	go dsp.updateDisplay(40, k)

	return &DcLed{
		dsp:         dsp,
		scroll_chan: nil,
		kill_scroll: make(chan bool),
		kill_update: k,
		scroll:      false}, nil
}

func (led *DcLed) PrintText(text *Text) {
	if led.scroll {
		led.kill_scroll <- true
		close(led.scroll_chan)
		led.scroll_chan = nil

		led.scroll = false
	}

	led.dsp.buffer.clear()
	led.dsp.buffer.drawPoints(text.points)
}

func (led *DcLed) ScrollText(text *Text) {
	if !led.scroll {
		led.scroll_chan = make(chan []Point)
		go led.dsp.scrollDisplay(led.scroll_chan, led.kill_scroll)
		led.scroll = true
	}

	led.scroll_chan <- text.points
}

func (led *DcLed) Kill() {
	if led.scroll {
		led.kill_scroll <- true
		close(led.kill_scroll)
		led.scroll = false
	}

	led.kill_update <- true
	close(led.kill_update)
}
