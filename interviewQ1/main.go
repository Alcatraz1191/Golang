package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*	input : 3421
	output : 3124
	input : 4321
	output : -1
	input : 8579
	output : 8975
*/

func findNext(num []string, n int) {
	for i := n - 1; i > 0; i-- {
		fmt.Println(num[i])
		numi, _ := strconv.Atoi(num[i])
		numi1, _ := strconv.Atoi(num[i-1])
		if numi > numi1 {
			break
		}
		if i == 1 && numi <= numi1 {
			fmt.Println("-1")
			return
		}
		x, _ := strconv.Atoi(num[i-1])
		smallest := i

		for j := i + 1; j < n; j++ {
			numj, _ := strconv.Atoi(num[j])
			nums, _ := strconv.Atoi(num[smallest])
			if numj > x && numj < nums {
				smallest = j
			}

		}
		num[smallest], num[i-1] = num[i-1], num[smallest]
		x = 0
		var numo int
		sort.Strings(num[i:]) //4321
		//fmt.Println("Sorted Num:", num)
		for i := 0; i < n; i++ {
			numo, _ = strconv.Atoi(num[i])
			x = x*10 + numo
		}

		fmt.Println("X:", x)
	}

}

func main() {
	digits := "8597"
	num := strings.Split(digits, "")
	fmt.Println(num)
	findNext(num, len(num))
}
