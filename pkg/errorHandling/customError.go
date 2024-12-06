package errorhandling

import (
	"errors"
	"fmt"
)

type customError struct {
	ErrorMsg  string
	ErrorCode int
}

func (ce *customError) Error() string {
	return fmt.Sprintf("%d - %v ", ce.ErrorCode, ce.ErrorMsg)
}

func fun(num int) (int, error) {
	if num&1 == 0 {
		return -1, &customError{"Number is Even", 404}
	}

	return num, nil
}

func CustomErrorHandling() {
	_, err := fun(50)
	var ce *customError
	if errors.As(err, &ce) {
		fmt.Println(ce.ErrorMsg)
		fmt.Println(ce.ErrorCode)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
