package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println()
	res, err := sumPositive(-1, 2)
	fmt.Errorf("%d, 错误, %w\n", res, err)

}

func sumPositive(i, j int) (int, error) {
	if i <= 0 || j <= 0 {
		//return -1, errors.New("必须为整数")
		// 自定义错误
		//return -1, New("必须为整数")

		return -1, errors.New("必须为整数，原始错误")
	}
	return i + j, nil
}

// 自定义错误
func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// 传递
type wrapError struct {
	msg string
	err error
}

func (e *wrapError) Error() string {
	return e.msg
}

func (e *wrapError) Unwrap() error {
	return e.err

}
