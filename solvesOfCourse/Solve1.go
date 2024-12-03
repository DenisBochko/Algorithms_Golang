package main

import (
	"fmt"
)

func main() {
	var mx int64
	var n int
	fmt.Scan(&n)

	arr := make([]int64, n)
	// res := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for _, v := range arr {
		if v > mx {
			mx = v
		}
	}

	// for i, v := range arr {
	// 	res[i] = strconv.FormatInt(v, 10)
	// }

	fmt.Println(mx)
}
