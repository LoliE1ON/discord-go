package HexHelper

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Convert hex to int
func HexToInt(color string) (colorInt int, err error) {

	n, err := strconv.ParseInt(strings.Replace(color, "#", "", -1), 16, 32)
	if err != nil {
		err = errors.Wrap(err, "Error convert")
		return
	}
	return int(n), err

}
