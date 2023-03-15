package main

import "fmt"

/*
*This function is used to take the values of 25 numbers into an array
 */
func main() {
	var num [25]int
	for i := 0; i < 25; i++ {
		fmt.Println("Enter any number")
		fmt.Scanln(&num[i])
	}
	Selection(num)
	fmt.Printf("\n")
	BubbleSort(num)
	fmt.Printf("\n")
	InsertionSort(num)
	return
}

/* This function is used to implement the algorithm of selection sort */
func Selection(num [25]int) {
	fmt.Printf("By selection method      ")
	for i := 0; i < 24; i++ {
		for j := i; j < 24; j++ {
			if num[i] > num[j+1] {
				swap(&num[i], &num[j+1])
			}
		}
	}
	for i := 0; i < 25; i++ {
		fmt.Printf("%d ", num[i])
	}
	return
}

/* This function is used to implement the algorithm of selection sort */
func BubbleSort(num [25]int) {
	fmt.Printf("By Bubble Sort method    ")
	for i := 0; i < 24; i++ {
		for j := 0; j < 24-i; j++ {
			if num[j] > num[j+1] {
				swap(&num[j], &num[j+1])
			}
		}
	}
	for i := 0; i < 25; i++ {
		fmt.Printf("%d ", num[i])
	}
	return
}

/* This function is used to implement the algorithm of Insertion sort */
func InsertionSort(num [25]int) {
	fmt.Printf("By Insertion Sort method ")
	var pin int
	for i := 1; i <= 24; i++ {
		pin = num[i]
		j := i - 1
		for ; j >= 0 && num[j] > pin; j-- {
			num[j+1] = num[j]
		}
		num[j+1] = pin
	}
	for i := 0; i < 25; i++ {
		fmt.Printf("%d ", num[i])
	}
	return
}

/* This function is used for swap the values */
func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
