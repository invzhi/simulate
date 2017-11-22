package simulate

import (
	"fmt"
	"log"
	"strconv"
)

func rightShift(num int64, l *uint) int64 {
	if num&(1<<(*l-1)) > 0 {
		num |= 1 << *l
	}
	*l = *l + 1
	return num
}

func printLine(l uint) {
	for n := uint(0); n < l; n++ {
		fmt.Print("-")
	}
	fmt.Println()
}

// Multiply will print calculation process with complement code and return x*y
// x_ should have double sign bit
func Multiply(x_, y_ string) string {
	x, err := strconv.ParseInt(x_, 2, 64)
	if err != nil {
		log.Fatalln("Can't parse string to int:", err)
	}
	y, err := strconv.ParseInt(y_, 2, 64)
	if err != nil {
		log.Fatalln("Can't parse string to int:", err)
	}

	fmt.Printf("%s * %s\n", x_, y_)

	sum := int64(0)
	t := int64(0)

	xl, yl := uint(len(x_)), uint(len(y_))
	sumL := xl

	rx := int64(uint64(-x) % (1 << xl))
	fmt.Printf("%0*b  [-x]\n\n", xl, rx)

	for i := uint(0); i < yl; i++ {
		fmt.Printf("%0*b|%0*b|%b\n", sumL, sum, yl-i, y, t)

		switch t - (y & 1) {
		case 1:
			fmt.Printf("%0*b +[+x]\n", xl, x)
			sum += x << i
		case -1:
			fmt.Printf("%0*b +[-x]\n", xl, rx)
			sum += rx << i
		}

		printLine(xl + yl + 3)

		sum %= 1 << sumL
		fmt.Printf("%0*b|%0*b|%b\n", sumL, sum, yl-i, y, t)

		if i < yl-1 {
			sum = rightShift(sum, &sumL)
		}

		t = y & 1
		y >>= 1
	}

	return fmt.Sprintf("%0*b", sumL, sum)
}
