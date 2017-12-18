package simulate

import (
	"fmt"
	"log"
	"strconv"
)

// Divide will print alculation process with origin code and return x/y
// x_, y_ should begin with '0'
func Divide(x_ string, y_ string) []int {
	x, err := strconv.ParseInt(x_, 2, 64)
	if err != nil {
		log.Fatalln("Can't parse string to int:", err)
	}
	y, err := strconv.ParseInt(y_, 2, 64)
	if err != nil {
		log.Fatalln("Can't parse string to int:", err)
	}

	fmt.Printf("%s / %s\n", x_, y_)

	if len(x_) != len(y_) {
		log.Fatalln("x_ and y_ diff")
	}

	l := uint(len(x_))

	ry := int64(uint64(-y) % (1 << l))
	fmt.Printf("%0*b  [-y]\n\n", l, ry)

	b := make([]int, 0, l)

	is := true

	for i := uint(0); i < l; i++ {
		fmt.Printf("%0*b %v\n", l, x, b)
		if i != 0 {
			x <<= 1
			x %= 1 << l
			fmt.Printf("%0*b %v\n", l, x, b)
		}
		if is {
			x += ry
			fmt.Printf("%0*b +[-y]\n\n", l, ry)
		} else {
			x += y
			fmt.Printf("%0*b +[+y]\n\n", l, y)
		}
		x %= 1 << l

		// fmt.Printf("%0*b << debug\n", l, x & (1 << (l-1)))
		if x&(1<<(l-1)) > 0 {
			b = append(b, 0)
			is = false
		} else {
			b = append(b, 1)
			is = true
		}
	}
	fmt.Printf("%0*b %v\n", l, x, b)
	return b
}
