// https://app.codility.com/programmers/lessons/8-leader/
package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// Implement your solution here
	if len(A) == 1 {
		return 0
	}
	checker := map[int]int32{}
	for i := 0; i < len(A); i++ {
		_, exist := checker[A[i]]
		// fmt.Println("the num = ", A[i])
		if exist {
			checker[A[i]] += 1
			// fmt.Println(checker[A[i]])
			if checker[A[i]] > int32(len(A)/2) {
				return i
			}
		} else {
			checker[A[i]] = 1
		}
	}
	return -1
}
