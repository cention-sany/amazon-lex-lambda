package lexlambda

import (
	"strings"

	"github.com/cention-sany/basex"
)

func init() {
	alpha = make(map[rune]bool, len(base26Encoder))
	for _, r := range base26Encoder {
		alpha[r] = true
	}
}

// Base26Only will modifies s to suit base26 and replace any non-base26 char
// with 'x'.
func Base26Only(s string) string {
	s = strings.ToLower(s)
	var sb strings.Builder
	sb.Grow(len(s))
	for _, r := range s {
		if alpha[r] {
			sb.WriteRune(r)
		} else {
			sb.WriteRune('x')
		}
	}
	return sb.String()
}

var alpha map[rune]bool

var Base26 = mustEncoding()

func mustEncoding() *basex.Encoding {
	e, _ := newEncoding()
	return e
}

func newEncoding() (*basex.Encoding, error) {
	return basex.NewEncoding(base26Encoder)
}
