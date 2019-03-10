package go_type

import "fmt"

// 二维数组
func TwoImensionalArrays(cols, rows int) {

	fmt.Printf("二维数组TwoImensionalArrays(%d,%d)\n", cols, rows)

	raw := make([]int, cols*rows)

	for i := range raw { // range slice array's index
		raw[i] = i + 1
	}

	fmt.Printf(" raw: %+v , %p\n", raw, &raw[0])

	tbl := make([][]int, rows)

	for i := range tbl {
		tbl[i] = raw[i*cols : i*cols+cols] // range slice array
	}

	fmt.Printf(" tbl: %+v , %p\n", tbl, &tbl[0][0])
}
