package helpers

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

func GenerateOTP(maxDigits uint32) (string, error) {
	bi, err := rand.Int(
		rand.Reader,
		big.NewInt(int64(math.Pow(10, float64(maxDigits)))),
	)
	if err != nil {
		fmt.Printf("Error while generating otp. Error: %v", err)
		return "", nil
	}
	return fmt.Sprintf("%0*d", maxDigits, bi), nil
}
