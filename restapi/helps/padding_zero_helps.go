package helps

import (
	"fmt"
)

func PaddingZeros(num int64) string {
	s := fmt.Sprintf("%019d", num)
	return s
}

func PaddingZerosFloat(num float64) string {
	s := fmt.Sprintf("%030.10f", num)
	return s
}

func PaddingZerosUint64(num uint64) string {
	s := fmt.Sprintf("%019d", num)
	return s
}
