package server

import (
	"fmt"
	"testing"
)

func TestDeleteElementFromSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	i := 2
	//in this operation, slice3's length is 5 and slice's length is 6. slice will be overwritten by [4,5,6] from index = 2.
	//the last element of slice(which is 6) is not untouched.
	slice3 := append(slice[:i], slice[i+1:]...)
	fmt.Printf("slice3:%v\n", slice3) //result is [1,2,4,5,6]
	fmt.Printf("slice:%v\n", slice)   //result is [1,2,4,5,6,6]
}

func TestDeleteElementFromSlice1(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}

	slice1 := delete(slice, 5)

	fmt.Printf("slice:%v\n", slice)
	fmt.Printf("slice1:%v\n", slice1)
}

func TestDeleteElementFromSlice2(t *testing.T) {
	slices := []int{6, 6, 5, 5, 4, 4, 3, 3}

	//for i, element := range slices {
	//	if element < 4 {
	//		if i == len(slices)-1 {
	//			slices = slices[:i]
	//		} else {
	//			slices = append(slices[:i], slices[i+1:]...)
	//		}
	//	}
	//}
	//
	//length := len(slices)
	//for i := 0; i < length; i++ {
	//	element := slices[i]
	//	if element < 4 {
	//		if i == len(slices)-1 {
	//			slices = slices[:i]
	//		} else {
	//			slices = append(slices[:i], slices[i+1:]...)
	//		}
	//	}
	//	length--
	//}
	var remains []int
	for _, element := range slices {
		if element >= 4 {
			remains = append(remains, element)
		}
	}
	slices = remains
	fmt.Printf("slices:%v\n", slices)
}

func delete(slice []int, i int) []int {
	var slice1 []int
	if i > len(slice)-1 {
		slice1 = slice
	} else if i == len(slice)-1 {
		slice1 = slice[:i]
	} else {
		slice1 = append(slice[:i], slice[i+1:]...)
	}
	return slice1
}
