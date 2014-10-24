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
	"math"
	"sync"
)

type buffer struct {
	mu         sync.Mutex
	brightness byte
	rows       [7][3]byte
}

func (buf *buffer) getPacket(row int) ([]byte, error) {
	if row%2 == 1 || row > 6 {
		return nil, errors.New("Not a valid row")
	}

	if row == 6 {
		return []byte{buf.brightness, 0x06, buf.rows[row][2], buf.rows[row][1], buf.rows[row][0], 0x00, 0x00, 0x00}, nil
	}

	return []byte{buf.brightness, byte(row), buf.rows[row][2], buf.rows[row][1], buf.rows[row][0], buf.rows[row+1][2], buf.rows[row+1][1], buf.rows[row+1][0]}, nil
}

func (buf *buffer) clear() {
	buf.mu.Lock()
	defer buf.mu.Unlock()

	for i := 0; i < len(buf.rows); i++ {
		for j := 0; j < len(buf.rows[i]); j++ {
			buf.rows[i][j] = 0xff
		}
	}
}

func (buf *buffer) revert() {
	buf.mu.Lock()
	defer buf.mu.Unlock()

	for i := 0; i < len(buf.rows); i++ {
		for j := 0; j < len(buf.rows[i]); j++ {
			buf.rows[i][j] = ^buf.rows[i][j]
		}
	}
}

func (buf *buffer) drawPoints(points []Point) {
	buf.mu.Lock()
	defer buf.mu.Unlock()

	for _, p := range points {
		buf.drawPixel(p.x, p.y)
	}
}

func (buf *buffer) drawPixel(x int, y int) {
	if x < 1 || x > LEDX || y < 1 || y > LEDY {
		return
	}

	buf.rows[y-1][(x-1)>>3] &= ^(byte(0x01) << uint((x+7)%8))
}

func (buf *buffer) drawLine(x0 int, y0 int, x1 int, y1 int) {
	dx := int(math.Abs(float64(x1 - x0)))
	dy := (int(math.Abs(float64(y1-y0))) * -1)

	err := dx + dy

	sx := -1
	sy := -1

	if x0 < x1 {
		sx = 1
	}
	if y0 < y1 {
		sy = 1
	}

	for {
		buf.drawPixel(x0, y0)
		if x0 == x1 && y0 == y1 {
			break
		}

		e2 := 2 * err
		if e2 > dy {
			err += dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}
