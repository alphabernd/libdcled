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
	switch row {
	case 0:
		return []byte{buf.brightness, 0x00, buf.rows[0][2], buf.rows[0][1], buf.rows[0][0], buf.rows[1][2], buf.rows[1][1], buf.rows[1][0]}, nil
	case 2:
		return []byte{buf.brightness, 0x02, buf.rows[2][2], buf.rows[2][1], buf.rows[2][0], buf.rows[3][2], buf.rows[3][1], buf.rows[3][0]}, nil
	case 4:
		return []byte{buf.brightness, 0x04, buf.rows[4][2], buf.rows[4][1], buf.rows[4][0], buf.rows[5][2], buf.rows[5][1], buf.rows[5][0]}, nil
	case 6:
		return []byte{buf.brightness, 0x06, buf.rows[6][2], buf.rows[6][1], buf.rows[6][0], 0x00, 0x00, 0x00}, nil
	default:
		return nil, errors.New("Unknown row")
	}
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
	if x < 1 || x > 21 || y < 1 || y > 7 {
		return
	}

	var pixel byte = 0x01

	offset := (x % 8) - 1
	if offset == -1 {
		offset = 7
	}

	pixel = pixel << uint(offset)
	pixel = ^pixel

	switch {
	case x < 9:
		x = 0
	case x < 17:
		x = 1
	default:
		x = 2
	}

	buf.rows[y-1][x] = buf.rows[y-1][x] & pixel
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
