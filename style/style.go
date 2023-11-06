package style

import (
	"fmt"

	"go.uber.org/zap"
)

func DoSmth(i int) error {
	res := make([]int, 0)

	const expectedI = 42
	if i == expectedI {
		for j := 0; j < i; j++ {
			res = append(res, j*i)
		}

		zap.S().Infof("полученный слайс: %v", res)
		return nil
	} else {
		return fmt.Errorf("переданный параметр не %d, а %d", expectedI, i)
	}
}
