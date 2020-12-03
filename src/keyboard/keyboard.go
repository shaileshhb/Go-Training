package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetFloat gets string from user and returns float
func GetFloat() (float64, error) {

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, nil
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}
