package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

type NorgateMathError struct {
	lat, long string
	err       error
}

func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}
func main() {
	fmt.Printf("%T\n", ErrNorgateMath)
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

	_, err2 := Sqrt2(-10.23)
	if err2 != nil {
		log.Println(err2)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		//return 0, errors.New("norgate math: square root of negative number")
		return 0, ErrNorgateMath
		// ErrNorgateMath2 = errors.New("norgate math: square root of negative number")
		// return 0, ErrNorgateMath2
	}
	return 42, nil
}

func Sqrt2(f float64) (float64, error) {
	if f < 0 {
		name := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, &NorgateMathError{"50.2289 N", "99.4656 W", name}
	}
	return 42, nil
}
