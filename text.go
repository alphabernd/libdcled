/*
 * ----------------------------------------------------------------------------
 * "THE PIZZA-WARE LICENSE" (Revision 42):
 * <whoami@dev-urandom.eu> wrote this file.  As long as you retain this notice you
 * can do whatever you want with this stuff. If we meet some day, and you think
 * this stuff is worth it, you can buy me a pizza in return.
 * ----------------------------------------------------------------------------
 */

package libdcled

const (
	STANDARD_FONT = iota
)

type Text struct {
	str    string
	points []Point
}

func NewText(message string, font int) *Text {
	return &Text{str: message, points: getTextPoints(message, font)}
}

func getLetter(c rune, font int) []Point {
	switch font {
	default:
		letter := make([]Point, len(stdfont[c]), len(stdfont[c]))
		copy(letter, stdfont[c])
		return letter
	}
}

func getTextPoints(text string, font int) []Point {
	p := []Point{}
	offset := 0

	for _, c := range text {
		letter := getLetter(c, font)
		width := 0
		for j := range letter {
			if letter[j].x > width {
				width = letter[j].x
			}
			letter[j].x += offset
		}

		p = append(p, letter...)

		if c == ' ' {
			offset += 2
		}
		offset += (width + 1)
	}

	return p
}

var stdfont = map[rune][]Point{
	'A': []Point{Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{1, 7}, Point{2, 1}, Point{2, 4}, Point{3, 1}, Point{3, 4}, Point{4, 2}, Point{4, 3}, Point{4, 4}, Point{4, 5}, Point{4, 6}, Point{4, 7}},
	'a': []Point{Point{1, 3}, Point{1, 6}, Point{2, 2}, Point{2, 5}, Point{2, 7}, Point{3, 2}, Point{3, 5}, Point{3, 7}, Point{4, 3}, Point{4, 4}, Point{4, 5}, Point{4, 6}, Point{4, 7}},
	'B': []Point{Point{1, 1}, Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{1, 7}, Point{2, 1}, Point{2, 4}, Point{2, 7}, Point{3, 1}, Point{3, 4}, Point{3, 7}, Point{4, 2}, Point{4, 3}, Point{4, 5}, Point{4, 6}},
	'b': []Point{Point{1, 1}, Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{1, 7}, Point{2, 4}, Point{2, 7}, Point{3, 4}, Point{3, 7}, Point{4, 5}, Point{4, 6}},
	'C': []Point{Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{2, 1}, Point{2, 7}, Point{3, 1}, Point{3, 7}, Point{4, 2}, Point{4, 6}},
	'c': []Point{Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{2, 2}, Point{2, 7}, Point{3, 2}, Point{3, 7}, Point{4, 3}, Point{4, 6}},
	'D': []Point{Point{1, 1}, Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{1, 7}, Point{2, 1}, Point{2, 7}, Point{3, 1}, Point{3, 7}, Point{4, 2}, Point{4, 3}, Point{4, 4}, Point{4, 5}, Point{4, 6}},
	'd': []Point{Point{1, 5}, Point{1, 6}, Point{2, 4}, Point{2, 7}, Point{3, 4}, Point{3, 7}, Point{4, 1}, Point{4, 2}, Point{4, 3}, Point{4, 4}, Point{4, 5}, Point{4, 6}, Point{4, 7}},
	'E': []Point{Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{2, 1}, Point{2, 4}, Point{2, 7}, Point{3, 1}, Point{3, 4}, Point{3, 7}, Point{4, 1}, Point{4, 7}},
	'e': []Point{Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{2, 2}, Point{2, 4}, Point{2, 7}, Point{3, 2}, Point{3, 4}, Point{3, 7}, Point{4, 3}, Point{4, 6}},
	'f': []Point{Point{1, 4}, Point{2, 2}, Point{2, 3}, Point{2, 4}, Point{2, 5}, Point{2, 6}, Point{2, 7}, Point{3, 1}, Point{3, 4}, Point{4, 2}},
	'g': []Point{Point{1, 3}, Point{1, 6}, Point{2, 2}, Point{2, 4}, Point{2, 7}, Point{3, 2}, Point{3, 4}, Point{3, 7}, Point{4, 3}, Point{4, 4}, Point{4, 5}, Point{4, 6}},
	'H': []Point{Point{1, 1}, Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{1, 7}, Point{2, 4}, Point{3, 4}, Point{4, 1}, Point{4, 2}, Point{4, 3}, Point{4, 4}, Point{4, 5}, Point{4, 6}, Point{4, 7}},
	'h': []Point{Point{1, 1}, Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{1, 7}, Point{2, 4}, Point{3, 4}, Point{4, 5}, Point{4, 6}, Point{4, 7}},
	'i': []Point{Point{1, 2}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{1, 7}},
	'j': []Point{Point{1, 6}, Point{2, 7}, Point{3, 2}, Point{3, 4}, Point{3, 5}, Point{3, 6}},
	'K': []Point{Point{1, 1}, Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{1, 7}, Point{2, 3}, Point{2, 5}, Point{3, 2}, Point{3, 6}, Point{4, 1}, Point{4, 7}},
	'k': []Point{Point{1, 1}, Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{1, 7}, Point{2, 4}, Point{2, 5}, Point{3, 3}, Point{3, 6}, Point{4, 2}, Point{4, 7}},
	'l': []Point{Point{1, 1}, Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{2, 7}, Point{3, 6}},
	'm': []Point{Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{1, 7}, Point{2, 2}, Point{3, 3}, Point{3, 4}, Point{3, 5}, Point{3, 6}, Point{3, 7}, Point{4, 2}, Point{5, 3}, Point{5, 4}, Point{5, 5}, Point{5, 6}, Point{5, 7}},
	'n': []Point{Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{1, 7}, Point{2, 2}, Point{3, 3}, Point{3, 4}, Point{3, 5}, Point{3, 6}, Point{3, 7}},
	'o': []Point{Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{2, 2}, Point{2, 7}, Point{3, 2}, Point{3, 7}, Point{4, 3}, Point{4, 4}, Point{4, 5}, Point{4, 6}},
	'p': []Point{Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{1, 7}, Point{2, 2}, Point{2, 5}, Point{3, 2}, Point{3, 5}, Point{4, 3}, Point{4, 4}},
	'q': []Point{Point{1, 3}, Point{1, 4}, Point{2, 2}, Point{2, 5}, Point{3, 2}, Point{3, 5}, Point{4, 3}, Point{4, 4}, Point{4, 5}, Point{4, 6}, Point{4, 7}},
	'r': []Point{Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{1, 7}, Point{2, 2}, Point{3, 2}, Point{4, 3}},
	's': []Point{Point{1, 3}, Point{1, 7}, Point{2, 2}, Point{2, 4}, Point{2, 7}, Point{3, 2}, Point{3, 4}, Point{3, 7}, Point{4, 2}, Point{4, 5}, Point{4, 6}},
	't': []Point{Point{1, 2}, Point{2, 1}, Point{2, 2}, Point{2, 3}, Point{2, 4}, Point{2, 5}, Point{2, 6}, Point{3, 2}, Point{3, 7}, Point{4, 6}},
	'u': []Point{Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{2, 7}, Point{3, 7}, Point{4, 2}, Point{4, 3}, Point{4, 4}, Point{4, 5}, Point{4, 6}},
	'v': []Point{Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{2, 7}, Point{3, 2}, Point{3, 3}, Point{3, 4}, Point{3, 5}, Point{3, 6}},
	'w': []Point{Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{2, 7}, Point{3, 2}, Point{3, 3}, Point{3, 4}, Point{3, 5}, Point{3, 6}, Point{4, 7}, Point{5, 2}, Point{5, 3}, Point{5, 4}, Point{5, 5}, Point{5, 6}},
	'x': []Point{Point{1, 2}, Point{1, 3}, Point{1, 6}, Point{1, 7}, Point{2, 4}, Point{2, 5}, Point{3, 4}, Point{3, 5}, Point{4, 2}, Point{4, 3}, Point{4, 6}, Point{4, 7}},
	'y': []Point{Point{1, 2}, Point{1, 3}, Point{2, 4}, Point{2, 7}, Point{3, 4}, Point{3, 7}, Point{4, 2}, Point{4, 3}, Point{4, 4}, Point{4, 5}, Point{4, 6}},
	'z': []Point{Point{1, 2}, Point{1, 6}, Point{2, 2}, Point{2, 5}, Point{2, 7}, Point{3, 2}, Point{3, 4}, Point{3, 7}, Point{4, 3}, Point{4, 7}},
	'1': []Point{Point{1, 2}, Point{2, 1}, Point{2, 2}, Point{2, 3}, Point{2, 4}, Point{2, 5}, Point{2, 6}, Point{2, 7}},
	'2': []Point{Point{1, 2}, Point{1, 6}, Point{2, 1}, Point{2, 5}, Point{2, 7}, Point{3, 1}, Point{3, 4}, Point{3, 7}, Point{4, 2}, Point{4, 3}, Point{4, 7}},
	'3': []Point{Point{1, 2}, Point{1, 6}, Point{2, 1}, Point{2, 4}, Point{2, 7}, Point{3, 1}, Point{3, 4}, Point{3, 7}, Point{4, 2}, Point{4, 3}, Point{4, 5}, Point{4, 6}},
	'4': []Point{Point{1, 1}, Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{2, 4}, Point{3, 4}, Point{4, 2}, Point{4, 3}, Point{4, 4}, Point{4, 5}, Point{4, 6}, Point{4, 7}},
	'5': []Point{Point{1, 1}, Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{1, 7}, Point{2, 1}, Point{2, 4}, Point{2, 7}, Point{3, 1}, Point{3, 4}, Point{3, 7}, Point{4, 1}, Point{4, 4}, Point{4, 5}, Point{4, 6}, Point{4, 7}},
	'6': []Point{Point{1, 1}, Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{1, 7}, Point{2, 1}, Point{2, 4}, Point{2, 7}, Point{3, 1}, Point{3, 4}, Point{3, 7}, Point{4, 1}, Point{4, 4}, Point{4, 5}, Point{4, 6}, Point{4, 7}},
	'7': []Point{Point{1, 1}, Point{2, 1}, Point{2, 4}, Point{3, 1}, Point{3, 2}, Point{3, 3}, Point{3, 4}, Point{3, 5}, Point{3, 6}, Point{3, 7}, Point{4, 4}},
	'8': []Point{Point{1, 1}, Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{1, 7}, Point{2, 1}, Point{2, 4}, Point{2, 7}, Point{3, 1}, Point{3, 4}, Point{3, 7}, Point{4, 1}, Point{4, 2}, Point{4, 3}, Point{4, 4}, Point{4, 5}, Point{4, 6}, Point{4, 7}},
	'9': []Point{Point{1, 1}, Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{1, 7}, Point{2, 1}, Point{2, 4}, Point{2, 7}, Point{3, 1}, Point{3, 4}, Point{3, 7}, Point{4, 1}, Point{4, 2}, Point{4, 3}, Point{4, 4}, Point{4, 5}, Point{4, 6}, Point{4, 7}},
	'0': []Point{Point{1, 2}, Point{1, 3}, Point{1, 4}, Point{1, 5}, Point{1, 6}, Point{2, 1}, Point{2, 7}, Point{3, 1}, Point{3, 7}, Point{4, 2}, Point{4, 3}, Point{4, 4}, Point{4, 5}, Point{4, 6}},
	' ': []Point{},
	'%': []Point{Point{1, 2}, Point{1, 3}, Point{1, 6}, Point{2, 2}, Point{2, 3}, Point{2, 5}, Point{3, 4}, Point{4, 3}, Point{4, 5}, Point{4, 6}, Point{5, 2}, Point{5, 5}, Point{5, 6}, Point{1, 2}, Point{1, 2}},
	'.': []Point{Point{1, 7}},
	'/': []Point{Point{1, 6}, Point{1, 7}, Point{2, 4}, Point{2, 5}, Point{2, 6}, Point{3, 2}, Point{3, 3}, Point{3, 4}, Point{4, 1}, Point{4, 2}},
}
