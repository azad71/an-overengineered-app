package helpers

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math"
	"math/big"
)

func GenerateOTP(maxDigits int8) (string, error) {

	if maxDigits > 10 {
		return "", errors.New("maxDigits limit overflowed. Can't be greater than 10")
	}

	bi, _ := rand.Int(
		rand.Reader,
		big.NewInt(int64(math.Pow(10, float64(maxDigits)))),
	)

	return fmt.Sprintf("%0*d", maxDigits, bi), nil
}
