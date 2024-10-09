package main

func main() {

	const name string = "jack"
	const msg = "hello"
	const num = 1
	const sum = (1+2+3)/2%100 + num

	const (
		Count = 1
		Name  = "Tom"
	)
	const (
		Size = 16
		Len  = 4
	)

	const (
		A = 1
		B
		c
	)
	//const iota = 0

	const (
		N = iota
		N1
		N2
	)
	const (
		M = iota * 2
		M1
		M2
	)

}
