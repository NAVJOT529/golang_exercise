/* 60. Write a golang program to shuffle a given array of integers.  */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
*This function is used to shuffle a given array of integers
 */
/* func main() {
var arr = []int{1, 2, 3, 4, 5, 6}
fmt.Println(arr) */
/* array := map[int]int{}
for i := range arr {
	array[arr[i]] = arr[i]
}
var newArr []int
for i := range array {
	newArr = append(newArr, i)
}
fmt.Println(newArr) */
/* 	a := []int{1, 2, 3, 4, 5, 6, 7, 8} */

/*
	rand.Seed(time.Now().UnixNano())

for i := len(arr) - 1; i > 0; i-- { // Fisher–Yates shuffle

		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

fmt.Println(arr)
*/
func main() {
	rand.Seed(time.Now().UnixNano())
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr)
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println(arr)
	return
}

/*
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println(arr)
} */
